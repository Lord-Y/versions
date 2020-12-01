const { createProxyMiddleware } = require('http-proxy-middleware');

export default {
  install(app) {
    const proxyOptions = {
      target: process.env.API_URL,
      changeOrigin: true,
    };
    app.use('/api/', createProxyMiddleware(proxyOptions));
  },
};
