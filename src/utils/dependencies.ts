// Types
import type { PackageManager } from '@/types/package-manger'

// Constants
import { INSTALLATION_COMMANDS } from '@/constants/dependencies'

// Third-Party libraries
import promiseSpawn from '@npmcli/promise-spawn'

// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'

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

    writeMessage({
      type: 'info',
      message: `Installing dependencies using: ${packageManagerToUse}...`
    })

    await promiseSpawn(
      packageManagerToUse,
      [installationCommand, packagesToInstall, '-D'].flat()
    )

    writeMessage({
      type: 'success',
      message: 'Dependencies installed successfully'
    })
  } catch (error) {
    writeMessage({
      type: 'error',
      message: getErrorMessage('Dependencies')
    })
    process.exit(1)
  }
}
