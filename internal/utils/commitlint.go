package utils

import (
	"os"
	"sync"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
)

type CommitlintProps struct {
	PackageManagerToUse types.PackageManager
	PackageJsonPath     string
	ShouldPublishToNpm  bool
}

func GenerateCommitlintConfig(commitlintProps CommitlintProps) {
	defer func() {
		if r := recover(); r != nil {
			WriteMessage(WriteMessageProps{
				Type:    "error",
				Message: GetErrorMessage("Commitlint"),
			})

			os.Exit(1)
		}
	}()

	WriteMessage(WriteMessageProps{
		Type:    "config",
		Message: "Configuring commitlint...",
	})

	InstallDependencies(InstallProps{
		PackageManagerToUse: commitlintProps.PackageManagerToUse,
		PackagesToInstall: []string{
			"lint-staged",
			"@commitlint/cli",
			"@commitlint/config-conventional",
		},
	})

	addScript(addScriptProps{
		PackageJsonPath: commitlintProps.PackageJsonPath,
		ScriptsToAdd: []packageJsonScript{
			{Key: "lint", Value: "eslint src"},
			{Key: "lint:fix", Value: "eslint src --fix"},
			{Key: "format", Value: "prettier src --check"},
			{Key: "format:write", Value: "prettier src --write"},
		},
	})

	createCommitlintConfigFiles(commitlintProps.PackageManagerToUse)

	if commitlintProps.ShouldPublishToNpm {
		modifyNpmIgnore([]string{".lintstagedrc", "commitlint.config.js"})
	}

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "commitlint's configuration generated successfully",
	})
}

func createCommitlintConfigFiles(packageManagerToUse types.PackageManager) {
	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: "Creating configuration files...",
	})

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		writeFile(
			".husky/commit-msg",
			[]byte(constants.COMMITLINT_CONFIG[string(packageManagerToUse)]),
		)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		content := "export default { extends: ['@commitlint/config-conventional'] }"

		writeFile("commitlint.config.js", []byte(content))
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		content := `{
			"src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}": [
				"prettier src --check",
				"eslint src --max-warnings 0"
			]
		}`

		writeFile(".lintstagedrc", []byte(content))
	}()

	wg.Wait()

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Configuration files (commit-msg, commitlint.config.js and .lintstagedrc) created successfully",
	})
}
