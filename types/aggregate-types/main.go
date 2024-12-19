package main

import "fmt"

// aggregate types (array, struct)
// reference types (pointers, slices, maps, functions, channels)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"
	for i := 0; i < len(myStrings); i++ {
		fmt.Println(myStrings[i])
	}
	// fmt.Println(myStrings)
	var myCar Car
	myCar.NumberOfTires = 4
	myCar.Luxury = true
	myCar.BucketSeats = true
	myCar.Make = "Hyundai"
	myCar.Model = "sx"
	myCar.Year = 2021

	// myCar := Car {
	// 	NumberOfTires: 4,
	// 	Model: "sx",
	// 	Make: "Hyundai",
	// }
	fmt.Printf("My car is %s.It's model is %s.It is manufactured in the year %d", myCar.Make, myCar.Model, myCar.Year)

}
