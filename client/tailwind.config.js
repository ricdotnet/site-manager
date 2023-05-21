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
        dark: '#181719',
        'dark-darker': '#0B0B0F',
        'dark-dim': '#6B696C',
        'dark-border': '#2B2C2D',
        light: '#F1F3F8',
        'light-dim': '#AEAFB1',
        'light-border': '#C3C4CC',
        'cobalt-green': '#A2E3C4'
      }
    }
  },
  plugins: []
};
