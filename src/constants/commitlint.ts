import type { PackageManager } from '@/types/package-manger'

// TODO: Add configuration for yarn
export const COMMITLINT_CONFIG: Record<PackageManager, string> = {
  npm: 'npx --no -- commitlint --edit $1',
  pnpm: 'pnpm dlx commitlint --edit $1',
  yarn: '',
  bun: 'bunx commitlint --edit $1'
}
