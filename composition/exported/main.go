package main

import (
	"fmt"
	"log"
	"myapp/staff"
)

var Employees = []staff.Employee{
	{FirstName: "Mohit ", LastName: "Sood", Salary: 100000, FullTime: true},
	{FirstName: "Sahil", LastName: "Rana", Salary: 40000, FullTime: true},
	{FirstName: "Sahil ", LastName: "Thakur", Salary: 50000, FullTime: false},
	{FirstName: "Chirag ", LastName: "Sharma", Salary: 65000, FullTime: false},
	{FirstName: "Vibhor", LastName: "Gautam", Salary: 10000, FullTime: true},
}

func main() {
	fmt.Println("I am learning  exported composition")

	myStaff := staff.Office{
		AllStaff: Employees,
	}

	// myStaff.All();
	// log.Println(myStaff.All())
	// overriding value of overpaidlimit
	staff.OverPaidLimit = 60000
	fmt.Println("Overpaid employeed are:")
	log.Println(myStaff.Overpaid())
	fmt.Println("Underpaid employeed are:")
	log.Println(myStaff.UnderPaid())
	log.Println(myStaff.NotVisible())
}
