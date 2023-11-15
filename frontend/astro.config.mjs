import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";
import svelte from "@astrojs/svelte";

// https://astro.build/config
export default defineConfig({
  integrations: [tailwind(), svelte()],
  vite: {
    server: {
      proxy: {
        "/api": {
          target: "http://localhost:8080",
          rewrite: (path) => path.replace(/^\/api/, ""),
        },
      },
    },
  },
});
