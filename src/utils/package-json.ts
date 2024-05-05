// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'

// Types
import type { PackageJson } from '@/types/package-json'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'

interface Props {
  key: string
  value: string
  packageJsonPath: string
}

export async function addScript({ key, value, packageJsonPath }: Props) {
  try {
    const packageJsonData = await fs.readFile(packageJsonPath, {
      encoding: UTF8_ENCODING
    })

    const packageJsonObj: PackageJson = JSON.parse(packageJsonData)

    packageJsonObj.scripts[key] = value

    await fs.writeFile(
      packageJsonPath,
      JSON.stringify(packageJsonObj, null, 2),
      {
        encoding: UTF8_ENCODING
      }
    )
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('AddScript')
    })
    process.exit(1)
  }
}
