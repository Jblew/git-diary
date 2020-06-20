// Configuration for development server only.
// To update proxies for production use 'firebase.json#hosting.rewrite'

module.exports = {
  devServer: {
    proxy: {
      "^/publish": {
        logLevel: "debug",
        pathRewrite: { "^/publish": "/PublishEntry" },
        target: "https://us-central1-git-diary.cloudfunctions.net/",
      },
    },
  },
};
