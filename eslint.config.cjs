const js = require("@eslint/js")
const globals = require("globals")
const { defineConfig } = require("eslint/config")

module.exports = defineConfig([
  {
    files: ["**/*.{js,mjs,cjs}"],
    plugins: { js },
    extends: ["js/recommended"],
    ignores: ["node_modules"],
  },
  {
    files: ["**/*.js"],
    languageOptions: { sourceType: "commonjs", ecmaVersion: "latest" },
  },
  {
    files: ["**/*.{js,mjs,cjs}"],
    languageOptions: { globals: globals.node },
  },
])
