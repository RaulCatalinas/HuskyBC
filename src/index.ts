// Config
import { configureCLI } from './config/cli'
import { configureDefaultOption, configureOptions } from './config/options'

// Constants
import { program } from './constants/huskybc'

configureCLI()
configureOptions()

program.parse()

configureDefaultOption()
