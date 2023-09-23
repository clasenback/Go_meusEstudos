package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println(k, m[k], v)
	}
}

/*

● use a for range loop to print each value and the key associated with each value
*/
