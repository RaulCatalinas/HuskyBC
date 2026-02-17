# 🐶 HuskyBC

**Simplify Husky configuration with a single command.**

HuskyBC is a blazing-fast CLI tool built in Rust that automates Git hooks setup in your Node.js projects.

> ⚠️ **Disclaimer:** HuskyBC is an **unofficial community tool**. It is not affiliated with or endorsed by [typicode/husky](https://github.com/typicode/husky). This CLI automates the setup of the official Husky package and related tools.

[![npm version](https://img.shields.io/npm/v/huskybc)](https://www.npmjs.com/package/huskybc)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)

---

## 📋 Table of Contents

- [Features](#-features)
- [Requirements](#-requirements)
- [Installation](#-installation)
- [Usage](#-usage)
- [Development](#-development)
- [Contributing](#-contributing)
- [Contributors](#-contributors)
- [License](#-license)
- [Acknowledgments](#-acknowledgments)

---

## ✨ Features

- 🚀 **Fast** - Built in Rust for maximum performance
- 🎯 **Simple** - Interactive wizard guides you through setup
- 🔧 **Flexible** - Configure Husky, Commitlint, Lint-staged, or all together
- 🌍 **Cross-platform** - Works on Windows, macOS, and Linux (x64 & ARM)

---

## 📋 Requirements

### For users

- **Node.js** 18.0.0 or higher
- **package.json** in your project root

### For development

- **Rust** 1.93.0 or higher ([Install Rust](https://rustup.rs/))
- **Node.js** 18.0.0 or higher
- **Git**

---

## 📦 Installation

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

---

## 🖥️ Usage

```bash
huskybc --init
```

| Command         | Alias | Description                    |
| --------------- | ----- | ------------------------------ |
| `--help`        | `-h`  | Display help information       |
| `--version`     | `-v`  | Show the current version       |
| `--init`        | `-i`  | Initialize Husky configuration |
| `--collaborate` | `-c`  | Open GitHub repository         |

---

## 🛠️ Development

### Clone the repository

```bash
git clone https://github.com/RaulCatalinas/HuskyBC.git
cd HuskyBC
```

### Project structure

```
HuskyBC/
├── src/                    # Rust source code
│   ├── main.rs
│   ├── cli/                # CLI input handling
│   ├── commands/           # Command handlers
│   ├── config/             # Configuration logic
│   │   └── presets/        # Configuration presets
│   ├── constants/          # App constants
│   ├── types/              # Type definitions
│   └── utils/              # Shared utilities
│
├── npm-wrapper/            # Node.js wrapper for npm distribution
│   ├── bin/
│   │   └── huskybc.js      # Entry point
│   └── postinstall.js      # Binary downloader
│
└── .github/
    └── workflows/
        └── github-release.yml     # CI/CD pipeline
```

### Build

```bash
# Debug build
cargo build

# Release build (optimized)
cargo build --release

# Run in debug mode
cargo run -- --help
```

### Release

Releases are automated via GitHub Actions. To trigger a new release:

```bash
# 1. Update version in Cargo.toml
# version = "x.x.x"

# 2. Commit and push
git add Cargo.toml
git commit -m "chore: bump version to x.x.x"
git push origin main

# 3. Create and push tag
git tag vx.x.x
git push origin vx.x.x

# GitHub Actions will automatically:
# - Build binaries for all platforms
# - Create a GitHub Release
# - Upload all binaries
```

### Supported platforms

| OS      | Architecture            | Binary                       |
| ------- | ----------------------- | ---------------------------- |
| Linux   | x86_64                  | `huskybc-x86_64-linux`       |
| Linux   | aarch64                 | `huskybc-aarch64-linux`      |
| macOS   | x86_64 (Intel)          | `huskybc-x86_64-macos`       |
| macOS   | aarch64 (Apple Silicon) | `huskybc-aarch64-macos`      |
| Windows | x86_64                  | `huskybc-x86_64-windows.exe` |

---

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details on how to get started, our code of conduct, and the process for submitting pull requests.

---

## 👥 Contributors

Learn more about the team behind HuskyBC on our [Contributors page](AUTHORS.md).

---

## 📄 License

This project is licensed under the [GPL-3.0 License](LICENSE).

---

## 🌟 Acknowledgments

- [Husky](https://github.com/typicode/husky) by @typicode - The tool that makes Git hooks management possible
- [Commitlint](https://commitlint.js.org/) - Lint commit messages
- [Lint-staged](https://github.com/okonet/lint-staged) - Run linters on staged files

---

Made with ❤️ and 🦀 Rust
