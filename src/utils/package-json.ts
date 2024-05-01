// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'

// Types
import type { PackageJson } from '@/types/package-json'

interface Props {
  key: string
  value: string
}

export async function addScript({ key, value }: Props) {
  try {
    const packageJsonPath = `${process.cwd()}/package.json`

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
  } catch (error) {
    console.error(error)
    throw new Error('Something went wrong, try again later.')
  }
}
