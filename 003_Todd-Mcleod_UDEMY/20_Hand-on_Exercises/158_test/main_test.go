package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(5, 5)
	if total != 0 {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestDoMath(t *testing.T) {
	total := doMath(5, 5, add)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
