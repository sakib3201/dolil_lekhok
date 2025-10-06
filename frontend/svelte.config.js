import adapter from '@sveltejs/adapter-static';
import { paraglide } from '@inlang/paraglide-js-adapter-sveltekit/vite';

export default {
  kit: {
    adapter: adapter({
      fallback: 'index.html'
    }),
    // Add Vite configuration
    vite: {
      plugins: [
        paraglide({
          // Point to your project directory
          project: './project.inlang',
          // Output directory for generated files
          outdir: './src/paraglide'
        })
      ]
    }
  }
};