import { ISSUES } from './github'

export const ERROR_MESSAGES = {
  NotFound: "The package.json file wasn't found in the current directory.",

  Default: `Something went wrong, please try again later, if the error persists please report it on ${ISSUES}.`,

  Dependencies: `An error occurred while installing dependencies, please try again later, if the error persists please report it on ${ISSUES}.`,

  Commitlint: `An error has occurred during the Commitlint configuration process, please try again later, if the error persists please report it on ${ISSUES}.`,
  Husky: `An error has occurred during the Husky configuration process, please try again later, if the error persists please report it on ${ISSUES}.`,

  PackageManagerSelection: `An error occurred while selecting the package manager, please try again later, if the error persists, please report it on ${ISSUES}.`,

  CommitlintSelection: `An error occurred while determining your choice for Commitlint, please try again later, if the error persists, please report it on ${ISSUES}.`,

  AddScript: `An error occurred while adding a script to the package.json file, please try again later, if the error persists, please report it on ${ISSUES}.`,

  CreateFolder: `An error occurred while creating the folder: {folderName}, please try again later, if the error persists, please report it on ${ISSUES}.`,

  CheckFileExists: `An error occurred while checking if the file/folder exists, please try again later, if the error persists, please report it on ${ISSUES}.`
} as const
