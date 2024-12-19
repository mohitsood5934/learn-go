package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 1000
	// execute a loop while i > 100
	for i > 100 {
		// i = i - 1
		// fmt.Println("i is:", i)

		i = rand.Intn(1000) + 1
		fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("i is ", i, "so loops keep going")
		}
	}

	fmt.Println("Go i as", i, "so breaking out of the loop")
}
