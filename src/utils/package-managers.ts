// Constants
import { PACKAGE_MANGERS } from '@/constants/package-mangers'

// Third-Party libraries
import inquirer from 'inquirer'

export async function getPackageManger(): Promise<string> {
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
