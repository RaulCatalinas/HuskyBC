package utils

import (
	"os"
	"os/exec"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
)

type InstallProps struct {
	PackageManagerToUse enums.PackageManager
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
			WriteMessage(WriteMessageProps{
				Type:    enums.MessageTypeError,
				Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.DependenciesError],
			})

			os.Exit(1)
		}
	}()

	existsDevDepsSection := existsSection(
		constants.PATH_PACKAGE_JSON,
		constants.DEV_DEPS_SECTION,
	)

	if !existsDevDepsSection {
		createEmptySection(constants.PATH_PACKAGE_JSON, constants.DEV_DEPS_SECTION)
	}

	message := "Installing dependencies using: " + props.PackageManagerToUse + "..."

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: string(message),
	})

	/* FIXME: Problemas al ejecutan bun, dice esto:
	error: Module not found "...\HuskyBC\node_modules\husky\bin.mjs"
	error: prepare script from "" exited with 1

	Investigacion:
		con bun no se instala todo lo que necesita husky, viendo con pnpm y npm, faltan varios archivos,
		hice una prueba haciendo "bun add husky" directamente y es el mismo caso, lo que quiere decir que HuskyBC no presenta problemas
		el problema viene desde el mismo bun
	*/

	installationCommand := constants.INSTALLATION_COMMANDS[props.PackageManagerToUse]

	args := append(
		[]string{installationCommand},
		append(
			props.PackagesToInstall,
			constants.DEV_DEPS_INSTALLATION_PARAM,
		)...,
	)

	err := promiseSpawn(string(props.PackageManagerToUse), args)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.DependenciesError],
		})

		os.Exit(1)
	}

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "Dependencies installed successfully",
	})
}
