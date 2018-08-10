const webpack = require('webpack')
const path = require('path')

module.exports = production => ({
  entry: {
    build: './src/index.js'
  },
  output: {
    filename: 'js/[name].js',
    path: path.resolve(__dirname, './build/')
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      },
      {
        test: /\.styl$/, 
        loader: 'style-loader!css-loader!stylus-loader' 
      }
    ]
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': production ? '"production"' : '"development"'
    })
  ],
  optimization: {
    minimize: production
  },
  devtool: (production) ? 'none' : 'source-map',
  mode: (production) ? 'production' : 'development',
  watch: !production
})