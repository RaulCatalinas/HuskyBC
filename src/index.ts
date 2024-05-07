// Config
import { configureCLI } from './config/cli'
import { configureDefaultOption, configureOptions } from './config/options'

// Constants
import { huskybc } from './constants/huskybc'

configureCLI()
configureOptions()

huskybc.parse()

configureDefaultOption()
