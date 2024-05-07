import { huskybc } from '@/constants/huskybc'
import { VERSION } from '@/constants/version'

export function configureCLI() {
  huskybc
    .description('Command line for easy Husky configuration')
    .version(VERSION)
    .showHelpAfterError()
}
