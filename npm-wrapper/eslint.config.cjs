const prettier = require("eslint-plugin-prettier/recommended")
const eslintConfigPrettier = require("eslint-config-prettier/flat")
const globals = require("globals")

module.exports = (async function config() {
  return [
    eslintConfigPrettier,
    prettier,

    {
      ignores: ["node_modules"],
    },

    {
      files: ["**/*.js"],
      languageOptions: {
        sourceType: "commonjs",
        ecmaVersion: 2022,
        globals: {
          ...globals.node,
        },
      },
      rules: {
        "prettier/prettier": "error",
      },
    },
  ]
})()
