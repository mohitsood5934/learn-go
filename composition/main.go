package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

// building/struct the car and embedding the vehicle in it this is what we refer to as composition
type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers", v.NumberOfPassengers)
	fmt.Println("Number of wheels", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("IsElectric:", c.IsElectric)
	fmt.Println("IsHybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {

	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2011,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoXC90.show()

}
