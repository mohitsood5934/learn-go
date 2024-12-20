package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println("Enter quit to exit the game")
	fmt.Println(whatToSay)

	// userInput, _ := reader.ReadString('\n')

	// fmt.Println(userInput)

	for {
		fmt.Print(" ->")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break

		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}
}
