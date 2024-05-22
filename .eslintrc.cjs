module.exports = {
  ...require('eslint-config-love'),
  files: ['*.ts'],
  extends: [
    'plugin:import/recommended',
    'plugin:import/typescript',
    'plugin:prettier/recommended'
  ],
  rules: {
    'import/order': [
      'error',
      {
        groups: [
          'NodeJS',
          'Third-Party libraries',
          'Utils',
          'Types',
          'User-Input',
          'Controllers',
          'Constants',
          'Config'
        ],
        pathGroups: [
          {
            pattern: 'node:**',
            group: 'NodeJS',
            position: 'before',
            patternOptions: {
              nocomment: false
            }
          },
          {
            pattern: '@/utils/**',
            group: 'Utils',
            position: 'before',
            patternOptions: {
              nocomment: false
            }
          },
          {
            pattern: '@/types/**',
            group: 'Types',
            position: 'before',
            patternOptions: {
              nocomment: false
            }
          },
          {
            pattern: '@/user-input/**',
            group: 'User-Input',
            position: 'before',
            patternOptions: {
              nocomment: false
            }
          },
          {
            pattern: '@/controllers/**',
            group: 'Controllers',
            position: 'before',
            patternOptions: {
              nocomment: false
            }
          },
          {
            pattern: '@/constants/**',
            group: 'Constants',
            position: 'before',
            patternOptions: {
              nocomment: false
            }
          },
          {
            pattern: '@/config/**',
            group: 'Config',
            position: 'before',
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
  },
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module'
  }
}
