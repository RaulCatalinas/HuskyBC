package constants

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var PACKAGE_MANAGERS = []enums.PackageManager{
	enums.PackageManagerNpm,
	enums.PackageManagerYarn,
	enums.PackageManagerPnpm,
	enums.PackageManagerBun,
}
