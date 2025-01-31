import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  define: {
    VITE_API_URL: JSON.stringify(process.env.VITE_API_URL || "http://localhost:8080"),
  },
  server: {
    proxy: {
      '/search': {
        target: process.env.VITE_API_URL || 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
