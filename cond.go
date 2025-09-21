package main

import (
	"fmt"
	"sync"
	"time"
)

func condExample() {
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	wg := sync.WaitGroup{}
	active := false

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			cond.L.Lock()
			defer cond.L.Unlock()

			for !active {
				cond.Wait()
			}

			fmt.Println("Active is true, printing: ", i)
		})
	}

	// Activate all goroutines after some time
	time.Sleep(time.Second * 5)
	fmt.Println("Setting Active to true...")
	active = true
	fmt.Println("Wake up one goroutine...")
	cond.Signal()

	time.Sleep(time.Second * 5)
	fmt.Println("Wake up all goroutines...")
	cond.Broadcast()

	wg.Wait()
	fmt.Println("All goroutines finished.")
}
