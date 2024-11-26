package enums

type PackageManager string

const (
	PackageManagerNpm  PackageManager = "npm"
	PackageManagerYarn PackageManager = "yarn"
	PackageManagerPnpm PackageManager = "pnpm"
	PackageManagerBun  PackageManager = "bun"
)
