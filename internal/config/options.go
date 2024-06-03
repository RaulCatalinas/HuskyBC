package config

import (
	"fmt"
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/handlers"
)

func printOptions() {
	fmt.Print("Usage: HuskyBC [options]\n\n")
	fmt.Print("Command line for easy Husky configuration\n\n")
	fmt.Println("Options:")
	fmt.Printf("%s \t%s \tOutput the version number\n", constants.OPTION_VERSION, constants.OPTION_VERSION_SORT)
	fmt.Printf("%s \t%s \tOpen GitHub repository for collaboration\n", constants.OPTION_COLLABORATE, constants.OPTION_COLLABORATE_SORT)
	fmt.Printf("%s \t%s \tStart Husky's configuration\n", constants.OPTION_BUILD, constants.OPTION_BUILD_SORT)
	fmt.Printf("%s  \t%s \tDisplay help for command\n\n", constants.OPTION_HELP, constants.OPTION_HELP_SORT)
}

func ConfigureOptions() {
	if len(os.Args) != 2 {
		printOptions()
		os.Exit(0)
	}
	switch os.Args[1] {
	case constants.OPTION_COLLABORATE_SORT, constants.OPTION_COLLABORATE:
		handlers.HandlerOptionCollaborate()
	case constants.OPTION_BUILD_SORT, constants.OPTION_BUILD:
		handlers.HandlerOptionBuild()
	case constants.OPTION_VERSION_SORT, constants.OPTION_VERSION:
		fmt.Println("HuskyBC ->", constants.VERSION)
	case constants.OPTION_HELP_SORT, constants.OPTION_HELP:
		printOptions()
	default:
		printOptions()
	}
}
