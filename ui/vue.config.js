var path = require('path');
const TerserPlugin = require('terser-webpack-plugin');

module.exports = {
  lintOnSave: process.env.NODE_ENV !== 'production',
  productionSourceMap: process.env.NODE_ENV === 'production' ? false : true,
  configureWebpack: {
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src/'),
        '@router': path.resolve(__dirname, 'src/router'),
        '@store': path.resolve(__dirname, 'src/store'),
        '@components': path.resolve(__dirname, 'src/components'),
        '@views': path.resolve(__dirname, 'src/views'),
        '@filters': path.resolve(__dirname, 'src/filters'),
        '@plugins': path.resolve(__dirname, 'src/plugins'),
        '@public': path.resolve(__dirname, './public'),
      },
    },
    optimization: {
      minimizer: [
        new TerserPlugin({
          parallel: true,
          cache: true,
          terserOptions: {
            ecma: 6,
            parallel: true,
            sourceMap: true,
            compress: {
              drop_debugger:
                process.env.NODE_ENV === 'production' ? false : true,
              drop_console:
                process.env.NODE_ENV === 'production' ? false : true,
              arrows: false,
              collapse_vars: false,
              comparisons: false,
              computed_props: false,
              hoist_funs: false,
              hoist_props: false,
              hoist_vars: false,
              inline: false,
              loops: false,
              negate_iife: false,
              properties: false,
              reduce_funcs: false,
              reduce_vars: false,
              switches: false,
              toplevel: false,
              typeofs: false,
              booleans: true,
              if_return: true,
              sequences: true,
              unused: true,
              conditionals: true,
              dead_code: true,
              evaluate: true,
            },
            mangle: {
              safari10: true,
            },
            output: { comments: false, beautify: false },
          },
          extractComments: false,
        }),
      ],
    },
  },
  devServer: {
    host: '0.0.0.0',
    hot: true,
    hotOnly: true,
  },
};
