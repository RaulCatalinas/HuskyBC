import type { PACKAGE_MANGERS } from '@/constants/package-mangers'

export type PackageManager = (typeof PACKAGE_MANGERS)[number]
