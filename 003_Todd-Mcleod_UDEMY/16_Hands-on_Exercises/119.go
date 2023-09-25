package main

import "fmt"

func main() {
	am := make(map[string][]string)
	am[`bond_james`] = []string{"shaken, not stirred", "martinis", "fast cars"}
	am[`moneypenny_jenny`] = []string{"intelligence", "literature", "computer science"}
	am[`no_dr`] = []string{"cats", "ice cream", "sunsets"}
	am[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	for k, _ := range am {
		fmt.Println(k)
		for i, v := range am[k] {
			fmt.Println("\t", i, v)
		}
	}

	delete(am, "moneypenny_jenny")
	for k, v := range am {
		fmt.Println(k, v)
	}
}

/*
Create a map with a key of TYPE string which is a person’s “last_first” name
and a value of TYPE []string which stores their favorite things.
Store three records in your map.
Print out all of the values, along with their index position in the slice.
key 				value
`bond_james` 		`shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` 	`intelligence`, `literature`, `computer science`
`no_dr` 			`cats`, `ice cream`, `sunsets`

Using the code from the previous example, add a record to your map.
Now print the map out using the “range” loop key value
`fleming_ian` 		`steaks`, `cigars`, `espionage`

Using the code from the previous example, delete a record from your map.
Now print the map out using the “range” loop
*/
