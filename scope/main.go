package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	var one = "This is a block level variable"
	fmt.Println(one)

	myFunc()

	newString := packageone.PublicVar
	fmt.Println("New String is", newString)
	packageone.Exported()
}

func myFunc() {
	fmt.Println(one)
}
