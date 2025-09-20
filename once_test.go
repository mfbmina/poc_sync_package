package main

import (
	"testing"
)

func TestDoSomething(t *testing.T) {
	result := doSomething()

	if result != 1 {
		t.Errorf("Expected result to be 1, but got %d", result)
	}
}
