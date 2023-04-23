import Icons from "unplugin-icons/vite"
import { defineConfig } from "vite"
import { sveltekit } from "@sveltejs/kit/vite"

export default defineConfig({
  plugins: [
    sveltekit(),
    Icons({
      compiler: "svelte"
    })
  ],
  build: {
    emptyOutDir: true
  }
})
