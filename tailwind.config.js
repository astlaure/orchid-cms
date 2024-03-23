/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './resources/assets/**/*.{css,js}',
    './resources/views/**/*.html',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

