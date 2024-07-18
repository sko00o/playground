const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/py/': {
        target: 'http://backend:8080',
      },
    }
  }
})
