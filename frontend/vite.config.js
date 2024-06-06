import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        proxy: {
            '/resource': 'http://127.0.0.1:80',
            '/api': {
                target: 'http://127.0.0.1:80',
                changeOrigin: true,
            },
            '/chat': {
                target: 'http://127.0.0.1:80',
                changeOrigin: true,
            },
        },
    },
})
