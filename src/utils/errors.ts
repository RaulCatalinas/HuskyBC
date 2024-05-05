import { ERROR_MESSAGES } from '@/constants/errors'

export function getErrorMessage(error: keyof typeof ERROR_MESSAGES) {
  return ERROR_MESSAGES[error]
}
