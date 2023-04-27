/** @type {import('tailwindcss').Config} */

export default {
  content: [
    'index.html',
    './src/**/*.{js,ts,jsx,tsx,vue}'
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        dark: '#071108',
        'dark-dim': '#222322',
        light: '#DFE0E2',
        'light-dim': '#EEF1F7',
        'cobalt-green': '#A2E3C4'
      }
    }
  },
  plugins: []
};
