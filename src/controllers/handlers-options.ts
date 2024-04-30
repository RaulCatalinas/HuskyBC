// Constants
import { REPOSITORY } from '@/constants/github'

// Third-Party libraries
import opener from 'opener'

export const handlerOptionBuild = async () => {
  console.log("Generating Husky's configuration...")

  // code...
}

export const handlerOptionCollaborate = () => {
  console.log('Opening the GitHub repository...')

  setTimeout(() => opener(REPOSITORY), 500)
}
