import type { PackageManager } from '@/types/package-manger'

export const LINT_STAGED_CONFIG: Record<PackageManager, string> = {
  npm: 'npx lint-staged',
  pnpm: 'pnpm dlx lint-staged',
  yarn: 'yarn dlx lint-staged',
  bun: 'bunx lint-staged'
}
