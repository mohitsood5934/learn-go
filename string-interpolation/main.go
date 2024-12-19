package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User

	user.UserName = readString("What is your name ?")
	user.Age = readInt("How old are you ?")
	user.FavouriteNumber = readFloat("What is your favourite number ?")
	user.OwnsADog = readBool("Do you own a dog (y/n)")

	fmt.Printf("Your name is %s. You are %d years old.Your favourite number if %.2f.You have a dog %t\n", user.UserName, user.Age, user.FavouriteNumber, user.OwnsADog)
}

func prompt() {
	fmt.Print(" ->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}

}

func readBool(s string) bool {
    for {
        fmt.Println(s)
        prompt()

        userInput, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input. Please try again.")
            continue
        }

        userInput = strings.TrimSpace(userInput)

        if userInput == "y" {
            return true
        } else if userInput == "n" {
            return false
        } else {
            fmt.Println("Please enter answer in y/n format")
        }
    }
}


