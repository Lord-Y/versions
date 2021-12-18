export default {
  install(app, options) {
    this.$options = options || {};
  },

  beforeRender(response, context) {
    context.runtimeConfig = {
      ...(this.$options.publicConfig || {}),
      ...(this.$options.privateConfig || {}),
    };
  },

  beforeBuild(response, context) {
    context.data.runtimeConfig = {
      ...this.$options.publicConfig,
    };
  },
};
