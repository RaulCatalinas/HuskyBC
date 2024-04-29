// Constants
import { program } from '@/constants/huskybc'

// Controllers
import {
  handlerOptionBuild,
  handlerOptionCollaborate
} from '@/controllers/handlers-options'

export function configureOptions() {
  program
    .option(
      '-co, --collaborate',
      'Open GitHub repository for collaboration',
      handlerOptionCollaborate
    )
    .option('-b, --build', "Start Husky's configuration", handlerOptionBuild)
}

export function configureDefaultOption() {
  const options = program.opts()

  // eslint-disable-next-line @typescript-eslint/strict-boolean-expressions
  if (!options.build && !options.collaborate && !options.b && !options.co) {
    program.help()
  }
}
