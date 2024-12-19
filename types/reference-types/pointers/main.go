package main

import "fmt"

func main() {
	var myInt int
	myInt = 10
	fmt.Println(myInt)

	x := 10
	myFirstPointer := &x;
	fmt.Println("Value of x is:", x)
	fmt.Println("My first pointer is", myFirstPointer)

	*myFirstPointer = 15;
	fmt.Println("changing the value of the x by pointing to memory")
	fmt.Println("Value of x is:", x)

}