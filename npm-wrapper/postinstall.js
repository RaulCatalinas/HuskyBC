#!/usr/bin/env node

const { mkdir, unlink, access, chmod, writeFile } = require("node:fs/promises")
const { createWriteStream } = require("node:fs")
const { Readable } = require("node:stream")
const { finished } = require("node:stream/promises")
const { join } = require("node:path")

const REPO = "RaulCatalinas/HuskyBC"

function getBinaryName() {
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
    console.error(`Your system: ${platform}-${arch}`)
    console.error(`Supported: x64 and arm64 on Windows, macOS, and Linux`)
    process.exit(1)
  }

  const ext = platform === "win32" ? ".exe" : ""
  return `huskybc-${mappedArch}-${mappedPlatform}${ext}`
}

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

async function getLatestReleaseUrl() {
  const binaryName = getBinaryName()
  const apiUrl = `https://api.github.com/repos/${REPO}/releases/latest`

  console.log("Fetching latest release info...")

  const res = await fetch(apiUrl, {
    headers: {
      "User-Agent": "huskybc-installer",
      Accept: "application/vnd.github+json",
    },
  })

  if (!res.ok) {
    throw new Error(
      `Failed to fetch latest release: ${res.status} ${res.statusText}`,
    )
  }

  const release = await res.json()
  const asset = release.assets.find((a) => a.name === binaryName)

  if (!asset) {
    throw new Error(
      `Binary "${binaryName}" not found in latest release (${release.tag_name})\n` +
        `Available assets: ${release.assets.map((a) => a.name).join(", ")}`,
    )
  }

  console.log(`Found version: ${release.tag_name}`)

  return asset.browser_download_url
}

async function download(url, dest) {
  console.log(`Downloading ${url}...`)

  const file = createWriteStream(dest)

  try {
    const res = await fetch(url, {
      headers: { "User-Agent": "huskybc-installer" },
    })

    if (!res.ok) {
      file.close()
      await unlink(dest).catch(() => {})
      throw new Error(`Failed to download: ${res.status} ${res.statusText}`)
    }

    await finished(Readable.fromWeb(res.body).pipe(file))

    console.log("✓ Download complete")
  } catch (err) {
    file.close()
    await unlink(dest).catch(() => {})
    throw err
  }
}

async function makeExecutable(filePath) {
  if (process.platform !== "win32") {
    try {
      await chmod(filePath, 0o755)
      console.log("✓ Made binary executable")
    } catch (err) {
      console.error("Warning: Could not make binary executable:", err.message)
    }
  }
}

async function fileExists(filePath) {
  try {
    await access(filePath)
    return true
  } catch {
    return false
  }
}

async function install() {
  try {
    const installDir = getInstallDir()
    const binaryPath = join(
      installDir,
      process.platform === "win32" ? "huskybc.exe" : "huskybc",
    )

    await mkdir(installDir, { recursive: true })

    if (await fileExists(binaryPath)) {
      console.log("✓ Binary already exists, skipping download")
      return
    }

    const downloadUrl = await getLatestReleaseUrl()

    await download(downloadUrl, binaryPath)

    await makeExecutable(binaryPath)

    console.log("✓ HuskyBC installed successfully!")
    console.log("  Run: npx huskybc --help")
  } catch (error) {
    console.error("✗ Installation failed:", error.message)
    console.error("\nPlease try:")
    console.error(
      `1. Manual download from: https://github.com/${REPO}/releases/latest`,
    )
    console.error(`2. Report this issue at: https://github.com/${REPO}/issues`)
    process.exit(1)
  }
}

install()
