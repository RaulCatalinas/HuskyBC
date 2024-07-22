package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
)

type InstallProps struct {
	PackageManagerToUse types.PackageManager
	PackagesToInstall   []string
}

func promiseSpawn(command string, args []string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func InstallDependencies(props InstallProps) {
	defer func() {
		if r := recover(); r != nil {
			WriteMessage("error", GetErrorMessage("Dependencies"))
			os.Exit(1)
		}
	}()

	installationCommand, exists := constants.INSTALLATION_COMMANDS[types.PackageManager(props.PackageManagerToUse)]
	if !exists {
		WriteMessage("error", "Invalid package manager")
		os.Exit(1)
	}

	WriteMessage("info", fmt.Sprintf("Installing dependencies using: %s...", props.PackageManagerToUse))

	/* FIXME: Problemas al ejecutan bun, dice esto:
	error: Module not found "...\HuskyBC\node_modules\husky\bin.mjs"
	error: prepare script from "" exited with 1

	Investigacion:
		con bun no se instala todo lo que necesita husky, viendo con pnpm y npm, faltan varios archivos,
		hice una prueba haciendo "bun add husky" directamente y es el mismo caso, lo que quiere decir que HuskyBC no presenta problemas
		el problema viene desde el mismo bun
	*/
	args := append([]string{installationCommand}, append(props.PackagesToInstall, "-D")...)
	err := promiseSpawn(string(props.PackageManagerToUse), args)
	if err != nil {
		WriteMessage("error", GetErrorMessage(err.Error()))
		os.Exit(1)
	}

	WriteMessage("success", "Dependencies installed successfully")
}
