package main

import "fmt"

func main() {
	apples := 18
	oranges := 9

	//expressions
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// if apples == "18" {
	// 	types should match when we are comparing values

	// }
	fmt.Printf("%d > %d : %t\n", apples,oranges, apples > oranges)
	fmt.Printf("%d < %d : %t\n", apples,oranges, apples < oranges)
	fmt.Printf("%d >= %d : %t\n", apples,oranges, apples >= oranges)
	fmt.Printf("%d <= %d : %t\n", apples,oranges, apples <= oranges)
}
