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
import { modifyNpmIgnore } from './npm'
import { addScript } from './package-json'

interface Props {
  packageManagerToUse: PackageManager
  packageJsonPath: string
  shouldPublishToNpm: boolean
}

export async function generateCommitlintConfig({
  packageManagerToUse,
  packageJsonPath,
  shouldPublishToNpm
}: Props) {
  try {
    writeMessage({
      type: 'config',
      message: 'Configuring commitlint...'
    })

    await Promise.all(
      [
        installDependencies({
          packageManagerToUse,
          packagesToInstall: [
            'lint-staged',
            '@commitlint/cli',
            '@commitlint/config-conventional'
          ]
        }),

        addNecessaryScriptsToPackageJson(packageJsonPath),

        createCommitlintConfigFiles(packageManagerToUse),

        shouldPublishToNpm
          ? modifyNpmIgnore(['.lintstagedrc', 'commitlint.config.js'])
          : null
      ].filter(promise => promise != null)
    )

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

async function createCommitlintConfigFiles(
  packageManagerToUse: PackageManager
) {
  writeMessage({
    type: 'info',
    message: 'Creating configuration files...'
  })

  await Promise.all([
    fs.writeFile('.husky/commit-msg', COMMITLINT_CONFIG[packageManagerToUse], {
      encoding: UTF8_ENCODING
    }),
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
          'src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}': [
            'prettier src --check',
            'eslint src --max-warnings 0'
          ]
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
}

async function addNecessaryScriptsToPackageJson(packageJsonPath: string) {
  writeMessage({
    type: 'info',
    message: 'Modifying package.json...'
  })

  await addScript({
    packageJsonPath,
    scriptsToAdd: [
      { key: 'lint', value: 'eslint src' },
      { key: 'lint:fix', value: 'eslint src --fix' },
      { key: 'format', value: 'prettier src --check' },
      { key: 'format:write', value: 'prettier src --write' }
    ]
  })

  writeMessage({
    type: 'info',
    message: 'package.json modified successfully'
  })
}
