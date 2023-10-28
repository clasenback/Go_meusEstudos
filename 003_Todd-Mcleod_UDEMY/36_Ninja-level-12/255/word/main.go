// Package word offers some fun functions to work with strings
package word

import "strings"

//UseCount maps the number of occurencies of each word in a given string. Words are considered separated by a space.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns the number of words in a given string, considering space chacacter as delimiter.
func Count(s string) int {
	// write the code for this func
	xs := strings.Split(s, " ")
	return len(xs)
}
