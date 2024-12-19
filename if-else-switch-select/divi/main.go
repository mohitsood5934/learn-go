package main

import "fmt"

func main() {
	count := 0
	divNumberCheck := 7

	for i := 0; i <= 100; i++ {
		if i%divNumberCheck == 0 {
			count += 1
			fmt.Println(i, "is divisible by 7")

		}
	}

	// for i := 0; i <= 100 ; i = i + 7 {
	// 	fmt.Println(i, "is divisible by 7")
	// }

	fmt.Sprintf("There are total %d numbers divisible by %d", divNumberCheck, divNumberCheck)

}
