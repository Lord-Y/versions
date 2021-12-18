const express = require('express');

export default {
  install(app) {
    app.use('/health', (req, res) => {
      res.send('OK');
    });
  },
};
