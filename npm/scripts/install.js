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

  let osName, archName, ext, format;

  if (platform === "win32") {
    osName = "Windows";
    archName = "x86_64";
    ext = ".exe";
    format = "zip";
  } else if (platform === "darwin") {
    osName = "Darwin";
    archName = arch === "arm64" ? "arm64" : "x86_64";
    ext = "";
    format = "tar.gz";
  } else if (platform === "linux") {
    osName = "Linux";
    archName = arch === "arm64" ? "arm64" : "x86_64";
    ext = "";
    format = "tar.gz";
  } else {
    throw new Error(`Unsupported platform: ${platform} ${arch}`);
  }

  return { os: osName, arch: archName, ext, format };
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

function extractZip(archivePath, destDir) {
  // Use PowerShell with .NET Framework method (more compatible)
  const psScript = `
    Add-Type -AssemblyName System.IO.Compression.FileSystem
    [System.IO.Compression.ZipFile]::ExtractToDirectory('${archivePath.replace(/\\/g, "\\\\")}', '${destDir.replace(/\\/g, "\\\\")}')
  `;

  try {
    execSync(
      `powershell -NoProfile -ExecutionPolicy Bypass -Command "${psScript}"`,
      {
        stdio: "inherit",
        windowsHide: true,
      },
    );
  } catch (error) {
    // Fallback: try using tar command (available in Windows 10+)
    try {
      execSync(`tar -xf "${archivePath}" -C "${destDir}"`, {
        stdio: "inherit",
      });
    } catch (tarError) {
      throw new Error("Failed to extract zip file. Please extract manually.");
    }
  }
}

async function install() {
  try {
    const { os: osName, arch, ext, format } = getPlatform();
    const version = `v${PACKAGE_VERSION}`;

    // Construct archive name and download URL
    const archiveExt = format === "zip" ? ".zip" : ".tar.gz";
    const archiveName = `better-next-app_${PACKAGE_VERSION}_${osName}_${arch}${archiveExt}`;
    const downloadUrl = `https://github.com/${REPO}/releases/download/${version}/${archiveName}`;

    console.log(
      `Downloading better-next-app ${version} for ${osName} ${arch}...`,
    );
    console.log(`URL: ${downloadUrl}`);

    const binDir = path.join(__dirname, "..", "bin");
    if (!fs.existsSync(binDir)) {
      fs.mkdirSync(binDir, { recursive: true });
    }

    const archivePath = path.join(binDir, archiveName);

    // Download the archive
    await download(downloadUrl, archivePath);

    console.log("Extracting binary...");

    // Extract the binary based on format
    if (format === "zip") {
      // Windows: Extract zip
      extractZip(archivePath, binDir);
    } else {
      // Unix: Extract tar.gz
      execSync(`tar -xzf "${archivePath}" -C "${binDir}"`, {
        stdio: "inherit",
      });
    }

    // Clean up archive
    fs.unlinkSync(archivePath);

    const binaryName = `better-next-app${ext}`;
    const binaryPath = path.join(binDir, binaryName);

    // Verify binary exists
    if (!fs.existsSync(binaryPath)) {
      throw new Error(`Binary not found after extraction: ${binaryPath}`);
    }

    // Make binary executable on Unix systems
    if (format !== "zip") {
      fs.chmodSync(binaryPath, 0o755);
    }

    console.log("✓ Installation complete!");
    console.log(`Binary installed at: ${binaryPath}`);
  } catch (error) {
    console.error('Installation failed:', error.message);
    console.error('\nYou can manually download the binary from:');
    console.error(`https://github.com/${REPO}/releases/tag/v${PACKAGE_VERSION}`);
    process.exit(1);
  }
}

install();
