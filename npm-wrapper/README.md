# 🐶 HuskyBC

**Simplify Husky configuration with a single command.**

HuskyBC is a blazing-fast CLI tool built in Rust that automates Git hooks setup in your Node.js projects.

> ⚠️ **Disclaimer:** HuskyBC is an **unofficial community tool**. It is not affiliated with or endorsed by [typicode/husky](https://github.com/typicode/husky). This CLI automates the setup of the official Husky package and related tools.

[![npm version](https://img.shields.io/npm/v/huskybc)](https://www.npmjs.com/package/huskybc)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](https://opensource.org/licenses/GPL-3.0)

## ✨ Features

- 🚀 **Fast** - Built in Rust for maximum performance
- 🎯 **Simple** - Interactive wizard guides you through setup
- 🔧 **Flexible** - Configure Husky, Commitlint, Lint-staged, or all together
- 🌍 **Cross-platform** - Works on Windows, macOS, and Linux (x64 & ARM)

## 📦 Installation

Install globally using your preferred package manager:

```bash
# npm
npm install -g huskybc

# pnpm
pnpm add -g huskybc

# yarn
yarn global add huskybc

# bun
bun add -g huskybc
```

Or use directly (no installation needed):

```bash
# npm
npx huskybc --init

# pnpm
pnpm dlx huskybc --init

# yarn
yarn dlx huskybc --init

# bun
bunx huskybc --init
```

## 🚀 Quick Start

Navigate to your Node.js project and run:

```bash
huskybc --init
```

The interactive wizard will ask:

1. **Which package manager** do you use? (npm, yarn, pnpm, bun)
2. **What to configure?**
   - Only Husky
   - Husky + Commitlint
   - Husky + Lint-staged
   - Full setup (Husky + Commitlint + Lint-staged)

That's it! HuskyBC handles the rest:

- ✅ Installs dependencies
- ✅ Creates configuration files
- ✅ Sets up Git hooks
- ✅ Adds scripts to `package.json`

## 📖 Commands

| Command         | Alias | Description                    |
| --------------- | ----- | ------------------------------ |
| `--help`        | `-h`  | Display help information       |
| `--version`     | `-v`  | Show the current version       |
| `--init`        | `-i`  | Initialize Husky configuration |
| `--collaborate` | `-c`  | Open GitHub repository         |

## 📋 Requirements

- **Node.js** 18.0.0 or higher
- **package.json** in your project root

## 🎯 What Gets Configured

### Husky Only

```
✓ Installs husky
✓ Creates .husky/ directory
✓ Sets up pre-commit hook
✓ Adds "prepare" script to package.json
```

### Husky + Commitlint

```
✓ Everything from "Husky Only"
✓ Installs @commitlint/cli and @commitlint/config-conventional
✓ Creates commitlint.config.mjs
✓ Sets up commit-msg hook
```

### Husky + Lint-staged

```
✓ Everything from "Husky Only"
✓ Installs lint-staged
✓ Configures lint-staged in package.json
✓ Updates pre-commit hook to run lint-staged
```

### Full Setup

```
✓ All of the above combined
✓ Production-ready Git hooks workflow
```

## 💡 Example

```bash
$ cd my-awesome-project
$ huskybc --init

? Which package manager do you want to use? npm
? What do you want to configure? Husky + Commitlint + Lint-staged (Full)

→ Configuring full setup with npm...
✓ Created .husky/pre-commit
✓ Created .husky/commit-msg
✓ Created commitlint.config.mjs
✓ Full configuration complete!
```

## 🤔 Why HuskyBC?

### Before HuskyBC 😫

```bash
npm install husky @commitlint/cli @commitlint/config-conventional lint-staged -D
npx husky init
echo "npx lint-staged" > ".husky/pre-commit"
echo "npx --no -- commitlint --edit \$1" > .husky/commit-msg
# Create commitlint.config.mjs manually
# Configure lint-staged manually
# Cross your fingers it works...
```

### With HuskyBC 🎉

```bash
huskybc --init
# Answer 2 questions
# Done! ✨
```

## 🐞 Troubleshooting

### Binary not found after installation

```bash
# Reinstall the package
npm install -g huskybc --force
```

### Permission denied on Unix systems

```bash
# The postinstall script should handle this, but if needed:
chmod +x $(which huskybc)
```

### Command not recognized

```bash
# Make sure npm global bin is in your PATH
npm config get prefix
# Add <prefix>/bin to your PATH
```

## 🌟 Contributing

Contributions are welcome! Check out our [Contributing Guidelines](https://github.com/RaulCatalinas/HuskyBC/blob/main/CONTRIBUTING.md).

## 📝 License

GPL-3.0 © [Raul Catalinas](https://github.com/RaulCatalinas)

## 🙏 Credits & Acknowledgments

HuskyBC automates the setup of these amazing tools:

- [Husky](https://github.com/typicode/husky) by @typicode - Git hooks made easy
- [Commitlint](https://commitlint.js.org/) - Lint commit messages
- [Lint-staged](https://github.com/okonet/lint-staged) - Run linters on staged files

**Note:** This is a community tool. For issues with the underlying packages, please report them in their respective repositories.

## 🔗 Links

- [GitHub Repository](https://github.com/RaulCatalinas/HuskyBC)
- [Report Issues](https://github.com/RaulCatalinas/HuskyBC/issues)
- [npm Package](https://www.npmjs.com/package/huskybc)

---

Made with ❤️ and 🦀 Rust
