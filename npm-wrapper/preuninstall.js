#!/usr/bin/env node

const { rm, access } = require("node:fs/promises")
const { join } = require("node:path")

function getInstallDir() {
  const home = process.env.HOME || process.env.USERPROFILE

  const dirs = {
    win32: join(
      process.env.APPDATA || join(home, "AppData", "Roaming"),
      "huskybc",
    ),
    darwin: join(home, "Library", "Application Support", "huskybc"),
    linux: join(
      process.env.XDG_DATA_HOME || join(home, ".local", "share"),
      "huskybc",
    ),
  }

  return dirs[process.platform] ?? join(home, ".huskybc")
}

async function fileExists(filePath) {
  try {
    await access(filePath)
    return true
  } catch {
    return false
  }
}

async function uninstall() {
  try {
    const installDir = getInstallDir()

    if (!(await fileExists(installDir))) {
      console.log("✓ Nothing to uninstall")
      return
    }

    await rm(installDir, { recursive: true, force: true })

    console.log("✓ HuskyBC uninstalled successfully!")
  } catch (error) {
    console.error("✗ Uninstallation failed:", error.message)
    process.exit(1)
  }
}

uninstall()
