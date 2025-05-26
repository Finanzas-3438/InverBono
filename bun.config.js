const fs = require('fs');
const path = require("path");
const vuePlugin = require("esbuild-plugin-vue3");
const tailwind = require("tailwindcss");
const autoprefixer = require("autoprefixer");
const postCssPlugin = require("@wemake/esbuild-plugin-postcss");
const esbuild = require("esbuild");

const BUILD_OPTIONS = {
  logLevel: 'info',
  entryPoints: ['./public/ts/application.ts'],
  bundle: true,
  minify: true,
  minifyWhitespace: true, //avoid leaking routes
  sourcemap: false,
  define: {
    'process.env.NODE_ENV': '"production"',
    'process.env.VERSION': '"1.0.0"'
  },
  // sourcemap: 'inline',
  // sourceRoot: '/js',
  outfile: 'public/js/app.js',
  external: ["/img/*", "/fonts/*"],
  plugins: [
    // Use Vue with Tailwind
    vuePlugin({
      postcss: {
        plugins: [tailwind, autoprefixer]
      }
      }
    ),
    // Use PostCSS with Tailwind
    postCssPlugin({
      plugins: [
        tailwind,
        autoprefixer
      ],
    processOptions: { syntax: require('postcss-scss') },
    }),
  ],
  metafile: true,
}

async function build() {
  try {
    let result = await esbuild.build(BUILD_OPTIONS) 
    fs.writeFileSync('meta.json', JSON.stringify(result.metafile))
  } catch(e) {
    console.info("Error building, try again...", e)
    // process.exit(1)
  }
}

module.exports = { BUILD_OPTIONS };
watchv2();


// const config = {
//   sourcemap: "external",
//   entrypoints: ["public/ts/application.ts"],
//   outdir: path.join(process.cwd(), "public/js/"),
// };

// const build = async (config) => {
//   const result = await Bun.build(config);

//   if (!result.success) {
//     if (process.argv.includes('--watch')) {
//       console.error("Build failed");
//       for (const message of result.logs) {
//         console.error(message);
//       }
//       return;
//     } else {
//       throw new AggregateError(result.logs, "Build failed");
//     }
//   }
// };

async function watchv2(){
  if (process.argv.includes('--watch')) {

    BUILD_OPTIONS.minify = false;
    BUILD_OPTIONS.minifyWhitespace = false;
    BUILD_OPTIONS.sourcemap = true;
    BUILD_OPTIONS.define['process.env.NODE_ENV'] = '"development"';

    await build();

    const watchDirs = [
      path.join(process.cwd(), "pkg/web/views"),
      path.join(process.cwd(), "public/ts"),
    ];

    watchDirs.forEach(dir => {
      fs.watch(dir, { recursive: true }, (eventType, filename) => {
        console.log(`File changed: ${filename} in ${dir}. Rebuilding...`);
        build();
      });
    });
  } else {
    // build once if not in watch mode
    await build(); 
    process.exit(0);
  }
}

