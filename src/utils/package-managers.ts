// Constants
import { ErrorMessages } from '@/constants/errors'
import { PACKAGE_MANGERS } from '@/constants/package-mangers'

// Third-Party libraries
import inquirer from 'inquirer'

// Types
import type { PackageManager } from '@/types/package-manger'

// NodeJS
import process from 'node:process'

export async function getPackageManger(): Promise<PackageManager> {
  try {
    const { packageManager } = await inquirer.prompt({
      type: 'rawlist',
      choices: PACKAGE_MANGERS.filter(
        packageManager => packageManager !== 'yarn'
      ),
      message: 'Which package manager do you wanna use?',
      name: 'packageManager'
    })

    return packageManager
  } catch (error) {
    console.error(ErrorMessages.Default)
    process.exit(1)
  }
}

await getPackageManger()
