package main

import "sync"

var numbers []int
var mu sync.RWMutex

func store(x int) {
	mu.Lock()
	numbers = append(numbers, x)
	mu.Unlock()
}

func avg() float64 {
	mu.RLock()
	defer mu.RUnlock()

	size := len(numbers)
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return float64(sum) / float64(size)
}
