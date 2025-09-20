package main

import (
	"testing"
)

func TestCountWithAtomic(t *testing.T) {
	result := countWithAtomic()

	if result != 1000 {
		t.Errorf("Expected counter to be 1000, but got %d", result)
	}
}
