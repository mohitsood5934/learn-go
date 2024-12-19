package main

import "fmt"

type Animal struct {
	Name string
	Sound string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func main() {
	var dog Animal
	dog.Name = "Bruno"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name: "Mithi",
		Sound: "meowwww",
		NumberOfLegs: 4,
	}
}
// func main() {
// 	a := 4
// 	b := 5
// 	output := addTwoNumbers(a, b)
// 	manySum := sumMany(1, 2, 3, 4, 5, 6)
// 	fmt.Printf("Sum of numbers is %d\n", output)
// 	fmt.Printf("Sum of many numbers is %d\n", manySum)
// }

func addTwoNumbers(x int, y int) (sum int) {
	sum = x + y
	return sum
}

func sumMany(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
