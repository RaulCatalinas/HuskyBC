// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'

// Third-Party libraries
import { exists as existsFolderOrFile } from 'fs-extra'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'

export async function exists(path: string) {
  try {
    return await existsFolderOrFile(path)
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('CheckFileExists')
    })

    process.exit(1)
  }
}

export async function createFolder(name: string) {
  try {
    await fs.mkdir(`${process.cwd()}/${name}`)
  } catch {
    const errorMessage = getErrorMessage('CreateFolder')

    writeMessage({
      type: 'error',
      message: errorMessage.replace('{folderName}', name)
    })
    process.exit(1)
  }
}

export async function createEmptyFile(fileName: string) {
  try {
    await fs.writeFile(`${process.cwd()}/${fileName}`, '', {
      encoding: UTF8_ENCODING
    })
  } catch {
    const errorMessage = getErrorMessage('EmptyFileCreate')

    writeMessage({
      type: 'error',
      message: errorMessage.replace('{fileName}', fileName)
    })
    process.exit(1)
  }
}
