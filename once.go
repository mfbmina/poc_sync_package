package main

import "sync"

func doSomething() int {
	wg := sync.WaitGroup{}
	o := sync.Once{}
	result := 0

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			o.Do(func() {
				result++
			})
		})
	}

	wg.Wait()
	return result
}
