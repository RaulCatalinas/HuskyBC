// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { ErrorMessages } from '@/constants/errors'

export async function exists(path: string) {
  try {
    await fs.access(path, fs.constants.F_OK)

    return true
  } catch (err) {
    return false
  }
}

export async function createFolder(name: string) {
  try {
    await fs.mkdir(`${process.cwd()}/${name}`)
  } catch {
    console.error(ErrorMessages.Default)
    process.exit(1)
  }
}
