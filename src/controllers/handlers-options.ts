// Constants
import { ErrorMessages } from '@/constants/errors'
import { REPOSITORY } from '@/constants/github'

// Third-Party libraries
import opener from 'opener'

// Utils
import { addCommitlint } from '@/utils/add-commitlint'
import { generateCommitlintConfig } from '@/utils/commitlint'
import { generateHuskyConfig } from '@/utils/husky-library'
import { existPackageJson } from '@/utils/package-json'
import { getPackageManger } from '@/utils/package-managers'

// NodeJS
import process from 'node:process'

export const handlerOptionBuild = async () => {
  try {
    const packageJsonPath = `${process.cwd()}/package.json`
    const existPackageJsonInTheCurrentDirectory =
      await existPackageJson(packageJsonPath)

    if (!existPackageJsonInTheCurrentDirectory) {
      console.error(ErrorMessages.NotFound)

      process.exit(1)
    }

    const packageManagerToUse = await getPackageManger()
    const useCommitlint = await addCommitlint()

    await generateHuskyConfig(packageManagerToUse)

    if (useCommitlint) {
      await generateCommitlintConfig(packageManagerToUse)
    }
  } catch {
    console.error(ErrorMessages.Default)
    process.exit(1)
  }
}

export const handlerOptionCollaborate = () => {
  console.log('Opening the GitHub repository...')

  setTimeout(() => opener(REPOSITORY), 500)
}
