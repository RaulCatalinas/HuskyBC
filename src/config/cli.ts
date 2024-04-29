import { program } from '@/constants/huskybc'
import { VERSION } from '@/constants/version'

export function configureCLI() {
  program
    .description('Command line for easy Husky configuration')
    .version(VERSION)
}
