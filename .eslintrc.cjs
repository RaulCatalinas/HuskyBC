module.exports = {
  overrides: [
    {
      files: ['*.ts'],
      extends: 'love',
      rules: {
        '@typescript-eslint/space-before-function-paren': 'off',
        '@typescript-eslint/explicit-function-return-type': 'off'
      }
    }
  ],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module'
  }
}
