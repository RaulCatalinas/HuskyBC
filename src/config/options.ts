// Third-Party libraries
import opener from 'opener'

// Constants
import { REPOSITORY } from '../constants/github'
import { program } from '../constants/huskybc'

export function configureOptions() {
  program.option(
    '--collaborate',
    'Open GitHub repository for collaboration',
    () => {
      opener(REPOSITORY)
    }
  )
}
