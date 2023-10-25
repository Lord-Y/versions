import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

let api_url: string
if (process.env.API_GATEWAY) {
  api_url = `${process.env.API_GATEWAY}`
} else {
  api_url = 'http://127.0.0.1:8081'
}
const port: number = 3000

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  base: '/ui/',
  preview: {
    port: port,
  },
  server: {
    port: port,
    strictPort: true,
    proxy: {
      '/api': {
        target: api_url,
      },
    },
  },
})
