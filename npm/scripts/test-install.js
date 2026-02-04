#!/usr/bin/env node

/**
 * Test script to verify the installation process locally
 * Run with: node npm/scripts/test-install.js
 */

const os = require('os');
const path = require('path');

function getPlatform() {
  const platform = os.platform();
  const arch = os.arch();

  let osName, archName, ext, format;

  if (platform === 'win32') {
    osName = 'Windows';
    archName = 'x86_64';
    ext = '.exe';
    format = 'zip';
  } else if (platform === 'darwin') {
    osName = 'Darwin';
    archName = arch === 'arm64' ? 'arm64' : 'x86_64';
    ext = '';
    format = 'tar.gz';
  } else if (platform === 'linux') {
    osName = 'Linux';
    archName = arch === 'arm64' ? 'arm64' : 'x86_64';
    ext = '';
    format = 'tar.gz';
  } else {
    throw new Error(`Unsupported platform: ${platform} ${arch}`);
  }

  return { os: osName, arch: archName, ext, format };
}

console.log('Platform Detection Test');
console.log('======================\n');

const { os: osName, arch, ext, format } = getPlatform();
const version = require('../package.json').version;

console.log(`Detected Platform: ${osName}`);
console.log(`Architecture: ${arch}`);
console.log(`Binary Extension: ${ext || '(none)'}`);
console.log(`Archive Format: ${format}`);
console.log(`Package Version: ${version}\n`);

const archiveExt = format === 'zip' ? '.zip' : '.tar.gz';
const archiveName = `better-next-app_${version}_${osName}_${arch}${archiveExt}`;
const binaryName = `better-next-app${ext}`;

console.log('Expected Files:');
console.log(`  Archive: ${archiveName}`);
console.log(`  Binary: ${binaryName}\n`);

const downloadUrl = `https://github.com/yeasin2002/better-next-app/releases/download/v${version}/${archiveName}`;
console.log('Download URL:');
console.log(`  ${downloadUrl}\n`);

console.log('Binary Path:');
console.log(`  ${path.join(__dirname, '..', 'bin', binaryName)}\n`);

console.log('✓ Platform detection working correctly!');
