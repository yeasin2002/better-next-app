# NPM Publishing Setup

This guide explains how to publish `create-better-next-app` to NPM registry.

## Overview

The NPM package is a lightweight wrapper that:
1. Downloads the appropriate binary from GitHub releases during `npm install`
2. Provides a Node.js shim to execute the binary
3. Works with `npx`, `pnpm dlx`, `bunx`, etc.

## One-Time Setup

### 1. Create NPM Account

If you don't have an NPM account:
1. Go to https://www.npmjs.com/signup
2. Create an account
3. Verify your email

### 2. Generate NPM Access Token

1. Log in to https://www.npmjs.com
2. Click your profile picture → "Access Tokens"
3. Click "Generate New Token" → "Classic Token"
4. Select "Automation" type (for CI/CD)
5. Copy the token (starts with `npm_...`)

### 3. Add NPM Token to GitHub Secrets

1. Go to your GitHub repository
2. Navigate to Settings → Secrets and variables → Actions
3. Click "New repository secret"
4. Name: `NPM_TOKEN`
5. Value: Paste your NPM token
6. Click "Add secret"

### 4. Claim Package Name (First Time Only)

Before your first release, publish manually to claim the package name:

```bash
cd npm

# Login to NPM
npm login

# Publish (this claims the package name)
npm publish --access public
```

After this, GitHub Actions will handle all future releases automatically.

## How It Works

### Package Structure

```
npm/
├── bin/
│   └── create-better-next-app.js    # Node.js wrapper script
├── scripts/
│   └── install.js                    # Post-install script (downloads binary)
├── package.json                      # NPM package metadata
├── README.md                         # NPM package documentation
└── .npmignore                        # Files to exclude from NPM
```

### Installation Flow

When a user runs `npx create-better-next-app@latest`:

1. NPM downloads the package
2. `postinstall` script runs (`scripts/install.js`)
3. Script detects OS and architecture
4. Downloads the correct binary from GitHub releases
5. Extracts binary to `npm/bin/`
6. Makes binary executable (Unix systems)
7. User can now run the CLI

### Execution Flow

When the user runs the command:

1. Node.js executes `bin/create-better-next-app.js`
2. Script spawns the Go binary with user's arguments
3. Binary runs and creates the Next.js project

## Release Process

### Automated Release (Recommended)

When you push a git tag, GitHub Actions automatically:

1. Builds binaries with GoReleaser
2. Creates GitHub release
3. Updates NPM package version
4. Publishes to NPM registry

```bash
# Create and push a tag
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

# GitHub Actions handles the rest!
```

### Manual Release (If Needed)

If you need to publish manually:

```bash
cd npm

# Update version to match git tag
npm version 0.1.0 --no-git-tag-version

# Login to NPM (if not already logged in)
npm login

# Publish
npm publish --access public
```

## Version Synchronization

**Important:** The NPM package version must match the GitHub release version.

- Git tag: `v0.1.0`
- NPM version: `0.1.0` (without the `v` prefix)

The GitHub Actions workflow automatically handles this synchronization.

## Testing Locally

Before publishing, test the NPM package locally:

```bash
cd npm

# Install dependencies (downloads binary)
npm install

# Test the CLI
node bin/create-better-next-app.js --help

# Or test with npx
npx . my-test-app
```

## Troubleshooting

### "Package name already taken"

If someone else claimed the package name:
1. Choose a different name (e.g., `@yeasin2002/create-better-next-app`)
2. Update `npm/package.json` → `name` field
3. Update documentation

### "Binary not found" error

The postinstall script failed to download the binary:
1. Check that the GitHub release exists
2. Verify the version matches
3. Check the binary naming in GoReleaser config
4. Test the download URL manually

### "Permission denied" on Unix

The binary isn't executable:
1. Check `scripts/install.js` sets permissions (`chmod 0o755`)
2. Manually fix: `chmod +x npm/bin/better-next-app`

### NPM publish fails in GitHub Actions

1. Verify `NPM_TOKEN` secret is set correctly
2. Check token hasn't expired
3. Ensure token has "Automation" permissions
4. Verify package name isn't taken

## Package Maintenance

### Updating the Package

To update package metadata (description, keywords, etc.):

1. Edit `npm/package.json`
2. Commit changes
3. Create a new release (version bump)

### Deprecating Old Versions

```bash
npm deprecate create-better-next-app@0.0.1 "Please upgrade to 0.1.0"
```

### Unpublishing (Use Carefully!)

You can only unpublish within 72 hours:

```bash
npm unpublish create-better-next-app@0.0.1
```

## Security Best Practices

1. **Never commit NPM tokens** to the repository
2. **Use Automation tokens** for CI/CD (not your personal token)
3. **Enable 2FA** on your NPM account
4. **Rotate tokens** periodically
5. **Review package contents** before publishing:
   ```bash
   npm pack --dry-run
   ```

## Monitoring

After publishing, monitor:

1. **NPM downloads**: https://www.npmjs.com/package/create-better-next-app
2. **GitHub releases**: https://github.com/yeasin2002/better-next-app/releases
3. **Issues**: Users reporting installation problems

## Future Enhancements

Consider adding:

1. **Scoped package**: `@yeasin2002/create-better-next-app` for namespace control
2. **Provenance**: NPM provenance for supply chain security
3. **Binary caching**: Cache downloaded binaries to speed up reinstalls
4. **Fallback mirrors**: Alternative download sources if GitHub is down
5. **Version checking**: Warn users about outdated versions

## Resources

- [NPM Publishing Guide](https://docs.npmjs.com/packages-and-modules/contributing-packages-to-the-registry)
- [NPM Access Tokens](https://docs.npmjs.com/about-access-tokens)
- [GitHub Actions NPM Publishing](https://docs.github.com/en/actions/publishing-packages/publishing-nodejs-packages)
