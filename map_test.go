package main

import (
	"testing"
)

func TestMapExample(t *testing.T) {
	result := mapExample()

	if result != 0 {
		t.Errorf("Expected return to be 0, but got %d", result)
	}
}
