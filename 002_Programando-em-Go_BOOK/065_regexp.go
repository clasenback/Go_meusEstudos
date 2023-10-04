package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	texto := "alex tem 33 anos."
	expr := regexp.MustCompile("\\d")
	fmt.Println(expr.ReplaceAllString(texto, "4"))

	expr = regexp.MustCompile("\\b\\w")
	funcao := func(s string) string { return strings.ToUpper(s) }
	processado := expr.ReplaceAllStringFunc(texto, funcao)
	fmt.Println(processado)
}
