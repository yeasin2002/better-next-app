#!/bin/bash

# NPM Publishing Setup Script
# This script helps you set up NPM publishing for the first time

set -e

echo "🚀 Better Next App - NPM Publishing Setup"
echo "=========================================="
echo ""

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    echo "❌ Error: npm is not installed"
    echo "Please install Node.js and npm first: https://nodejs.org/"
    exit 1
fi

echo "✓ npm is installed"
echo ""

# Check if user is logged in to npm
echo "Checking NPM authentication..."
if npm whoami &> /dev/null; then
    NPM_USER=$(npm whoami)
    echo "✓ You are logged in as: $NPM_USER"
else
    echo "⚠️  You are not logged in to NPM"
    echo ""
    echo "Please log in to NPM:"
    npm login
    
    if npm whoami &> /dev/null; then
        NPM_USER=$(npm whoami)
        echo "✓ Successfully logged in as: $NPM_USER"
    else
        echo "❌ Login failed"
        exit 1
    fi
fi

echo ""
echo "Checking package availability..."

# Check if package name is available
PACKAGE_NAME="create-better-next-app"
if npm view $PACKAGE_NAME &> /dev/null; then
    CURRENT_OWNER=$(npm view $PACKAGE_NAME maintainers --json | grep -o '"name":"[^"]*"' | head -1 | cut -d'"' -f4)
    echo "⚠️  Package '$PACKAGE_NAME' already exists"
    echo "   Current owner: $CURRENT_OWNER"
    
    if [ "$CURRENT_OWNER" = "$NPM_USER" ]; then
        echo "✓ You own this package - you can publish updates"
    else
        echo ""
        echo "❌ You don't own this package. Options:"
        echo "   1. Use a scoped package: @$NPM_USER/$PACKAGE_NAME"
        echo "   2. Choose a different package name"
        echo ""
        echo "To use a scoped package, update npm/package.json:"
        echo '   "name": "@'$NPM_USER'/'$PACKAGE_NAME'"'
        exit 1
    fi
else
    echo "✓ Package name '$PACKAGE_NAME' is available"
fi

echo ""
echo "Testing package build..."

cd npm

# Test installation script
echo "Running postinstall script..."
if node scripts/install.js; then
    echo "✓ Installation script works"
else
    echo "❌ Installation script failed"
    echo "This is expected if no GitHub release exists yet"
fi

echo ""
echo "Checking package contents..."
npm pack --dry-run

echo ""
echo "=========================================="
echo "✅ Setup Complete!"
echo ""
echo "Next steps:"
echo ""
echo "1. Add NPM_TOKEN to GitHub Secrets:"
echo "   - Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions"
echo "   - Create a new secret named: NPM_TOKEN"
echo "   - Get your token from: https://www.npmjs.com/settings/$NPM_USER/tokens"
echo ""
echo "2. Create your first release:"
echo "   git tag -a v0.1.0 -m 'Release v0.1.0'"
echo "   git push origin v0.1.0"
echo ""
echo "3. GitHub Actions will automatically:"
echo "   - Build binaries with GoReleaser"
echo "   - Publish to NPM registry"
echo ""
echo "Or publish manually now:"
echo "   cd npm && npm publish --access public"
echo ""
