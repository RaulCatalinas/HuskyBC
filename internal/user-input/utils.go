package userinput

import (
	"fmt"
	"log"
)

func ConfirmationHandler(message string) bool {
	var response string

	if _, err := fmt.Print(message); err != nil {
		log.Fatalln(err)
	}

	if _, err := fmt.Scan(&response); err != nil {
		log.Fatalln(err)
	}

	if response != "y" && response != "n" {
		fmt.Println(">> Please enter a valid option")
		return ConfirmationHandler(message)
	}

	if response == "y" {
		return true
	}

	return false
}
