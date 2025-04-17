module.exports = {
  devServer: {
    port: 8086,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    },
    client: {
      webSocketURL: {
        hostname: 'localhost',
        pathname: '/dev-server-ws',
        port: 8086
      }
    }
  }
} 