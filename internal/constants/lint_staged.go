package constants

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var LINT_STAGED_CONFIG = map[enums.PackageManager]string{
	enums.PackageManagerNpm:  "npx lint-staged",
	enums.PackageManagerYarn: "yarn dlx lint-staged",
	enums.PackageManagerPnpm: "pnpm dlx lint-staged",
	enums.PackageManagerBun:  "bunx lint-staged",
}
