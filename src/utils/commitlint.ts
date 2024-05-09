// Types
import type { PackageManager } from '@/types/package-manger'

// NodeJS
import fs from 'node:fs/promises'

// Constants
import { COMMITLINT_CONFIG } from '@/constants/commitlint'
import { UTF8_ENCODING } from '@/constants/encoding'

// Utils
import { writeMessage } from './console'
import { installDependencies } from './dependencies'
import { getErrorMessage } from './errors'

export async function generateCommitlintConfig(
  packageManagerToUse: PackageManager
) {
  try {
    writeMessage({
      type: 'config',
      message: 'Configuring commitlint...'
    })

    await installDependencies({
      packageManagerToUse,
      packagesToInstall: [
        'lint-staged',
        '@commitlint/cli',
        '@commitlint/config-conventional'
      ]
    })

    writeMessage({
      type: 'info',
      message: 'Creating configuration files...'
    })

    await Promise.all([
      fs.writeFile(
        '.husky/commit-msg',
        COMMITLINT_CONFIG[packageManagerToUse],
        {
          encoding: UTF8_ENCODING
        }
      ),
      fs.writeFile(
        'commitlint.config.js',
        "export default { extends: ['@commitlint/config-conventional'] }",
        {
          encoding: UTF8_ENCODING
        }
      ),
      fs.writeFile(
        '.lintstagedrc',
        JSON.stringify(
          {
            'src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}': []
          },
          null,
          2
        ),
        {
          encoding: UTF8_ENCODING
        }
      )
    ])

    writeMessage({
      type: 'info',
      message:
        'Configuration files (commit-msg, commitlint.config.js and .lintstagedrc) created successfully'
    })
    writeMessage({
      type: 'success',
      message: "commitlint's configuration generated successfully"
    })
  } catch (error) {
    writeMessage({
      type: 'error',
      message: getErrorMessage('Commitlint')
    })
    process.exit(1)
  }
}
