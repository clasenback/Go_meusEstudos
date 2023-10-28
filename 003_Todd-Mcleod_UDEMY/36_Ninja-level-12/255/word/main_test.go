package word

import (
	"Go_meusEstudos/003_Todd-Mcleod_UDEMY/36_Ninja-level-12/255/quote"
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	expected := 3
	got := Count("Alexander Clasen Back")
	if got != expected {
		t.Error("got:", got, "expected:", expected)
	}
}

func TestUseCount(t *testing.T) {
	expected := make(map[string]int)
	expected["Alexander"] = 1
	expected["Clasen"] = 1
	expected["Back"] = 1
	expected["Abacaxi"] = 3
	got := UseCount("Alexander Abacaxi Clasen Abacaxi Back Abacaxi")
	for k, v := range expected {
		if got[k] != v {
			t.Error("got:", got[k], "expected:", v)
			fmt.Println(got)

		}
	}
}

func BenchmarkCount(b *testing.B) {
	//s := "Alexander Clasen Back"
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	//s := "Alexander Abacaxi Clasen Abacaxi Back Abacaxi"
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

/* func ExampleCount() {
	fmt.Println(Count("Alexander Clasen Back"))
	//Output:
	//3
}

func ExampleUseCount() {
	fmt.Println(Count("quilo carro gordo"))
	//Output:
	//map[carro:1 gordo:1 quilo:1]
} */
