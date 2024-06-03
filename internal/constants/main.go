package constants

const (
	PATH_PACKAGE_JSON = "package.json"
	PATH_DIR_HUSKY    = ".husky"
	REPOSITORY        = "https://github.com/RaulCatalinas/HuskyBC"
	ISSUES            = REPOSITORY + "/issues"
	VERSION           = "1.0.0"
	NPM               = "npm"
	YARN              = "yarn"
	PNPM              = "pnpm"
	//BUN                     = "bun"
	UTF8_ENCODING           = "UTF-8"
	SPECIAL_CHARS_REGEX     = "[^a-zA-Z0-9._-]"
	OPTION_VERSION          = "--version"
	OPTION_VERSION_SORT     = "-v"
	OPTION_COLLABORATE      = "--collaborate"
	OPTION_COLLABORATE_SORT = "-co"
	OPTION_BUILD            = "--build"
	OPTION_BUILD_SORT       = "-b"
	OPTION_HELP             = "--help"
	OPTION_HELP_SORT        = "-h"
)

var PACKAGEMANAGERS = []string{NPM, YARN, PNPM /* BUN */}

var INSTALLATION_COMMANDS = map[string]string{
	NPM:  "install",
	YARN: "add",
	PNPM: "add",
	/* BUN:  "add", */
}

var LINT_STAGED_CONFIG = map[string]string{
	NPM:  "npx lint-staged",
	YARN: "yarn dlx lint-staged",
	PNPM: "pnpm dlx lint-staged",
	/* BUN:  "bunx lint-staged", */
}

var COMMITLINT_CONFIG = map[string]string{
	"npm":  "#!/bin/sh\n. \"$(dirname \"$0\")/_/husky.sh\"\nnpx --no-install commitlint --edit \"$1\"\n",
	"yarn": "#!/bin/sh\n. \"$(dirname \"$0\")/_/husky.sh\"\nyarn commitlint --edit \"$1\"\n",
}
