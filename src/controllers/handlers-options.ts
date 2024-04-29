// Constants
import { REPOSITORY } from '@/constants/github'

// Third-Party libraries
import opener from 'opener'

export const handlerOptionBuild = () => {
  console.log("Generating Husky's configuration...")

  // code...
}

export const handlerOptionCollaborate = () => {
  console.log('Open GitHub repository for collaboration...')

  setTimeout(() => opener(REPOSITORY), 500)
}
