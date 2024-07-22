package constants

import "github.com/RaulCatalinas/HuskyBC/internal/types"

var LINT_STAGED_CONFIG = map[types.PackageManager]string{
	"npm":  "npx lint-staged",
	"yarn": "yarn dlx lint-staged",
	"pnpm": "pnpm dlx lint-staged",
	"bun":  "bunx lint-staged",
}
