package constants

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var INSTALLATION_COMMANDS = map[enums.PackageManager]string{
	enums.PackageManagerNpm:  "install",
	enums.PackageManagerYarn: "add",
	enums.PackageManagerPnpm: "add",
	enums.PackageManagerBun:  "add",
}
