package main

import "fmt"

var one = "One"

func main () {
  var one = "This is a block level variable"
	fmt.Println(one)

	myFunc()
}

func myFunc() {
	fmt.Println(one)
}
