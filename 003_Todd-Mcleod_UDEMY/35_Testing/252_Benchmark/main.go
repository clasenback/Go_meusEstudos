package main

import (
	"00_Estudos/003_Todd-Mcleod_UDEMY/35_Testing/252_Benchmark/mystr"
	"fmt"
	"strings"
)

const s = "Nunc eu felis ac erat maximus eleifend nec non urna. In pellentesque ligula vitae ex lobortis, vitae sodales nisl viverra. Curabitur congue molestie odio, id hendrerit tortor euismod sit amet. Phasellus vitae nunc a nunc interdum imperdiet sed vitae augue. In eu urna diam. Nunc mattis ex velit, eu imperdiet nibh blandit quis. Suspendisse euismod in eros et scelerisque. Nunc vel efficitur urna. Nunc porttitor ex in condimentum vehicula."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%v\n", mystr.Cat(xs))
	fmt.Printf("\n%v\n\n", mystr.Join(xs))
}
