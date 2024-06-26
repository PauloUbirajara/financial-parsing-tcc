import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  plugins: [sveltekit()],
  resolve: {
    alias: {
      $components: path.resolve(__dirname, "src/components"),
      $public: path.resolve(__dirname, "src/public"),
      $domain: path.resolve(__dirname, "src/domain"),
      $data: path.resolve(__dirname, "src/data"),
    },
  },
});
