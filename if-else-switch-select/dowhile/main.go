package main

import "fmt"

func main() {

	counter := 0

	for {
		// do some work
		fmt.Println("Value of counter is", counter)
		counter++

		// check condition & break
		if counter > 5 {
			break
		}
	}
}
