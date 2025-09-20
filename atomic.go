package main

import (
	"sync"
	"sync/atomic"
)

func countWithAtomic() int {
	var counter atomic.Int32
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			counter.Add(1)
		})
	}

	wg.Wait()
	return int(counter.Load())
}
