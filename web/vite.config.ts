import { fileURLToPath, URL } from 'node:url'

import vue from '@vitejs/plugin-vue'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import { defineConfig, UserConfig } from 'vite'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const NODE_ENV = mode || 'development'
  const envFiles = [`.env.${NODE_ENV}`]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }

  const alias = {
    '@': fileURLToPath(new URL('./src', import.meta.url)),
  }

  const esbuild = {}
  const optimizeDeps = {}

  const config = {
    base: '/',
    root: './',
    publicDir: 'public',
    resolve: { alias },
    define: {
      'process.env': {},
    },
    server: {
      open: true,
      port: Number(process.env.VITE_WEB_PORT),
      proxy: {
        '/api': {
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
    plugins: [vue(), vueDevTools()],
    esbuild,
    optimizeDeps,
  } satisfies UserConfig

  return config
})
