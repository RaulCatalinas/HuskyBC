#!/usr/bin/env node

const { spawn } = require("child_process")
const { join } = require("path")
const { platform, arch } = require("os")

const userPlatform = platform()
const userArch = arch()

const binName = Object.freeze({
  win32: {
    x64: "huskybc-windows-amd64.exe",
    arm64: "huskybc-windows-arm64.exe",
  },
  darwin: {
    x64: "huskybc-darwin-amd64",
    arm64: "huskybc-darwin-arm64",
  },
  linux: {
    x64: "huskybc-linux-amd64",
    arm64: "huskybc-linux-arm64",
  },
})

const platformBin = binName[userPlatform]?.[userArch]

if (!platformBin) {
  console.error(
    `Platform ${userPlatform} with architecture ${userArch} not supported`
  )
  process.exit(1)
}

const binPath = join(__dirname, "bin", platformBin)

const args = process.argv.slice(2).map((arg) => arg.trim())

try {
  const proc = spawn(binPath, args, { stdio: "inherit" })

  proc.on("close", (code) => {
    process.exit(code ?? 0)
  })

  proc.on("error", (err) => {
    console.error(`Error executing ${binPath}: ${err.message}`)
    process.exit(1)
  })
} catch (err) {
  console.error(`Error starting process: ${err.message}`)
  process.exit(1)
}
