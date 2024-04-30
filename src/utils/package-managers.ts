// Constants
import { PACKAGE_MANGERS } from '@/constants/package-mangers'

// Third-Party libraries
import inquirer from 'inquirer'

// Types
import type { PackageManager } from '@/types/package-manger'

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
    console.error(error)
    throw new Error('Something went wrong, try again later.')
  }
}
