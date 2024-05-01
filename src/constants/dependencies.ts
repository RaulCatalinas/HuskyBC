import type { PackageManager } from '@/types/package-manger'

export const INSTALLATION_COMMANDS: Record<PackageManager, string> = {
  npm: 'install',
  yarn: 'add',
  pnpm: 'add',
  bun: 'add'
}
