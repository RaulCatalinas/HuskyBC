export interface PackageJson {
  name: string
  module: string
  type: string
  version: string
  scripts: Record<string, string>
  dependencies: Record<string, string>
  devDependencies: Record<string, string>
}

export interface PackageJsonScript {
  key: string
  value: string
}
