// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'

// Types
import type { PackageJson, PackageJsonScript } from '@/types/package-json'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'

interface AddScriptProps {
  packageJsonPath: string
  scriptsToAdd: PackageJsonScript | PackageJsonScript[]
}

interface ExistsSectionProps {
  packageJsonPath: string
  sectionToCheck: string
}

interface CreateEmptySectionProps {
  packageJsonPath: string
  sectionToCreate: string
}

export async function addScript({
  scriptsToAdd,
  packageJsonPath
}: AddScriptProps) {
  try {
    const existsScriptsSection = await existsSection({
      packageJsonPath,
      sectionToCheck: 'scripts'
    })

    if (!existsScriptsSection) {
      await createEmptySection({ packageJsonPath, sectionToCreate: 'scripts' })
    }

    const packageJsonData = await fs.readFile(packageJsonPath, {
      encoding: UTF8_ENCODING
    })

    const packageJsonObj: PackageJson = JSON.parse(packageJsonData)

    const scriptsToAddIsAnArray = Array.isArray(scriptsToAdd)

    if (scriptsToAddIsAnArray) {
      for (const { key, value } of scriptsToAdd) {
        packageJsonObj.scripts[key] = value
      }

      await fs.writeFile(
        packageJsonPath,
        JSON.stringify(packageJsonObj, null, 2),
        {
          encoding: UTF8_ENCODING
        }
      )

      return
    }

    const { key, value } = scriptsToAdd

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

async function existsSection({
  packageJsonPath,
  sectionToCheck
}: ExistsSectionProps) {
  try {
    const packageJsonData = await fs.readFile(packageJsonPath, {
      encoding: UTF8_ENCODING
    })

    const packageJsonObj = JSON.parse(packageJsonData)

    return packageJsonObj[sectionToCheck] !== undefined
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('CheckSection')
    })

    process.exit(1)
  }
}

async function createEmptySection({
  packageJsonPath,
  sectionToCreate
}: CreateEmptySectionProps) {
  try {
    const packageJsonData = await fs.readFile(packageJsonPath, {
      encoding: UTF8_ENCODING
    })

    const packageJsonObj = JSON.parse(packageJsonData)

    packageJsonObj[sectionToCreate] = {}

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
      message: getErrorMessage('CreateSection')
    })

    process.exit(1)
  }
}
