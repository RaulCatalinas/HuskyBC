import pluginJs from '@eslint/js'
import pluginImport from 'eslint-plugin-import'
import eslintPluginPrettier from 'eslint-plugin-prettier/recommended'
import globals from 'globals'
import tsEslint from 'typescript-eslint'

export default [
  pluginJs.configs.recommended,
  ...tsEslint.configs.recommended,
  eslintPluginPrettier,
  {
    languageOptions: { globals: globals.nodeBuiltin },
    plugins: {
      import: pluginImport
    },
    rules: {
      'import/order': [
        'error',
        {
          'newlines-between': 'always',
          alphabetize: {
            order: 'asc',
            caseInsensitive: true
          },
          groups: [
            'NodeJS',
            'Third-Party libraries',
            'Constants',
            'Utils',
            'User-Input',
            'Controllers',
            'Config',
            'Types'
          ],
          pathGroups: [
            {
              pattern: 'node:**',
              group: 'NodeJS',
              patternOptions: {
                nocomment: false
              }
            },
            {
              pattern: '@/utils/**',
              group: 'Utils',
              patternOptions: {
                nocomment: false
              }
            },
            {
              pattern: '@/types/**',
              group: 'Types',
              patternOptions: {
                nocomment: false
              }
            },
            {
              pattern: '@/user-input/**',
              group: 'User-Input',
              patternOptions: {
                nocomment: false
              }
            },
            {
              pattern: '@/controllers/**',
              group: 'Controllers',
              patternOptions: {
                nocomment: false
              }
            },
            {
              pattern: '@/constants/**',
              group: 'Constants',
              patternOptions: {
                nocomment: false
              }
            },
            {
              pattern: '@/config/**',
              group: 'Config',
              patternOptions: {
                nocomment: false
              }
            }
          ]
        }
      ],
      'import/newline-after-import': [
        'error',
        {
          count: 2,
          considerComments: true
        }
      ]
    }
  }
]
