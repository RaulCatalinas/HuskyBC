// Types
import type { PackageManager } from '@/types/package-manger'

// Utils
import { installDependencies } from './dependencies'
import { addScript } from './package-json'

// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'
import { ErrorMessages } from '@/constants/errors'
import { HUSKY_CONFIG } from '@/constants/husky-library'

export async function generateHuskyConfig(packageManagerToUse: PackageManager) {
  try {
    console.log("Generating Husky's configuration...")

    await installDependencies({
      packageManagerToUse,
      packagesToInstall: 'husky'
    })

    console.log('Creating configuration file...')

    await fs.writeFile('.husky/pre-commit', HUSKY_CONFIG[packageManagerToUse], {
      encoding: UTF8_ENCODING
    })

    console.log('Configuration file (pre-commit) created')

    console.log('Modifying package.json')

    await addScript({ key: 'prepare', value: 'husky' })

    console.log('Modified package.json')
    console.log("Husky's configuration generated successfully")
  } catch {
    console.error(ErrorMessages.Husky)
    process.exit(1)
  }
}
