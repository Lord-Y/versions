export default {
  plugins: [
    '@uvue/core/plugins/asyncData',
    '@uvue/core/plugins/prefetch',
    [
      '@uvue/core/plugins/vuex',
      {
        onHttpRequest: true,
        fetch: true,
      },
    ],
    '@uvue/core/plugins/middlewares',
    '@uvue/core/plugins/errorHandler',
    '../../configs/runtimeConfig',
  ],
};
