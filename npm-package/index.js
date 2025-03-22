#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const os = require('os');


const binName = {
  darwin: 'husky-cli-macos',
  linux: 'husky-cli-linux',
  win32: 'husky-cli-windows.exe'
}

const platform = os.platform();

if (!binName[platform]) {
  console.error('Unsupported platform');
  process.exit(1);
}

const binPath = path.join(__dirname, 'bin', binName[platform]);
const args = process.argv.slice(2);

const proc = spawn(binPath, args, { stdio: 'inherit' });

proc.on('close', (code) => {
  process.exit(code);
});