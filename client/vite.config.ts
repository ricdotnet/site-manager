import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: './dist'
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@components': path.resolve(__dirname, './src/components'),
      '@pages': path.resolve(__dirname, './src/pages'),
      '@composables': path.resolve(__dirname, './src/composables'),
      '@stores': path.resolve(__dirname, './src/stores'),
      '@utils': path.resolve(__dirname, './src/utils'),
      '@validators': path.resolve(__dirname, './src/validators'),
      '@types': path.resolve(__dirname, './src/types.ts'),
    },
  }
});
