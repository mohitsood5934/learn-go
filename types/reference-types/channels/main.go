package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// rune is a single character that is used to build a string

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key / q to quit")
	_ = keyboard.Open()

	defer func() {
		fmt.Println("Closing the keyboard after executing main/if any error occurs in the program!!!")
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		keyPressChan <- char
	}
}

func listenForKeyPress() {

	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}

}

// go get github.com/eiannone/keyboard
