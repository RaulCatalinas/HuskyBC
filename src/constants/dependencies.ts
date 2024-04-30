import type { PackageManager } from '@/types/package-manger'

export const INSTALLATION_COMMANDS: Record<PackageManager, string> = {
  npm: 'npm install',
  yarn: 'yarn add',
  pnpm: 'pnpm add',
  bun: 'bun add'
}
