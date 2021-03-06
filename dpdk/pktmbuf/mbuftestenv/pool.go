package mbuftestenv

import (
	"sync"

	"github.com/usnistgov/ndn-dpdk/dpdk/eal"
	"github.com/usnistgov/ndn-dpdk/dpdk/ealtestenv"
	"github.com/usnistgov/ndn-dpdk/dpdk/pktmbuf"
)

// TestPool adds convenience functions to pktmbuf.Pool for unit testing.
type TestPool struct {
	Template pktmbuf.Template
	poolInit sync.Once
	pool     *pktmbuf.Pool
}

// Pool returns the mempool.
func (p *TestPool) Pool() *pktmbuf.Pool {
	p.poolInit.Do(func() {
		ealtestenv.Init()
		p.pool = p.Template.MakePool(eal.NumaSocket{})
	})
	return p.pool
}

// Alloc allocates a packet.
func (p *TestPool) Alloc() *pktmbuf.Packet {
	vec := p.Pool().MustAlloc(1)
	return vec[0]
}

// TestPool instances.
var (
	Direct   TestPool
	Indirect TestPool
)

func init() {
	Direct.Template = pktmbuf.Direct.Update(pktmbuf.PoolConfig{
		Capacity: 4095,
	})
	Indirect.Template = pktmbuf.Indirect.Update(pktmbuf.PoolConfig{
		Capacity: 4095,
	})
}
