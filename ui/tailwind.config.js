const flattenColorPalette =
  require('tailwindcss/lib/util/flattenColorPalette').default

/* eslint-env node */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      animation: {
        'spin-fast': 'spin 0.60s linear infinite',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    ({ addUtilities, theme, variants }) => {
      const colors = flattenColorPalette(theme('borderColor'))
      delete colors['default']

      const colorMap = Object.keys(colors).map((color) => ({
        [`.border-t-${color}`]: { borderTopColor: colors[color] },
        [`.border-r-${color}`]: { borderRightColor: colors[color] },
        [`.border-b-${color}`]: { borderBottomColor: colors[color] },
        [`.border-l-${color}`]: { borderLeftColor: colors[color] },
      }))
      const utilities = Object.assign({}, ...colorMap)
      addUtilities(utilities, variants('borderColor'))
    },
  ],
}
