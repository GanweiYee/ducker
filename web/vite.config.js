import { defineConfig } from "vite";
import * as path from 'path'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import vue from '@vitejs/plugin-vue';

export default ({
  command,
  mode
}) => {
  const NODE_ENV = process.env.NODE_ENV || 'development'
  const envFiles = [
    `.env.${NODE_ENV}`
  ]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }
  return {
    base: "",
    plugins: [
      vue()
    ],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, 'src'),
      }
    },
    optimizeDeps: {
      include: ['request'],
    },
    build: {
      target: 'modules',
      outDir: 'dist',
      assetsDir: 'assets',
      minify: 'terser' // 混淆器
    },
    server: {
      cors: true,
      open: true,
      port: process.env.VITE_CLI_PORT,
      proxy: {
        [process.env.VITE_BASE_API]: {
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/v1/`,   //代理接口
          changeOrigin: true,
          rewrite: path => path.replace(new RegExp('^' + process.env.VITE_BASE_API), ''),
        }
      }
    }
  }
}