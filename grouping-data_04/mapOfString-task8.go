package main

import (
	"fmt"
)

func main() {
	// Create a map with a key of TYPE string which is a person’s “last_first” name,
	// and a value of TYPE []string which stores their favorite things. Store three
	// records in your map. Print out all of
	// the values, along with their index position in the slice.
	// `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	// `no_dr`, `Being evil`, `Ice cream`, `Sunsets

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Println("the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	// for k, v := range m {
	// 	fmt.Printf("the key: %v \t  holds the value: %v\n", k, v)
	// }

	fmt.Println("Hello World!")
}

