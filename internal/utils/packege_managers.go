package utils

import "github.com/RaulCatalinas/HuskyBC/internal/constants"

func GetPackageManagersAsStrings() []string {
	packageManagesAsStrings := make([]string, len(constants.PACKAGE_MANAGERS))

	for i, manager := range constants.PACKAGE_MANAGERS {
		packageManagesAsStrings[i] = string(manager)
	}

	return packageManagesAsStrings
}
