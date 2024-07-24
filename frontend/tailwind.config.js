/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line no-undef
const colors = require('tailwindcss/colors')

export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      'black': '#232536',
      'light': '#fff',
      'yellow': '#FFD050',
      'light-yellow': '#FBF6EA',
      'purple': '#592EA9',
      'lavender': '#F4F0F8',
      'gray': colors.gray,
    },
    extend: {},
  },
  plugins: [],
}

