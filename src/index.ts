// Config'
import { configureCLI } from './config/cli'
import { configureOptions } from './config/options'

// Controllers
import { huskyConfigController } from './controllers/husky-config'

// Constants
import { program } from './constants/huskybc'

configureCLI()
configureOptions()

const thereIsFirstArgument = program.args[0] !== undefined

if (!thereIsFirstArgument) {
  program.action(huskyConfigController)
}

program.parse()
