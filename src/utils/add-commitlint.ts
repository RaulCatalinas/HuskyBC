import inquirer from 'inquirer'

export async function addCommitlint(): Promise<boolean> {
  try {
    const { addCommitlint } = await inquirer.prompt({
      type: 'confirm',
      default: true,
      name: 'addCommitlint',
      message: 'Do you wanna add commitlint?:'
    })

    return addCommitlint
  } catch (error) {
    console.error(error)
    throw new Error('Something went wrong, try again later.')
  }
}
