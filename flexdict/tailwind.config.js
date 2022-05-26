module.exports = {
  content: [
    './components/**/*.{js,vue,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.{js,ts}',
    './nuxt.config.{js,ts}',
  ],
  theme: {
    extend: {
      colors: {
        primary: '#AEE0FF',
        secondary: '#10204D',
        tertiary: '#FF9800',
        'lightest-gray': '#F8F7FA',
        'light-gray': '#CBCBD4',
        'medium-gray': '#A6A6AA',
        'dark-gray': '#3A3A3A',
      },
      fontFamily: {
        sans: ['PT Serif Caption', 'sans-serif'],
        serif: ['Otomanopee One', 'serif'],
      },
    },
  },
  plugins: [],
}
