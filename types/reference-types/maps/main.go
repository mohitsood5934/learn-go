package main

import "fmt"

func main() {
	fmt.Println("Implementing maps in Go")
	var intMap = make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	intMap["six"] = 6

	delete(intMap, "four")
	// for key, value := range intMap {
	// 	fmt.Printf("Key is %s and Value is %d\n", key, value)
	// }

	// checking if element exist in map
	el , ok :=  intMap["four"]
	if ok {
		fmt.Println(el , "is present in map")
	} else {
		fmt.Println(el , "is not present in map")
	}
	// modifying value in map
	intMap["two"] = 4
}

// Map is a refrence type & it is always needed to pass by reference. So you do not need a pointer
// Maps are not sorted