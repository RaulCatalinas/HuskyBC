// Types
import type { PackageManager } from '@/types/package-manger'

// Utils
import { writeMessage } from './console'
import { installDependencies } from './dependencies'
import { getErrorMessage } from './errors'
import { addScript } from './package-json'
import { createFolder, exists } from './user-os'

// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'
import { HUSKY_CONFIG } from '@/constants/husky-library'

interface Props {
  packageManagerToUse: PackageManager
  packageJsonPath: string
}

export async function generateHuskyConfig({
  packageManagerToUse,
  packageJsonPath
}: Props) {
  try {
    writeMessage({
      type: 'config',
      message: "Generating Husky's configuration..."
    })

    await installDependencies({
      packageManagerToUse,
      packagesToInstall: 'husky'
    })

    writeMessage({
      type: 'info',
      message: 'Creating configuration file...'
    })

    const existsHuskyFolder = await exists(`${process.cwd()}/.husky`)

    if (!existsHuskyFolder) {
      await createFolder('.husky')
    }

    await fs.writeFile('.husky/pre-commit', HUSKY_CONFIG[packageManagerToUse], {
      encoding: UTF8_ENCODING
    })

    writeMessage({
      type: 'info',
      message: 'Configuration file (pre-commit) created successfully'
    })

    writeMessage({
      type: 'info',
      message: 'Modifying package.json'
    })

    await addScript({ key: 'prepare', value: 'husky', packageJsonPath })

    writeMessage({
      type: 'info',
      message: 'Modified package.json'
    })
    writeMessage({
      type: 'success',
      message: "Husky's configuration generated successfully"
    })
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('Husky')
    })
    process.exit(1)
  }
}
