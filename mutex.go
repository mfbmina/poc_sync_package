package main

import "sync"

func count() int {
	counter := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			mu.Lock()
			counter++
			mu.Unlock()
		})
	}

	wg.Wait()
	return counter
}
