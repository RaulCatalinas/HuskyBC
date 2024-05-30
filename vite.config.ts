// Vite
import { defineConfig } from 'vite'

// Plugins
import tsConfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [tsConfigPaths()],
  ssr: {
    // bundle and treeshake everything
    noExternal: true
  },
  build: {
    ssr: true,
    lib: {
      entry: ['src/index.ts'],
      formats: ['cjs']
    },
    outDir: 'dist',
    minify: true
  }
})
