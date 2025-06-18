#!/usr/bin/env node

const { spawn } = require("child_process")
const path = require("path")
const os = require("os")

const platform = os.platform()
const arch = os.arch()

const binName = {
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
}

const platformBin = binName[platform]?.[arch]

if (!platformBin) {
  console.error(`Platform ${platform} with architecture ${arch} not supported`)
  process.exit(1)
}

const binPath = path.join(__dirname, "bin", platformBin)

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
