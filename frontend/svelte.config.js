import adapter from "@sveltejs/adapter-static"
import { vitePreprocess } from "@sveltejs/kit/vite"

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: vitePreprocess(),
  kit: {
    adapter: adapter({
      fallback: "index.html",
      pages: "dist",
      assets: "dist",
      strict: true
    }),
    alias: {
      $wails: "src/lib/wailsjs"
    }
  }
}

export default config
