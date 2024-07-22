package userinput

import (
	"fmt"
	"log"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
)

func PrintListPackageManeger() {
	for index, packageManager := range constants.PACKAGEMANAGERS {
		if _, err := fmt.Printf("%d. %s\n", index, packageManager); err != nil {
			log.Fatalln(err)
		}
	}
}

func GetPackageManager() string {
	var selectedPackageManagerNumber int

	if _, err := fmt.Println("Which package manager do you wanna use?"); err != nil {
		log.Fatalln(err)
	}

	PrintListPackageManeger()

	if _, err := fmt.Print("Answer: "); err != nil {
		log.Fatalln(err)
	}

	if _, err := fmt.Scan(&selectedPackageManagerNumber); err != nil {
		log.Fatalln(err)
	}

	if len(constants.PACKAGEMANAGERS) <= selectedPackageManagerNumber {
		if _, err := fmt.Println(">> Please enter a valid index"); err != nil {
			log.Fatalln(err)
		}

		return GetPackageManager()
	}

	selectedPackageManager := constants.PACKAGEMANAGERS[selectedPackageManagerNumber]

	return selectedPackageManager
}
