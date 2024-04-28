// Config
import { configureCLI, program } from './config/huskybc'

// Controllers
import { huskyConfigController } from './controllers/husky-config'

configureCLI()

const thereIsFirstArgument = program.args[0] !== undefined

if (!thereIsFirstArgument) {
  program.action(huskyConfigController)
}

program.parse()
