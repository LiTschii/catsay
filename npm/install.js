#!/usr/bin/env node
'use strict';

const https = require('https');
const fs = require('fs');
const path = require('path');
const os = require('os');
const { execSync } = require('child_process');
const { version } = require('./package.json');

const REPO = 'LiTLiTschi/catsay';
const BIN_DIR = path.join(__dirname, 'bin');
const platform = os.platform();
const arch = os.arch() === 'arm64' ? 'arm64' : 'amd64';

fs.mkdirSync(BIN_DIR, { recursive: true });

const base = `https://github.com/${REPO}/releases/download/v${version}`;

let url, dest, isWindows = platform === 'win32';

if (isWindows) {
  url = `${base}/catsay.exe`;
  dest = path.join(BIN_DIR, 'catsay.exe');
} else {
  const plat = platform === 'darwin' ? 'darwin' : 'linux';
  url = `${base}/catsay-${plat}-${arch}.tar.gz`;
  dest = path.join(BIN_DIR, 'catsay.tar.gz');
}

function download(url, dest, cb) {
  const file = fs.createWriteStream(dest);
  https.get(url, res => {
    if (res.statusCode === 302 || res.statusCode === 301) {
      file.close();
      fs.unlinkSync(dest);
      return download(res.headers.location, dest, cb);
    }
    res.pipe(file);
    file.on('finish', () => file.close(cb));
  }).on('error', err => {
    fs.unlinkSync(dest);
    console.error('Download failed:', err.message);
    process.exit(1);
  });
}

console.log(`Downloading catsay v${version} for ${platform}/${arch}...`);

download(url, dest, () => {
  if (!isWindows) {
    execSync(`tar -xzf ${dest} -C ${BIN_DIR}`);
    fs.unlinkSync(dest);
    fs.chmodSync(path.join(BIN_DIR, 'catsay'), 0o755);
  }
  console.log('catsay installed successfully.');
});
