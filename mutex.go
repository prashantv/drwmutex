package drwmutex

import (
	"runtime"
	"sync"

	"github.com/prashantv/goid"
)

type RWMutex struct {
	shards []sync.RWMutex
}

func NewRWMutex() *RWMutex {
	return &RWMutex{
		shards: make([]sync.RWMutex, runtime.GOMAXPROCS(0)),
	}
}

func (m *RWMutex) Lock() {
	for i := range m.shards {
		m.shards[i].Lock()
	}
}

func (m *RWMutex) Unlock() {
	for i := range m.shards {
		m.shards[i].Unlock()
	}
}

func (m *RWMutex) RLock() *sync.RWMutex {
	shard := &m.shards[goid.ProcID()]
	shard.RLock()
	return shard
}
