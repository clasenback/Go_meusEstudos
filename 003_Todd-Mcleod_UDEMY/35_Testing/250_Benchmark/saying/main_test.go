package saying

import (
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Alex")
	expected := "Bem vindo, Alex"
	if s != expected {
		t.Error("got:", s, "expected:", expected)
	}
}

/* func ExampleGreet() {
	fmt.Println(Greet("Alex"))
	// Output:
	// "Bem vindo, Alex"
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Alex")
	}
}
*/

/*
go test -coverprofile c.out
go tool cover -html=c.out
*/
