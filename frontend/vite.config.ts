import {defineConfig} from "vite"
import react from "@vitejs/plugin-react"

// https://vitejs.dev/config/
// noinspection JSUnusedGlobalSymbols
export default defineConfig({
    build: {
        manifest: true,
        rollupOptions: {
            input: "src/main.tsx"
        },
        modulePreload: {
            polyfill: false
        }
    },
    server: {
        host: "0.0.0.0",
        port: 5173,
        strictPort: true,
        hmr: {
            protocol: "ws",
            port: 5173,
        }
    },
    plugins: [react()],
    assetsInclude: ["**/*.gltf"],
})
