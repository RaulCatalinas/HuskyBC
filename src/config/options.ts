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
      '--collaborate',
      'Open GitHub repository for collaboration',
      handlerOptionCollaborate
    )
    .option('--build', "Start Husky's configuration", handlerOptionBuild)
}
