package mymath

import (
	"testing"
)

/* func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{0, 1, 2, 3, 4, 5, 6})
	}
} */

/* func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{0, 1, 2, 3, 4, 5, 6}))
	//Output:
	//3
} */

func TestCenteredAvg(t *testing.T) {
	type teste struct {
		dados    []int
		esperado float64
	}

	a := teste{[]int{1, 4, 6, 8, 100}, 6.0}
	b := teste{[]int{0, 8, 10, 1000}, 9.0}
	c := teste{[]int{9000, 4, 10, 8, 6, 12}, 9.0}
	d := teste{[]int{123, 744, 140, 200}, 170.0}
	testes := []teste{a, b, c, d}

	for _, test := range testes {
		got := CenteredAvg(test.dados)
		if got != test.esperado {
			t.Error("Got:", got, "Expected:", test.esperado)
		}
	}
}

/* 	a := []int{1, 4, 6, 8, 100}
b := []int{0, 8, 10, 1000}
c := []int{9000, 4, 10, 8, 6, 12}
d := []int{123, 744, 140, 200} */
/* t.Run("A=1", func(t *testing.B) {
	expected := 6.0
	got := CenteredAvg(a)
	if got != expected {
		t.Error("Got:", got, "Expected:", expected)
	}
})
t.Run("B=1", func(t *testing.B) {
	expected := 9.0
	got := CenteredAvg(b)
	if got != expected {
		t.Error("Got:", got, "Expected:", expected)
	}
})
t.Run("C=1", func(t *testing.B) {
	expected := 9.0
	got := CenteredAvg(c)
	if got != expected {
		t.Error("Got:", got, "Expected:", expected)
	}
})
t.Run("D=1", func(t *testing.B) {
	expected := 170.0
	got := CenteredAvg(d)
	if got != expected {
		t.Error("Got:", got, "Expected:", expected)
	}
}) */
