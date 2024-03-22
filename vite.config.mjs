import { defineConfig } from 'vite';

export default defineConfig({
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
