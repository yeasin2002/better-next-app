#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const os = require('os');
const fs = require("fs");

const platform = os.platform();
const ext = platform === 'win32' ? '.exe' : '';
const binaryName = `better-next-app${ext}`;
const binaryPath = path.join(__dirname, '..', 'bin', binaryName);

// Check if binary exists
if (!fs.existsSync(binaryPath)) {
  console.error('Error: Binary not found at', binaryPath);
  console.error('Please try reinstalling the package:');
  console.error('  npm install -g create-better-next-app@latest');
  process.exit(1);
}

// Pass all arguments to the binary
const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
  shell: platform === "win32",
  windowsHide: true,
});

child.on("error", (err) => {
  console.error("Failed to start binary:", err.message);
  process.exit(1);
});

child.on("exit", (code, signal) => {
  if (signal) {
    process.kill(process.pid, signal);
  } else {
    process.exit(code || 0);
  }
});
