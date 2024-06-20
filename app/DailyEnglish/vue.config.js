// vue.config.js
module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://47.107.81.75:8080',
        pathRewrite: {
          '^/api': ''
        }
      }
    },
  }
}