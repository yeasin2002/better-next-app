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

#### Step 2.1: Navigate to Access Tokens

1. Log in to https://www.npmjs.com
2. Click your profile picture (top right)
3. Click "Access Tokens"
4. Or go directly to: https://www.npmjs.com/settings/YOUR_USERNAME/tokens

#### Step 2.2: Create New Token

Click "Generate New Token" and select **"Granular Access Token"** (recommended for better security)

#### Step 2.3: Configure Token Settings

Fill in the token configuration:

**Token name:**
```
better-next-app-ci
```
(or any descriptive name like "github-actions-publish")

**Description (optional):**
```
GitHub Actions automation for publishing create-better-next-app
```

**Expiration:**
- Select **"90 days"** or **"Custom"** with a longer duration
- For production, consider 1 year and set a calendar reminder to rotate it
- Avoid "No expiration" for security reasons

**Packages and scopes - Permissions:**

Select **"Read and write"** for:
- ✅ **Packages and scopes** - This allows publishing packages

**Organizations - Permissions:**

Leave as **"No access"** (unless you're publishing under an organization)

**Summary should show:**
- Provide **read and write** access to packages and scopes
- Expires on [your selected date]

#### Step 2.4: Generate and Copy Token

1. Click **"Generate token"**
2. **IMPORTANT:** Copy the token immediately (starts with `npm_...`)
3. Store it securely - you won't be able to see it again!

Example token format: `npm_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

### 3. Add NPM Token to GitHub Secrets

#### Step 3.1: Navigate to Repository Secrets

1. Go to your repository: https://github.com/yeasin2002/better-next-app
2. Click **Settings** (top menu)
3. In the left sidebar, expand **Secrets and variables**
4. Click **Actions**
5. Or go directly to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions

#### Step 3.2: Create New Secret

1. Click **"New repository secret"** (green button)
2. Fill in:
   - **Name:** `NPM_TOKEN` (must be exactly this)
   - **Secret:** Paste your NPM token (the one starting with `npm_...`)
3. Click **"Add secret"**

#### Step 3.3: Verify Secret

You should see `NPM_TOKEN` listed under "Repository secrets" with:
- Updated: Just now
- A green checkmark indicating it's set

### 4. Verify Setup with Task

Run the setup task to verify everything is configured correctly:

```bash
task npm:setup
```

This will check:
- ✓ NPM is installed
- ✓ You're logged in to NPM
- ✓ Package name is available (or you own it)

### 5. Claim Package Name (First Time Only)

Before your first automated release, manually publish to claim the package name:

```bash
task npm:publish
```

This publishes the current version (v0.0.2) to NPM and reserves the package name for your account.

**Expected output:**
```
Updated package.json to version 0.0.2
npm notice Publishing to https://registry.npmjs.org/
+ create-better-next-app@0.0.2
✅ Published to NPM successfully!
```

## Token Types Comparison

### Granular Access Token (Recommended) ✅

**Pros:**
- Fine-grained permissions (only what you need)
- Can scope to specific packages
- Better security audit trail
- Recommended by NPM for automation

**Cons:**
- Slightly more complex setup
- Must configure permissions explicitly

**Use for:** GitHub Actions, CI/CD pipelines

### Classic Token (Legacy)

**Types:**
- **Automation:** For CI/CD (no 2FA required)
- **Publish:** For publishing packages (requires 2FA)
- **Read-only:** For downloading private packages

**Note:** Classic tokens are being phased out. Use Granular tokens for new projects.

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
# Ensure you're logged in
npm login

# Publish using Task (recommended)
task npm:publish

# Or manually
cd npm
npm version 0.1.0 --no-git-tag-version
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
# Preview what will be published
task npm:pack

# Test installation and execution
task npm:test

# Check version synchronization
task npm:version
```

## Troubleshooting

### "Package name already taken"

If someone else claimed the package name:

```bash
# Check who owns it
task npm:setup
```

If you don't own it:
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

### Version mismatch

If NPM version doesn't match git tag:

```bash
# Sync version automatically
task npm:version
```

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

## Quick Reference

### Available Tasks

```bash
# Setup and verification
task npm:setup          # Check NPM login and package availability
task npm:test           # Test package locally
task npm:pack           # Preview package contents
task npm:version        # Sync version with git tag

# Publishing
task npm:publish        # Publish to NPM manually
```

### First-Time Setup Checklist

- [ ] Install Node.js and npm
- [ ] Run `npm login`
- [ ] Run `task npm:setup` to verify
- [ ] Generate NPM token (Automation type)
- [ ] Add `NPM_TOKEN` to GitHub Secrets
- [ ] Run `task npm:publish` to claim package name
- [ ] Create first release: `git tag -a v0.1.0 -m "Release v0.1.0" && git push origin v0.1.0`

### Automated Release Workflow

1. Push a git tag: `git push origin v0.1.0`
2. GitHub Actions automatically:
   - Builds binaries with GoReleaser
   - Creates GitHub release
   - Publishes to NPM registry
3. Users can install: `npx create-better-next-app@latest`
