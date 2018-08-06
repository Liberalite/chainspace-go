package broadcast

import (
	"context"
	"sync"

	"chainspace.io/prototype/byzco"
	"chainspace.io/prototype/log"
	"chainspace.io/prototype/log/fld"
)

type blockData struct {
	id    byzco.BlockID
	links []byzco.BlockID
	ref   *SignedData
}

type depgraph struct {
	await   map[byzco.BlockID][]byzco.BlockID
	cond    *sync.Cond // protects in
	ctx     context.Context
	icache  map[byzco.BlockID]bool
	in      []*blockData
	pending map[byzco.BlockID]*blockData
	out     chan *blockData
	self    uint64
	store   *store
}

func (d *depgraph) add(info *blockData) {
	d.cond.L.Lock()
	d.in = append(d.in, info)
	d.cond.L.Unlock()
	d.cond.Signal()
}

func (d *depgraph) addPending(block *blockData, deps []byzco.BlockID) {
	d.pending[block.id] = block
	for _, ref := range deps {
		await, exists := d.await[ref]
		if exists {
			exists = false
			for _, id := range await {
				if id == block.id {
					exists = true
					break
				}
			}
			if !exists {
				d.await[ref] = append(await, block.id)
			}
		} else {
			d.await[ref] = []byzco.BlockID{block.id}
		}
	}
}

func (d *depgraph) isIncluded(id byzco.BlockID) bool {
	inc, exists := d.icache[id]
	if exists {
		return inc
	}
	inc, err := d.store.isIncluded(id)
	if err != nil {
		log.Fatal("Couldn't check if block has been included", fld.Err(err))
	}
	d.icache[id] = inc
	return inc
}

func (d *depgraph) markIncluded(id byzco.BlockID) {
	d.icache[id] = true
}

func (d *depgraph) process() {
	i := 0
	for {
		// Prune the included cache every 100 iterations.
		i++
		if i%100 == 0 {
			if len(d.icache) > 1000 {
				ncache := map[byzco.BlockID]bool{}
				j := 0
				for k, v := range d.icache {
					ncache[k] = v
					j++
					if j == 1000 {
						break
					}
				}
				d.icache = ncache
			}
		}
		d.cond.L.Lock()
		for len(d.in) == 0 {
			d.cond.Wait()
			select {
			case <-d.ctx.Done():
				d.cond.L.Unlock()
				return
			default:
			}
		}
		info := d.in[0]
		d.in = d.in[1:]
		d.cond.L.Unlock()
		if !d.processBlock(info) {
			continue
		}
		first := true
		processed := []byzco.BlockID{info.id}
		seen := map[byzco.BlockID]bool{
			info.id: true,
		}
		for len(processed) > 0 {
			next := processed[0]
			processed = processed[1:]
			var nawait []byzco.BlockID
			for _, revdep := range d.await[next] {
				if seen[revdep] {
					nawait = append(nawait, revdep)
				} else if d.processBlock(d.pending[revdep]) {
					processed = append(processed, revdep)
					seen = map[byzco.BlockID]bool{}
				} else {
					nawait = append(nawait, revdep)
				}
			}
			if len(nawait) > 0 {
				d.await[next] = nawait
			} else {
				delete(d.await, next)
			}
			if first {
				first = false
			} else {
				delete(d.pending, next)
			}
		}
	}
}

func (d *depgraph) processBlock(block *blockData) bool {
	// Skip full processing of any blocks that have already been included into
	// one of our blocks.
	if d.isIncluded(block.id) {
		return true
	}
	// Check if all the referenced blocks have been included already.
	var deps []byzco.BlockID
	for _, ref := range block.links {
		if ref.NodeID == d.self {
			continue
		}
		if !d.isIncluded(ref) {
			log.Debug("Missing dependency", fld.BlockID(block.id), log.String("dep", ref.String()))
			deps = append(deps, ref)
		}
	}
	// Mark the block as pending if any of the referenced blocks, including the
	// previous block, haven't been included.
	if len(deps) > 0 {
		d.addPending(block, deps)
		return false
	}
	// Queue the block for inclusion and mark it as included.
	d.out <- block
	d.markIncluded(block.id)
	return true
}

func (d *depgraph) release() {
	<-d.ctx.Done()
	d.cond.Signal()
}
