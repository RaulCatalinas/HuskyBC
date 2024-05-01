// Types
import type { PackageManager } from '@/types/package-manger'

// Constants
import { INSTALLATION_COMMANDS } from '@/constants/dependencies'

// Third-Party libraries
import promiseSpawn from '@npmcli/promise-spawn'

interface Props {
  packageManagerToUse: PackageManager
  packagesToInstall: string | string[]
}

export async function installDependencies({
  packageManagerToUse,
  packagesToInstall
}: Props) {
  const installationCommand = INSTALLATION_COMMANDS[packageManagerToUse]

  console.log(`Installing dependencies using: ${packageManagerToUse}...`)

  await promiseSpawn(
    packageManagerToUse,
    [installationCommand, packagesToInstall, '-D'].flat()
  )

  console.log('Installed dependencies')
}
