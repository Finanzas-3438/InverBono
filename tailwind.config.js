/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pkg/web/views/**/*.{html,js}",
    "./public/ts/**/*.{ts,js,vue}",
    "./public/pages/**/*.html",
  ],
  theme: {
    extend: {
      colors: {
        
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/container-queries'),
  ],
}
