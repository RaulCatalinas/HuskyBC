package constants

var LINT_STAGED_CONFIG = map[string]string{
	"npm":  "npx lint-staged",
	"yarn": "yarn dlx lint-staged",
	"pnpm": "pnpm dlx lint-staged",
	"bun":  "bunx lint-staged",
}
