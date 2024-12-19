package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public (or exported)"

func notExported() {
	fmt.Println("Not exported function")

}

func Exported() {
	fmt.Println("This is a exported function ")
}