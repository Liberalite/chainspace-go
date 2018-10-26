package sbac // import "chainspace.io/prototype/sbac"

import (
	"sync"

	"chainspace.io/prototype/internal/log"
	"chainspace.io/prototype/internal/log/fld"
)

type StateTable struct {
	actions     map[State]Action
	transitions map[StateTransition]Transition
	onEvent     func(tx *TxDetails, event *Event) error
}

type StateTransition struct {
	From State
	To   State
}

type SignedDecision struct {
	Decision  SBACDecision
	Signature []byte
}

type TxDetails struct {
	Consensus1Tx      *SBACTransaction
	Consensus2Tx      *SBACTransaction
	ConsensusCommitTx *SBACTransaction

	Mu              sync.Mutex
	CommitDecisions map[uint64]SignedDecision
	Phase1Decisions map[uint64]SignedDecision
	Phase2Decisions map[uint64]SignedDecision

	CheckersEvidences map[uint64][]byte
	ID                []byte
	Raw               []byte
	Result            chan bool
	Tx                *Transaction
	HashID            uint32
}

// Action specify an action to execute when a new event is triggered.
// it returns a State, which will be either the new actual state, the next state
// which may required a transition from the current state to the new one (see the transition
// table
type Action func(tx *TxDetails) (State, error)

// Transition are called when the state is change from a current state to a new one.
// return a State which may involved a new transition as well.
type Transition func(tx *TxDetails) (State, error)

type Event struct {
	msg    *SBACMessage
	peerID uint64
}

type StateMachine struct {
	txDetails *TxDetails
	events    *pendingEvents
	state     State
	table     *StateTable
	mu        sync.Mutex
}

func (sm *StateMachine) Close() {
	sm.events.Close()
}

func (sm *StateMachine) Reset() {
	sm.state = StateWaitingForConsensus1
}

func (sm *StateMachine) StateReport() *StateReport {
	sm.txDetails.Mu.Lock()
	defer sm.txDetails.Mu.Unlock()
	s := StateReport{
		HashID:          sm.txDetails.HashID,
		State:           sm.State().String(),
		CommitDecisions: map[uint64]bool{},
		Phase1Decisions: map[uint64]bool{},
		Phase2Decisions: map[uint64]bool{},
		PendingEvents:   int32(sm.events.Len()),
	}
	for k, v := range sm.txDetails.CommitDecisions {
		s.CommitDecisions[k] = v.Decision == SBACDecision_ACCEPT
	}
	for k, v := range sm.txDetails.Phase1Decisions {
		s.Phase1Decisions[k] = v.Decision == SBACDecision_ACCEPT
	}
	for k, v := range sm.txDetails.Phase2Decisions {
		s.Phase2Decisions[k] = v.Decision == SBACDecision_ACCEPT
	}
	return &s
}

func (sm *StateMachine) State() State {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.state
}

func (sm *StateMachine) setState(newState State) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.state = newState
}

func (sm *StateMachine) applyTransition(transitionTo State) error {
	for {
		curState := sm.State()
		txtransition := StateTransition{curState, transitionTo}
		f, ok := sm.table.transitions[txtransition]
		if !ok {
			// no more transitions available, this is not an error
			return nil
		}

		log.Info("applying transition",
			log.Uint32("id", sm.txDetails.HashID),
			log.String("old_state", curState.String()),
			log.String("new_state", transitionTo.String()),
		)
		nextstate, err := f(sm.txDetails)
		if err != nil {
			log.Error("unable to apply transition",
				log.Uint32("id", sm.txDetails.HashID),
				log.String("old_state", curState.String()),
				log.String("new_state", transitionTo.String()),
				fld.Err(err),
			)
			return err
		}
		sm.setState(transitionTo)
		transitionTo = nextstate
	}
}

func (sm *StateMachine) moveState() error {
	for {
		curState := sm.State()
		// first try to execute an action if possible with the current state
		action, ok := sm.table.actions[curState]
		if !ok {
			log.Error("unable to find an action to map with the current state",
				log.String("state", curState.String()), fld.TxID(sm.txDetails.HashID))
			return nil
		}
		log.Info("applying action",
			log.Uint32("id", sm.txDetails.HashID),
			log.String("state", curState.String()),
		)
		newstate, err := action(sm.txDetails)
		if err != nil {
			log.Error("unable to execute action", fld.Err(err))
			return err
		}
		if newstate == sm.state {
			// action returned the same state, we can return now as this action is not ready to be completed
			// although this is not an error
			return nil
		}
		// action succeed, we try to find a transition, if any we apply it, if none available, just set the new state.
		txtransition := StateTransition{curState, newstate}
		// if a transition exist for the new state, apply it
		if _, ok := sm.table.transitions[txtransition]; ok {
			err = sm.applyTransition(newstate)
			if err != nil {
				log.Error("unable to apply transition", fld.Err(err))
			}
		} else {
			// else save the new state directly
			sm.setState(newstate)
		}
	}
}

func (sm *StateMachine) consumeEvent(e *Event) bool {
	curState := sm.State()
	if log.AtDebug() {
		log.Info("processing new event",
			log.Uint32("id", sm.txDetails.HashID),
			log.String("state", curState.String()),
			fld.PeerID(e.peerID),
		)
	}
	sm.table.onEvent(sm.txDetails, e)
	if curState == StateSucceeded || curState == StateAborted {
		if log.AtDebug() {
			log.Debug("statemachine reach end", log.String("final_state", sm.state.String()))
		}
		return true
	}
	err := sm.moveState()
	if err != nil {
		log.Error("something happend while moving states", fld.Err(err))
	}

	return true
}

func (sm *StateMachine) OnEvent(e *Event) {
	sm.events.OnEvent(e)
}

func NewStateMachine(table *StateTable, txDetails *TxDetails, initialState State) *StateMachine {
	if log.AtDebug() {
		log.Debug("starting new statemachine", fld.TxID(txDetails.HashID))
	}
	log.Info("starting new statemachine", fld.TxID(txDetails.HashID))
	sm := &StateMachine{
		state:     initialState,
		table:     table,
		txDetails: txDetails,
	}
	sm.events = NewPendingEvents(sm.consumeEvent)
	go sm.events.Run()
	return sm
}

func NewTxDetails(txID, raw []byte, tx *Transaction, evidences map[uint64][]byte) *TxDetails {
	return &TxDetails{
		CheckersEvidences: evidences,
		CommitDecisions:   map[uint64]SignedDecision{},
		ID:                txID,
		Phase1Decisions:   map[uint64]SignedDecision{},
		Phase2Decisions:   map[uint64]SignedDecision{},
		Raw:               raw,
		Result:            make(chan bool),
		Tx:                tx,
		HashID:            ID(txID),
	}
}