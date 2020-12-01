module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: ['plugin:vue/strongly-recommended', '@vue/prettier'],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    'no-console': 'off',
    'no-debugger': 'off',
    'vue/no-use-v-if-with-v-for': 'off',
  },
};
