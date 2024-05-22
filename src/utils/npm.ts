// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'
import { SPECIAL_CHARS_REGEX } from '@/constants/regex'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'
import { createEmptyFile, exists } from './user-os'

export async function modifyNpmIgnore(filesToAdd: string | string[]) {
  try {
    writeMessage({
      type: 'info',
      message: 'Writing in the file ".npmignore"...'
    })

    const npmIgnoreFilePath = `${process.cwd()}/.npmignore`

    const existsNpmIgnoreFile = await exists(npmIgnoreFilePath)

    if (!existsNpmIgnoreFile) {
      await createEmptyFile('.npmignore')
    }

    const ignoredFiles = await fs.readFile(npmIgnoreFilePath, {
      encoding: UTF8_ENCODING
    })

    if (ignoredFiles === '') {
      const filesToAddIsAnArray = Array.isArray(filesToAdd)

      await fs.writeFile(
        npmIgnoreFilePath,
        filesToAddIsAnArray ? filesToAdd.join('\n') : filesToAdd,
        {
          encoding: UTF8_ENCODING
        }
      )

      writeMessage({
        type: 'info',
        message: '".npmignore" file modified successfully'
      })

      return
    }

    const ignoredFilesArray = ignoredFiles
      .split('\n')
      .filter(ignoreFile => ignoreFile !== '')
      .map(ignoreFile => ignoreFile.trim().replace(SPECIAL_CHARS_REGEX, ''))

    await fs.writeFile(
      npmIgnoreFilePath,
      ignoredFilesArray.concat(filesToAdd).join('\n'),
      {
        encoding: UTF8_ENCODING
      }
    )

    writeMessage({
      type: 'info',
      message: '".npmignore" file modified successfully'
    })
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('NpmIgnoreWrite')
    })

    process.exit(1)
  }
}
