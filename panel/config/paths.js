"use strict";

const path = require("path");
const fs = require("fs");
const url = require("url");

// Make sure any symlinks in the project folder are resolved:
// https://github.com/facebookincubator/create-react-app/issues/637
const appDirectory = fs.realpathSync(process.cwd());
const resolveApp = relativePath => path.resolve(appDirectory, relativePath);

const envPublicUrl = process.env.PUBLIC_URL;

function ensureSlash(pathName, needsSlash) {
  const hasSlash = pathName.endsWith("/");
  if (hasSlash && !needsSlash) {
    return pathName.substr(pathName, pathName.length - 1);
  } else if (!hasSlash && needsSlash) {
    return `${pathName}/`;
  } else {
    return pathName;
  }
}

const getPublicUrl = appPackageJson =>
  envPublicUrl || require(appPackageJson).homepage;

// We use `PUBLIC_URL` environment variable or "homepage" field to infer
// "public path" at which the app is served.
// Webpack needs to know it to put the right <script> hrefs into HTML even in
// single-page apps that may serve index.html for nested URLs like /todos/42.
// We can't use a relative path in HTML because we don't want to load something
// like /todos/42/static/js/bundle.7289d.js. We have to know the root.
function getServedPath(appPackageJson) {
  const publicUrl = getPublicUrl(appPackageJson);
  const servedUrl =
    envPublicUrl || (publicUrl ? url.parse(publicUrl).pathname : "/");
  return ensureSlash(servedUrl, true);
}

// config after eject: we're in ./config/
module.exports = {
  appBuild: resolveApp("panel/build"),
  appHtml: resolveApp("panel/public/index.html"),
  appIndexJs: resolveApp("panel/src/index.tsx"),
  appNodeModules: resolveApp("node_modules"),
  appPackageJson: resolveApp("package.json"),
  appPublic: resolveApp("panel/public"),
  appSrc: resolveApp("panel/src"),
  appTsConfig: resolveApp("tsconfig.json"),
  appTsLint: resolveApp("tslint.json"),
  appTsProdConfig: resolveApp("tsconfig.prod.json"),
  dotenv: resolveApp(".env"),
  publicUrl: getPublicUrl(resolveApp("package.json")),
  servedPath: getServedPath(resolveApp("package.json")),
  testsSetup: resolveApp("panel/src/setupTests.ts"),
  yarnLockFile: resolveApp("yarn.lock")
};
