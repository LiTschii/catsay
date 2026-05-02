#!/usr/bin/env node
'use strict';

// This shim locates the downloaded native binary and execs it,
// passing through all arguments. Works with node, npx, bunx.

const path = require('path');
const os = require('os');
const { spawnSync } = require('child_process');

const isWindows = os.platform() === 'win32';
const bin = path.join(__dirname, isWindows ? 'catsay.exe' : 'catsay');

const result = spawnSync(bin, process.argv.slice(2), { stdio: 'inherit' });
process.exit(result.status ?? 1);
