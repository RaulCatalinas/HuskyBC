// Third-Party libraries
import chalk from 'chalk'

// Constants
import type { ERROR_MESSAGES } from '@/constants/errors'

// Types
import type { MessageType } from '@/types/message'

interface Props {
  type: MessageType
  message: typeof ERROR_MESSAGES | string
}

export function writeMessage({ type, message }: Props) {
  if (type === 'success') console.log(chalk.green(message))
  if (type === 'info') console.log(chalk.blue(message))
  if (type === 'error') console.log(chalk.red(message))
  if (type === 'config') console.log(chalk.white(message))
}
