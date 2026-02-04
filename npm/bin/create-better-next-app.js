#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const os = require('os');

const platform = os.platform();
const ext = platform === 'win32' ? '.exe' : '';
const binaryName = `better-next-app${ext}`;
const binaryPath = path.join(__dirname, '..', 'bin', binaryName);

// Pass all arguments to the binary
const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: 'inherit',
  shell: platform === 'win32'
});

child.on('exit', (code) => {
  process.exit(code || 0);
});
