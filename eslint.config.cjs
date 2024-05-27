const eslintPluginPrettier = require('eslint-plugin-prettier/recommended')

module.exports = [
  eslintPluginPrettier,
  {
    ...require('eslint-config-love'),
    files: ['**/*.ts'],
    rules: {
      "@typescript-eslint/explicit-function-return-type": "off",
      "@typescript-eslint/explicit-function-return-type": "off"
    }
  }
]
