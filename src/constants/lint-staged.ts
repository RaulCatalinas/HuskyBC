import type { PackageManager } from '@/types/package-manger'

// TODO: Add configuration for yarn
export const LINT_STAGED_CONFIG: Record<PackageManager, string> = {
  npm: 'npx lint-staged',
  pnpm: 'pnpm dlx lint-staged',
  yarn: '',
  bun: 'bunx lint-staged'
}
