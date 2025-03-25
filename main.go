package main

import "github.com/RaulCatalinas/HuskyBC/internal/config"

func main() {
	optionsManager := config.NewOptionManager()
	optionsManager.ExecuteSelectedOption()
}
