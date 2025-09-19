package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	result := count()

	if result != 1000 {
		t.Errorf("Expected counter to be 1000, but got %d", result)
	}
}
