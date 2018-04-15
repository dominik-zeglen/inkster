// TODO: rewrite all of this to ES6
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var path = require("path");
var webpack = require("webpack");

var resolve = path.resolve.bind(path, __dirname);

var extractTextPlugin;
var fileLoaderPath;
var output;
var reactPath;
var reactDomPath;

if (process.env.NODE_ENV === "production") {
  reactPath = "node_modules/react/umd/react.production.min.js";
  reactDomPath = "node_modules/react-dom/umd/react-dom.production.min.js";
  output = {
    path: resolve("static/"),
    filename: "[name].[chunkhash].js",
    chunkFilename: "[name].[chunkhash].js",
    publicPath: process.env.STATIC_URL || "/static/"
  };
  fileLoaderPath = "file-loader?name=[name].[hash].[ext]";
  extractTextPlugin = new ExtractTextPlugin("[name].[contenthash].css");
} else {
  reactPath = "node_modules/react/umd/react.development.js";
  reactDomPath = "node_modules/react-dom/umd/react-dom.development.js";
  output = {
    path: resolve("static/"),
    filename: "[name].js",
    chunkFilename: "[name].js",
    publicPath: "/static/"
  };
  fileLoaderPath = "file-loader?name=[name].[ext]";
  extractTextPlugin = new ExtractTextPlugin("[name].css");
}

var config = {
  entry: {
    app: "./panel/index.tsx"
  },
  output: output,
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "babel-loader"
      },
      {
        test: /\.tsx?$/,
        loader: "ts-loader"
      },
      {
        test: /\.(eot|otf|png|svg|jpg|ttf|woff|woff2)(\?v=[0-9.]+)?$/,
        loader: fileLoaderPath,
        include: [
          resolve("node_modules"),
          resolve("assets/fonts"),
          resolve("assets/images")
        ]
      }
    ]
  },
  plugins: [extractTextPlugin],
  resolve: {
    alias: {
      react: resolve(reactPath),
      "react-dom": resolve(reactDomPath)
    },
    extensions: [".ts", ".tsx", ".js", ".jsx"]
  },
  devtool: "sourceMap"
};

module.exports = config;
