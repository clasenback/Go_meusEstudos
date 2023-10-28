package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	expected := 49
	got := Years(7)
	if got != expected {
		t.Error("got:", got, "expected:", expected)
	}
}

func TestYearsTwo(t *testing.T) {
	expected := 49
	got := YearsTwo(7)
	if got != expected {
		t.Error("got:", got, "expected:", expected)
	}
}

func BenchmarkYears(b *testing.B) {
	xs := 7
	for i := 0; i < b.N; i++ {
		Years(xs)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	xs := 7
	for i := 0; i < b.N; i++ {
		YearsTwo(xs)
	}
}

func ExampleYears() {
	hy := 7
	dy := Years(hy)
	fmt.Println(dy)
	// Output:
	// 49
}

func ExampleYearsTwo() {
	hy := 7
	dy := Years(hy)
	fmt.Println(dy)
	// Output:
	// 49
}
