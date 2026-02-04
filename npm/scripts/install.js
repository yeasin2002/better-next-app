#!/usr/bin/env node

const https = require('https');
const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');
const os = require('os');

const PACKAGE_VERSION = require('../package.json').version;
const REPO = 'yeasin2002/better-next-app';

function getPlatform() {
  const platform = os.platform();
  const arch = os.arch();

  if (platform === 'win32') {
    return { os: 'Windows', arch: 'x86_64', ext: '.exe' };
  }
  
  if (platform === 'darwin') {
    if (arch === 'arm64') {
      return { os: 'Darwin', arch: 'arm64', ext: '' };
    }
    return { os: 'Darwin', arch: 'x86_64', ext: '' };
  }
  
  if (platform === 'linux') {
    if (arch === 'arm64') {
      return { os: 'Linux', arch: 'arm64', ext: '' };
    }
    return { os: 'Linux', arch: 'x86_64', ext: '' };
  }

  throw new Error(`Unsupported platform: ${platform} ${arch}`);
}

function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    
    https.get(url, (response) => {
      if (response.statusCode === 302 || response.statusCode === 301) {
        // Follow redirect
        return download(response.headers.location, dest).then(resolve).catch(reject);
      }
      
      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: ${response.statusCode}`));
        return;
      }

      response.pipe(file);
      
      file.on('finish', () => {
        file.close();
        resolve();
      });
    }).on('error', (err) => {
      fs.unlink(dest, () => {});
      reject(err);
    });
  });
}

async function install() {
  try {
    const { os: osName, arch, ext } = getPlatform();
    const version = `v${PACKAGE_VERSION}`;
    
    // Construct download URL
    const archiveName = `better-next-app_${PACKAGE_VERSION}_${osName}_${arch}.tar.gz`;
    const downloadUrl = `https://github.com/${REPO}/releases/download/${version}/${archiveName}`;
    
    console.log(`Downloading better-next-app ${version} for ${osName} ${arch}...`);
    
    const binDir = path.join(__dirname, '..', 'bin');
    if (!fs.existsSync(binDir)) {
      fs.mkdirSync(binDir, { recursive: true });
    }

    const archivePath = path.join(binDir, archiveName);
    
    // Download the archive
    await download(downloadUrl, archivePath);
    
    console.log('Extracting binary...');
    
    // Extract the binary
    if (osName === 'Windows') {
      // For Windows, we need to handle .zip files
      const zipName = `better-next-app_${PACKAGE_VERSION}_${osName}_${arch}.zip`;
      const zipUrl = `https://github.com/${REPO}/releases/download/${version}/${zipName}`;
      const zipPath = path.join(binDir, zipName);
      
      await download(zipUrl, zipPath);
      
      // Extract using PowerShell
      execSync(`powershell -command "Expand-Archive -Path '${zipPath}' -DestinationPath '${binDir}' -Force"`, {
        stdio: 'inherit'
      });
      
      fs.unlinkSync(zipPath);
    } else {
      // Extract tar.gz for Unix systems
      execSync(`tar -xzf "${archivePath}" -C "${binDir}"`, {
        stdio: 'inherit'
      });
      
      fs.unlinkSync(archivePath);
    }
    
    const binaryName = `better-next-app${ext}`;
    const binaryPath = path.join(binDir, binaryName);
    
    // Make binary executable on Unix systems
    if (osName !== 'Windows') {
      fs.chmodSync(binaryPath, 0o755);
    }
    
    console.log('✓ Installation complete!');
  } catch (error) {
    console.error('Installation failed:', error.message);
    console.error('\nYou can manually download the binary from:');
    console.error(`https://github.com/${REPO}/releases/tag/v${PACKAGE_VERSION}`);
    process.exit(1);
  }
}

install();
