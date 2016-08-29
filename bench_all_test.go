package drwmutex

import (
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkStdLibrary(b *testing.B) {
	var cnt int
	var mu sync.RWMutex

	b.SetParallelism(10)

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		limit := rand.Intn(1000) + 1000
		for pb.Next() {
			i++
			if i == limit {
				i = 0
				// write
				mu.Lock()
				cnt++
				mu.Unlock()
				continue
			}
			mu.RLock()
			_ = cnt
			mu.RUnlock()
		}
	})
}

func BenchmarkPrwMutex(b *testing.B) {
	var cnt int
	mu := NewRWMutex()

	b.SetParallelism(10)

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		limit := rand.Intn(1000) + 1000
		for pb.Next() {
			i++
			if i == limit {
				i = 0
				// write
				mu.Lock()
				cnt++
				mu.Unlock()
				continue
			}
			locker := mu.RLock()
			_ = cnt
			locker.RUnlock()
		}
	})

}
