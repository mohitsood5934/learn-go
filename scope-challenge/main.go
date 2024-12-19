package main

import (
	"fmt"
	"myapp/packageone"
)

var myVar = "My package level variable"

func main() {
	var blockVar = "My block level variable"
	fmt.Print(myVar, blockVar, packageone.PackageVar)

}
