# NPM Publishing Setup Script (Windows PowerShell)
# This script helps you set up NPM publishing for the first time

$ErrorActionPreference = "Stop"

Write-Host "🚀 Better Next App - NPM Publishing Setup" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Check if npm is installed
try {
    $null = Get-Command npm -ErrorAction Stop
    Write-Host "✓ npm is installed" -ForegroundColor Green
} catch {
    Write-Host "❌ Error: npm is not installed" -ForegroundColor Red
    Write-Host "Please install Node.js and npm first: https://nodejs.org/"
    exit 1
}

Write-Host ""

# Check if user is logged in to npm
Write-Host "Checking NPM authentication..."
try {
    $npmUser = npm whoami 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✓ You are logged in as: $npmUser" -ForegroundColor Green
    } else {
        throw "Not logged in"
    }
} catch {
    Write-Host "⚠️  You are not logged in to NPM" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Please log in to NPM:"
    npm login
    
    $npmUser = npm whoami 2>$null
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✓ Successfully logged in as: $npmUser" -ForegroundColor Green
    } else {
        Write-Host "❌ Login failed" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""
Write-Host "Checking package availability..."

# Check if package name is available
$packageName = "create-better-next-app"
$packageExists = $false

try {
    npm view $packageName 2>$null | Out-Null
    if ($LASTEXITCODE -eq 0) {
        $packageExists = $true
    }
} catch {
    $packageExists = $false
}

if ($packageExists) {
    $currentOwner = (npm view $packageName maintainers --json | ConvertFrom-Json)[0].name
    Write-Host "⚠️  Package '$packageName' already exists" -ForegroundColor Yellow
    Write-Host "   Current owner: $currentOwner"
    
    if ($currentOwner -eq $npmUser) {
        Write-Host "✓ You own this package - you can publish updates" -ForegroundColor Green
    } else {
        Write-Host ""
        Write-Host "❌ You don't own this package. Options:" -ForegroundColor Red
        Write-Host "   1. Use a scoped package: @$npmUser/$packageName"
        Write-Host "   2. Choose a different package name"
        Write-Host ""
        Write-Host "To use a scoped package, update npm/package.json:"
        Write-Host "   `"name`": `"@$npmUser/$packageName`""
        exit 1
    }
} else {
    Write-Host "✓ Package name '$packageName' is available" -ForegroundColor Green
}

Write-Host ""
Write-Host "Testing package build..."

Set-Location npm

# Test installation script
Write-Host "Running postinstall script..."
try {
    node scripts/install.js
    Write-Host "✓ Installation script works" -ForegroundColor Green
} catch {
    Write-Host "❌ Installation script failed" -ForegroundColor Yellow
    Write-Host "This is expected if no GitHub release exists yet"
}

Write-Host ""
Write-Host "Checking package contents..."
npm pack --dry-run

Write-Host ""
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "✅ Setup Complete!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:"
Write-Host ""
Write-Host "1. Add NPM_TOKEN to GitHub Secrets:"
Write-Host "   - Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions"
Write-Host "   - Create a new secret named: NPM_TOKEN"
Write-Host "   - Get your token from: https://www.npmjs.com/settings/$npmUser/tokens"
Write-Host ""
Write-Host "2. Create your first release:"
Write-Host "   git tag -a v0.1.0 -m 'Release v0.1.0'"
Write-Host "   git push origin v0.1.0"
Write-Host ""
Write-Host "3. GitHub Actions will automatically:"
Write-Host "   - Build binaries with GoReleaser"
Write-Host "   - Publish to NPM registry"
Write-Host ""
Write-Host "Or publish manually now:"
Write-Host "   cd npm && npm publish --access public"
Write-Host ""
