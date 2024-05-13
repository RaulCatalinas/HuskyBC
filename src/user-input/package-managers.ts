// Constants
import { PACKAGE_MANGERS } from '@/constants/package-mangers'

// Third-Party libraries
import inquirer from 'inquirer'

// Types
import type { PackageManager } from '@/types/package-manger'

// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from '@/utils/console'
import { getErrorMessage } from '@/utils/errors'

export async function getPackageManger(): Promise<PackageManager> {
  try {
    const { packageManager } = await inquirer.prompt({
      type: 'rawlist',
      choices: PACKAGE_MANGERS,
      message: 'Which package manager do you wanna use?',
      name: 'packageManager'
    })

    return packageManager
  } catch (error) {
    writeMessage({
      type: 'error',
      message: getErrorMessage('PackageManagerSelection')
    })
    process.exit(1)
  }
}
