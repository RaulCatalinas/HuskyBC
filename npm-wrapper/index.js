#!/usr/bin/env node

const { spawn } = require("child_process")
const { join } = require("path")

const platform = process.platform
const arch = process.arch

const platformMap = {
  win32: "windows",
  darwin: "macos",
  linux: "linux",
}

const archMap = {
  x64: "x86_64",
  arm64: "aarch64",
}

const mappedPlatform = platformMap[platform]
const mappedArch = archMap[arch]

if (!mappedPlatform || !mappedArch) {
  console.error(`Unsupported platform: ${platform}-${arch}`)
  process.exit(1)
}

const ext = platform === "win32" ? ".exe" : ""
const binaryName = `bin/huskybc-${mappedArch}-${mappedPlatform}${ext}`
const binPath = join(__dirname, binaryName)

const proc = spawn(binPath, process.argv.slice(2), {
  stdio: "inherit",
  shell: false,
})

proc.on("close", (code) => {
  process.exit(code ?? 0)
})

proc.on("error", (err) => {
  console.error(`✗ Failed to execute HuskyBC binary`)
  console.error(`Error: ${err.message}`)
  console.error(`Binary: ${binPath}`)
  console.error(`\nTry reinstalling: npm install huskybc`)

  process.exit(1)
})
