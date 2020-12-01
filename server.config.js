import { ExpressAdapter } from '@uvue/server';

export default {
  plugins: [
    '@uvue/server/plugins/gzip',
    '@uvue/server/plugins/serverError',
    '@uvue/server/plugins/static',
    './configs/proxy',
    './configs/statics',
    [
      './configs/runtimeConfigServer',
      {
        publicConfig: {
          BASE_URL: process.env.BASE_URL
            ? process.env.BASE_URL
            : 'http://localhost:8080',
          API_URL: process.env.API_URL
            ? process.env.API_URL
            : 'http://localhost:8081',
          RANGE_LIMIT: process.env.RANGE_LIMIT ? process.env.RANGE_LIMIT : 50,
        },
        privateConfig: {
          BASE_URL: process.env.BASE_URL
            ? process.env.BASE_URL
            : 'http://localhost:8080',
          API_URL: process.env.API_URL
            ? process.env.API_URL
            : 'http://localhost:8081',
          RANGE_LIMIT: process.env.RANGE_LIMIT ? process.env.RANGE_LIMIT : 50,
        },
      },
    ],
  ],
  adapter: ExpressAdapter,
  watch: ['server.config.js', 'vue.config.js', './configs/*.js'],
  watchIgnore: ['node_modules'],
};
