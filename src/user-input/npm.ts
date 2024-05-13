// Third-Party libraries
import inquirer from 'inquirer'

// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from '@/utils/console'
import { getErrorMessage } from '@/utils/errors'

export async function shouldPublishToNPM(): Promise<boolean> {
  try {
    const { shouldPublishToNPM } = await inquirer.prompt({
      type: 'confirm',
      default: false,
      name: 'shouldPublishToNPM',
      message: 'Will you publish it on npm?:'
    })

    return shouldPublishToNPM
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('PublishConfirmation')
    })
    process.exit(1)
  }
}
