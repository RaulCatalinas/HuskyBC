// Types
import type { PackageJsonScript } from '@/types/package-json'
import type { PackageManager } from '@/types/package-manger'

// Utils
import { writeMessage } from './console'
import { installDependencies } from './dependencies'
import { getErrorMessage } from './errors'
import { modifyNpmIgnore } from './npm'
import { addScript } from './package-json'
import { createFolder, exists } from './user-os'

// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'
import { LINT_STAGED_CONFIG } from '@/constants/lint-staged'

interface Props {
  packageManagerToUse: PackageManager
  packageJsonPath: string
  useCommitlint: boolean
  shouldPublishToNpm: boolean
}

type AddNecessaryScriptsToPackageJsonProps = Omit<Props, 'useCommitlint'>
type CreateHuskyConfigFiles = Pick<
  Props,
  'packageManagerToUse' | 'useCommitlint'
>

export async function generateHuskyConfig({
  packageManagerToUse,
  packageJsonPath,
  useCommitlint,
  shouldPublishToNpm
}: Props) {
  try {
    writeMessage({
      type: 'config',
      message: "Generating Husky's configuration..."
    })

    await Promise.all(
      [
        installDependencies({
          packageManagerToUse,
          packagesToInstall: 'husky'
        }),

        createHuskyConfigFiles({ packageManagerToUse, useCommitlint }),

        addNecessaryScriptsToPackageJson({
          packageJsonPath,
          packageManagerToUse,
          shouldPublishToNpm
        }),

        shouldPublishToNpm ? modifyNpmIgnore('.husky') : null
      ].filter(promise => promise != null)
    )

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

async function createHuskyConfigFiles({
  packageManagerToUse,
  useCommitlint
}: CreateHuskyConfigFiles) {
  writeMessage({
    type: 'info',
    message: 'Creating configuration file...'
  })

  const existsHuskyFolder = await exists(`${process.cwd()}/.husky`)

  if (!existsHuskyFolder) {
    await createFolder('.husky')
  }

  const preCommitFileValue = useCommitlint
    ? LINT_STAGED_CONFIG[packageManagerToUse]
    : `${packageManagerToUse} test`

  await fs.writeFile('.husky/pre-commit', preCommitFileValue, {
    encoding: UTF8_ENCODING
  })

  writeMessage({
    type: 'success',
    message: 'Configuration file (pre-commit) created successfully'
  })
}

async function addNecessaryScriptsToPackageJson({
  packageJsonPath,
  packageManagerToUse,
  shouldPublishToNpm
}: AddNecessaryScriptsToPackageJsonProps) {
  const huskyScriptsForYarn: PackageJsonScript | PackageJsonScript[] =
    !shouldPublishToNpm
      ? { key: 'postinstall', value: 'husky' }
      : [
          { key: 'postinstall', value: 'husky' },
          { key: 'prepack', value: 'pinst --disable' },
          { key: 'postpack', value: 'pinst --enable' }
        ]

  const scriptsToAdd: PackageJsonScript | PackageJsonScript[] =
    packageManagerToUse !== 'yarn'
      ? { key: 'prepare', value: 'husky' }
      : huskyScriptsForYarn

  await addScript({
    packageJsonPath,
    scriptsToAdd
  })
}
