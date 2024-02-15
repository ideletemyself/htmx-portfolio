const defaultTheme = require('tailwindcss/defaultTheme')
const { addDynamicIconSelectors } = require('@iconify/tailwind');

module.exports = {
  content: [
    "./static/*.{html,js}",
    "../templates/*.{html,js}",
  ],
  darkMode: 'class',
  theme: {
    screens: {
      'xs': '375px',
      ...defaultTheme.screens,
    },
    extend: {
       colors: {
        transparent: 'transparent',
        current: 'currentColor',
        'med-light-magenta': '#bd96c5',
        'cyan-blue': '#6d8dc4',
        'light-cyan-blue': '#a9b9d4',
        'behr-debonair-blue': '#d7dfe7',
        'very-light-brown': '#ece4da',
        'dark-green': '#1b401f',
      },
      fontFamily: {
        sans: ['Inter var', ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
    addDynamicIconSelectors(),
  ], 
}

