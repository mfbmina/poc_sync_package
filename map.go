package main

import (
	"sync"
)

func mapExample() int {
	var m sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			m.LoadOrStore(i, i*i)
		})
	}

	wg.Wait()

	v, _ := m.Load(0)
	return v.(int)
}
