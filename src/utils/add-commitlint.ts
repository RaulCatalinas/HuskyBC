// Third-Party libraries
import inquirer from 'inquirer'

// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'

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
    writeMessage({
      type: 'error',
      message: getErrorMessage('CommitlintSelection')
    })
    process.exit(1)
  }
}
