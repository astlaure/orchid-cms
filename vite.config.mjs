import viteReact from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

export default defineConfig({
    plugins: [viteReact()],
    publicDir: 'resources/assets/public',
    build: {
        copyPublicDir: true,
        rollupOptions: {
            input: 'resources/assets/main.js',
            output: {
                dir: 'public',
                entryFileNames: '[name].js',
                assetFileNames: '[name].[ext]',
            },
        }
    },
});
