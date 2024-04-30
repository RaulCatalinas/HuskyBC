// Types
import type { PackageManager } from '@/types/package-manger'

// Constants
import { INSTALLATION_COMMANDS } from '@/constants/dependencies'

interface Props {
  packageManagerToUse: PackageManager
  packagesToInstall: string | string[]
}

export function installDependencies({
  packageManagerToUse,
  packagesToInstall
}: Props) {
  const packageManager = INSTALLATION_COMMANDS[packageManagerToUse]

  console.log(`Installing dependencies using: ${packageManager}...`)

  // TODO: Implement the logic for installing dependencies using the package manager of the user's choice.

  console.log('Installed dependencies')
}
