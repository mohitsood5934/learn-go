package packageone

import "fmt"

var PackageVar = "My exported package variable"


func PrintMe() {
	fmt.Print(PackageVar);
}