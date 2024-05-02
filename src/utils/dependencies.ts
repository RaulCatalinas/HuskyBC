// Types
import type { PackageManager } from '@/types/package-manger'

// Constants
import { INSTALLATION_COMMANDS } from '@/constants/dependencies'
import { ErrorMessages } from '@/constants/errors'

// Third-Party libraries
import promiseSpawn from '@npmcli/promise-spawn'

// NodeJS
import process from 'node:process'

interface Props {
  packageManagerToUse: PackageManager
  packagesToInstall: string | string[]
}

export async function installDependencies({
  packageManagerToUse,
  packagesToInstall
}: Props) {
  try {
    const installationCommand = INSTALLATION_COMMANDS[packageManagerToUse]

    console.log(`Installing dependencies using: ${packageManagerToUse}...`)

    await promiseSpawn(
      packageManagerToUse,
      [installationCommand, packagesToInstall, '-D'].flat()
    )

    console.log('Installed dependencies')
  } catch (error) {
    console.error(ErrorMessages.Dependencies)
    process.exit(1)
  }
}
