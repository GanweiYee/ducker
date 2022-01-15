module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: process.env.VUE_APP_API_URL,
                pathRewrite: {
                    '^/api':'/v1'
                }
            }
        }
    }
}