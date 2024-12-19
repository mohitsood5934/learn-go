package main

import "log"

// basic types (numbers, strings, booleans
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64

var myUnit uint

var myFloat float32
var myFloat64 float64
// aggregate types (array , struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	myInt = 10
	myUnit = 20
	myFloat = 10.2
	myFloat64 = 110.233
	log.Println(myInt, myUnit, myFloat, myFloat64)

	myString := ""
	log.Println(myString)

	var myBool = true
	myBool = false
	log.Println("My boolean value is", myBool)

}

