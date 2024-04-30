// Types
import type { PackageManager } from '@/types/package-manger'

// NodeJS
import fs from 'node:fs/promises'

// Constants
import { COMMITLINT_CONFIG } from '@/constants/commitlint'
import { UTF8_ENCODING } from '@/constants/encoding'

// Utils
import { installDependencies } from './dependencies'

export async function configureCommitlint(packageManagerToUse: PackageManager) {
  try {
    console.log('Configuring commitlint...')

    installDependencies({
      packageManagerToUse,
      packagesToInstall: [
        'lint-staged',
        '@commitlint/cli',
        '@commitlint/config-conventional'
      ]
    })

    console.log('Creating configuration files...')

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
        `
          {
            "src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}": []
          }      
        `,
        {
          encoding: UTF8_ENCODING
        }
      )
    ])

    console.log(
      'Configuration files (commit-msg, commitlint.config.js and .lintstagedrc) created'
    )
  } catch (error) {
    console.error(error)
    throw new Error('Something went wrong, try again later.')
  }
}
