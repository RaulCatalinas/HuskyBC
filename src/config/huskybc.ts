// Commander
import { createCommand } from 'commander'

// Constants
import { VERSION } from '../constants/version'

export const program = createCommand('HuskyBC')

export function configureCLI() {
  program
    .description('Command line for easy Husky configuration')
    .version(VERSION)
}
