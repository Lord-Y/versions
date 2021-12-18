export default {
  beforeCreate(context, inject) {
    const config = process.server
      ? context.ssr.runtimeConfig
      : window.__DATA__.runtimeConfig;

    inject('config', config);
  },
};
