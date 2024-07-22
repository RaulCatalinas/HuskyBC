package types

type PackageManager string

const (
	NPM  PackageManager = "npm"
	YARN PackageManager = "yarn"
	PNPM PackageManager = "pnpm"
	BUN  PackageManager = "bun"
)
