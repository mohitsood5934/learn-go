package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cow")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	fmt.Println(animals)
	for i := 0; i < len(animals); i++ {
		println(animals[i])
	}

	// for i, x := range(animals) {
	// 	fmt.Println(i,x)
	// }

	fmt.Println("First element is:", animals[0])
	fmt.Println("First three elements are:", animals[0:3])
	fmt.Println("Is it sorted ?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("String sorted , now check")
	fmt.Println(animals)
	fmt.Println("Is it sorted now ?", sort.StringsAreSorted(animals))

	animals = DeleteFromSlice(animals, 1);
	fmt.Println("Slice after deleting item at index 1", animals)
}

func DeleteFromSlice(a []string,i int) []string {
	a[i] = a[len(a) - 1]
	a[len(a) - 1] = ""
	a = a[:len(a) - 1]
	return a 
}