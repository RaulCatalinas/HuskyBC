// Constants
import { huskybc } from '@/constants/huskybc'

// Controllers
import {
  handlerOptionBuild,
  handlerOptionCollaborate
} from '@/controllers/handlers-options'

export function configureOptions() {
  huskybc
    .option(
      '-co, --collaborate',
      'Open GitHub repository for collaboration',
      handlerOptionCollaborate
    )
    .option('-b, --build', "Start Husky's configuration", handlerOptionBuild)
}

export function configureDefaultOption() {
  const options = huskybc.opts()

  // eslint-disable-next-line @typescript-eslint/strict-boolean-expressions
  if (!options.build && !options.collaborate && !options.b && !options.co) {
    huskybc.help()
  }
}
