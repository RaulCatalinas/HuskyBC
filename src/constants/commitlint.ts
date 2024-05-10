import type { PackageManager } from '@/types/package-manger'

export const COMMITLINT_CONFIG: Record<PackageManager, string> = {
  npm: 'npx --no -- commitlint --edit $1',
  pnpm: 'pnpm dlx commitlint --edit $1',
  yarn: 'yarn dlx commitlint --edit $1',
  bun: 'bunx commitlint --edit $1'
}
