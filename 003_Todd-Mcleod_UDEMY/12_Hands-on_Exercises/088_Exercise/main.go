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

	age := m["James"]
	fmt.Println(age)

	tentativa := m["Q"]
	fmt.Println(tentativa)

	if v, ok := m["Q"]; !ok {
		fmt.Printf("%v \t %T \t %v \n", v, ok, ok)
		fmt.Println("Q não é uma chave deste mapa.")
	}
}

/*
● use the code from the previous exercise
● add this code to print a single value stored in the map
	age := m["James"]
	fmt.Println(age)
● now use similar code to use the lookup of "Q" and print that value
● now use the "comma ok" idiom to test whether "Q" is actually stored in the map, then
print out a statement if it is not stored in the map
	○ hint: check effective go for the "comma ok" idiom
*/
