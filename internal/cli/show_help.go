package cli

import (
	"fmt"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
)

func ShowHelp() {
	fmt.Print("Usage: huskybc [options]\n\n")
	fmt.Print("Command line for easy Husky configuration\n\n")
	fmt.Println("Options:")

	for _, option := range constants.Options {
		fmt.Printf("%-15s %-5s %s\n", option.Name, option.Alias, option.Description)
	}

	fmt.Println()
}
