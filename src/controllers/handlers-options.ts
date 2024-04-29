import opener from 'opener'
import { REPOSITORY } from '../constants/github'

export const handlerOptionBuild = () => {
  console.log("Generating Husky's configuration")
  // code...
}

export const handlerOptionCollaborate = async () => {
  console.log('Open GitHub repository for collaboration...')
  await Bun.sleep(500)
  opener(REPOSITORY)
}
