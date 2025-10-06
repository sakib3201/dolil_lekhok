import adapter from '@sveltejs/adapter-static';
import { paraglide } from '@inlang/paraglide-js-adapter-sveltekit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter({
      fallback: 'index.html'
    })
  },
  // Vite config should be at the root level, not inside kit
  vite: {
    plugins: [
      paraglide({
        project: './project.inlang',
        outdir: './src/paraglide'
      })
    ]
  }
};

export default config;