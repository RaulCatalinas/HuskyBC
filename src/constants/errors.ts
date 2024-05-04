import { ISSUES } from './github'

export enum ErrorMessages {
  NotFound = "The package.json file wasn't found in the current directory.",

  Default = `Something went wrong, try again later, please try again later, if the error persists please report it on ${ISSUES}.`,

  Dependencies = `An error occurred while installing dependencies, please try again later, if the error persists please report it on ${ISSUES}.`,

  Commitlint = `An error has occurred during the Commitlint configuration process, please try again later, if the error persists please report it on ${ISSUES}.`,

  Husky = `An error has occurred during the Husky configuration process, please try again later, if the error persists please report it on ${ISSUES}.`
}
