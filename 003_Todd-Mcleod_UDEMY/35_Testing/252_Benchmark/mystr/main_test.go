package mystr

import "testing"

var xs = []string{"Nunc", "eu", "felis", "ac", "erat", "maximus", "eleifend", "nec", "non", "urna.", "In", "pellentesque", "ligula", "vitae", "ex", "lobortis,", "vitae", "sodales", "nisl", "viverra.", "Curabitur", "congue", "molestie", "odio,", "id", "hendrerit", "tortor", "euismod", "sit", "amet.", "Phasellus", "vitae", "nunc", "a", "nunc", "interdum", "imperdiet", "sed", "vitae", "augue.", "In", "eu", "urna", "diam.", "Nunc", "mattis", "ex", "velit,", "eu", "imperdiet", "nibh", "blandit", "quis.", "Suspendisse", "euismod", "in", "eros", "et", "scelerisque.", "Nunc", "vel", "efficitur", "urna.", "Nunc", "porttitor", "ex", "in", "condimentum", "vehicula."}

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
