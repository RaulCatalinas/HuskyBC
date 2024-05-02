// Third-Party libraries
import inquirer from 'inquirer'

// Constants
import { ErrorMessages } from '@/constants/errors'

// NodeJS
import process from 'node:process'

export async function addCommitlint(): Promise<boolean> {
  try {
    const { addCommitlint } = await inquirer.prompt({
      type: 'confirm',
      default: true,
      name: 'addCommitlint',
      message: 'Do you wanna add commitlint?:'
    })

    return addCommitlint
  } catch {
    console.error(ErrorMessages.Default)
    process.exit(1)
  }
}
