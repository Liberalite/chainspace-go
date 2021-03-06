package conns // import "chainspace.io/chainspace-go/internal/conns"

import (
	"sync"

	"chainspace.io/chainspace-go/internal/crypto/signature"
	"chainspace.io/chainspace-go/network"
	"chainspace.io/chainspace-go/service"
)

type pool struct {
	mu    sync.Mutex
	i     int
	size  int
	conns []*cache
}

type Pool interface {
	MessageAckPending() int
	Close()
	Borrow() Cache
}

func (c *pool) Borrow() Cache {
	c.mu.Lock()
	cc := c.conns[c.i]
	c.i += 1
	if c.i >= c.size {
		c.i = 0
	}
	c.mu.Unlock()
	return cc
}

func (c *pool) Close() {
	for _, conn := range c.conns {
		conn := conn
		conn.Close()
	}
}

func (c *pool) MessageAckPending() int {
	c.mu.Lock()
	cnt := 0
	for _, v := range c.conns {
		v := v
		cnt += len(v.pendingAcks)
	}
	c.mu.Unlock()
	return cnt
}

func NewPool(size int, nodeID uint64, top network.NetTopology, maxPayload int, key signature.KeyPair, connection service.CONNECTION) *pool {
	conns := make([]*cache, 0, size)
	for i := 0; i < size; i += 1 {
		cc := NewCache(
			nodeID, top, maxPayload, key, connection)
		conns = append(conns, cc)
	}
	return &pool{
		i:     0,
		size:  size,
		conns: conns,
	}
}
