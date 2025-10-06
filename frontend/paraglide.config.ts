import { defineConfig } from "@inlang/paraglide-js";

export default defineConfig({
  outdir: "./src/paraglide",
  project: "./project.inlang",
  languages: ["en", "bn"],
  // Enable Vite plugin for HMR support
  vite: true
});
