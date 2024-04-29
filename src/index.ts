// Config
import { configureCLI } from './config/cli'
import { configureOptions } from './config/options'

// Constants
import { program } from './constants/huskybc'

configureCLI()
configureOptions()

program.parse()
