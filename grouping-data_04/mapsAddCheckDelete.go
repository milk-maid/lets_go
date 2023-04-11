package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Bello": 35,
		"Oscar": 43,
	}
	fmt.Println(m)
	fmt.Println(m["Bello"])
	fmt.Println(m["Barnabas"])
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	//CHECK IF THE ELEMENT EXIST
	//this won't print since its False
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the IF print!", v)
	}

	// this will print because the keyalue exist  .'. True
	if v, ok := m["Oscar"]; ok {
		fmt.Println("This is the IF print!", v)
	}

	//ADD AN ELEMENT
	m["todd"] = 323

	//Print the elements in k(keys) & v(value) pair out
	for k, v := range m {
		fmt.Println(k, v)
	}

  // DELETE AN ENTRY
  if v, ok := m["Bello"]; ok {
    fmt.Println("value: ", v)
    delete(m, "Bello")
  }

  fmt.Println(m)
}

