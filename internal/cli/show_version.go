package cli

import (
	"fmt"
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
)

func ShowVersion() {
	fmt.Println("HuskyBC -> " + constants.VERSION)
}