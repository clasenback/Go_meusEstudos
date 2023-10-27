package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}

func TestMySub(t *testing.T) {
	x := mySub(2, 3)
	if x != -5 {
		t.Error("Expected", -5, "Got", x)
	}
}
