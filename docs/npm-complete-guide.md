# NPM Publishing - Complete Guide

A comprehensive guide for publishing `create-better-next-app` to NPM registry, including setup, troubleshooting, and release management.

## Table of Contents

1. [Overview](#overview)
2. [Quick Start (5 Minutes)](#quick-start-5-minutes)
3. [Detailed Setup](#detailed-setup)
4. [Token Generation Guide](#token-generation-guide)
5. [Release Process](#release-process)
6. [Troubleshooting](#troubleshooting)
7. [Maintenance](#maintenance)

---

## Overview

### What This Does

The NPM package is a lightweight wrapper that:
1. Downloads the appropriate binary from GitHub releases during `npm install`
2. Provides a Node.js shim to execute the binary
3. Works with `npx`, `pnpm dlx`, `bunx`, etc.

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

### How It Works

**Installation Flow:**
1. User runs: `npx create-better-next-app@latest`
2. NPM downloads the package
3. `postinstall` script detects OS and architecture
4. Downloads matching binary from GitHub releases
5. Extracts and makes binary executable
6. User can create Next.js projects!

**Execution Flow:**
1. Node.js executes `bin/create-better-next-app.js`
2. Script spawns the Go binary with user's arguments
3. Binary runs and creates the Next.js project

---

## Quick Start (5 Minutes)

### Prerequisites

- ✅ Node.js and npm installed
- ✅ NPM account created at https://www.npmjs.com/signup
- ✅ Repository pushed to GitHub

### Step 1: Login to NPM

```bash
npm login
```

**Note:** This creates a 2-hour session token for local use only. For CI/CD, you need a Granular Access Token (next step).

### Step 2: Verify Setup

```bash
task npm:setup
```

This checks:
- ✓ NPM is installed
- ✓ You're logged in
- ✓ Package name is available (or you own it)

### Step 3: Generate NPM Token

**⚠️ CRITICAL:** You must enable "Bypass 2FA for automation"

1. Go to: https://www.npmjs.com/settings/YOUR_USERNAME/tokens
2. Click "Generate New Token" → "Granular Access Token"
3. Configure:
   - Token name: `better-next-app-ci`
   - Expiration: `90 days` (maximum for write tokens)
   - Packages permission: `Read and write` ✅
   - **Bypass 2FA:** `Enabled` ✅ (CRITICAL!)
   - Organizations permission: `No access`
4. Copy token immediately (starts with `npm_...`)

### Step 4: Add Token to GitHub

1. Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Click "New repository secret"
3. Name: `NPM_TOKEN` (exactly this, case-sensitive)
4. Secret: Paste your NPM token
5. Click "Add secret"

### Step 5: Create Release

```bash
# Commit any changes
git add .
git commit -m "chore: prepare for release"
git push origin main

# Create and push tag
git tag -a v0.0.3 -m "Release v0.0.3"
git push origin v0.0.3
```

GitHub Actions will automatically:
1. Build binaries with GoReleaser
2. Create GitHub release
3. Publish to NPM

### Step 6: Verify

**Check GitHub Actions:** https://github.com/yeasin2002/better-next-app/actions

**Check NPM Package:** https://www.npmjs.com/package/create-better-next-app

**Test Installation:**
```bash
npx create-better-next-app@latest --help
```

---

## Detailed Setup

### 1. Create NPM Account

If you don't have an NPM account:
1. Go to https://www.npmjs.com/signup
2. Create an account
3. Verify your email
4. Enable 2FA (recommended)

### 2. NPM Token Requirements (December 2025 Update)

**⚠️ Important Changes:**
- ❌ Classic tokens permanently revoked (December 9, 2025)
- ✅ Granular Access Tokens are now required
- ⏰ Session tokens (from `npm login`) expire after 2 hours
- 🔒 Must enable "Bypass 2FA" for CI/CD automation
- ⏳ Write tokens limited to 90 days maximum

[Read the official announcement](https://github.blog/changelog/2025-12-09-npm-classic-tokens-revoked-session-based-auth-and-cli-token-management-now-available/)

### 3. Token Types Comparison

| Feature | Granular Token | Session Token |
|---------|---------------|---------------|
| **Use Case** | CI/CD automation | Local development |
| **Expiration** | 90 days (max) | 2 hours |
| **2FA Required** | Can bypass | Yes |
| **Suitable for GitHub Actions** | ✅ Yes | ❌ No |
| **Visibility** | In token list | Hidden |

**For this project:** Use Granular Access Token with "Bypass 2FA" enabled.

---

## Token Generation Guide

### Visual Step-by-Step

#### Step 1: Access Token Page

**URL:** https://www.npmjs.com/settings/YOUR_USERNAME/tokens

**Path:** NPM Website → Profile Picture → Access Tokens

#### Step 2: Create New Token

Click "Generate New Token" → Select "Granular Access Token"

**Note:** Classic tokens are no longer available (removed December 9, 2025)

#### Step 3: Configure Token

**General Section:**

```
Token name: better-next-app-ci
Description: GitHub Actions automation for publishing
Expiration: 90 days (maximum for write tokens)
```

**Packages and Scopes Section:**

```
┌─────────────────────────────────────────┐
│  Permissions:                            │
│  ● Read and write  ← Select this! ✅    │
│                                          │
│  ☑ Bypass 2FA for automation  ← MUST!  │
└─────────────────────────────────────────┘
```

**Organizations Section:**

```
Permissions: ○ No access
```

#### Step 4: Generate and Copy

1. Click "Generate token"
2. **IMMEDIATELY COPY** the token (starts with `npm_...`)
3. Store securely (you won't see it again!)

#### Step 5: Add to GitHub Secrets

1. Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Click "New repository secret"
3. Name: `NPM_TOKEN`
4. Value: Paste token
5. Click "Add secret"

### Verification Checklist

Before proceeding, verify:

- [ ] Token type: Granular Access Token
- [ ] Token name: Descriptive (e.g., `better-next-app-ci`)
- [ ] Expiration: 90 days
- [ ] Packages permission: Read and write ✅
- [ ] **Bypass 2FA: Enabled** ✅ (CRITICAL!)
- [ ] Organizations permission: No access
- [ ] Token copied and saved securely
- [ ] GitHub secret created with name `NPM_TOKEN`
- [ ] Calendar reminder set for token rotation (85 days)

---

## Release Process

### Automated Release (Recommended)

This is the safest and most automated approach.

#### Step 1: Prepare Changes

```bash
# Make your changes
git add .
git commit -m "feat: your changes"
git push origin main
```

#### Step 2: Create Release Tag

```bash
# Create tag
git tag -a v0.0.4 -m "Release v0.0.4"

# Push tag
git push origin v0.0.4
```

#### Step 3: Monitor Progress

**GitHub Actions:** https://github.com/yeasin2002/better-next-app/actions

**Expected Timeline:**
- 2-3 minutes: GoReleaser builds binaries
- 1 minute: Creates GitHub release
- 30 seconds: Publishes to NPM
- **Total: ~4 minutes**

#### Step 4: Verify Success

**GitHub Release:**
https://github.com/yeasin2002/better-next-app/releases

Should show:
- ✅ Release with version tag
- ✅ Binaries for all platforms
- ✅ Changelog

**NPM Package:**
https://www.npmjs.com/package/create-better-next-app

Should show:
- ✅ New version number
- ✅ "Published X minutes ago"
- ✅ Your username as maintainer

**Test Installation:**
```bash
npx create-better-next-app@latest --help
```

### Manual Release (Fallback)

Only use if GitHub Actions fails.

#### Option 1: Use Token Locally

```powershell
# Set token in environment (PowerShell)
$env:NPM_TOKEN="npm_YOUR_TOKEN_HERE"

# Configure NPM
npm config set //registry.npmjs.org/:_authToken $env:NPM_TOKEN

# Publish
task npm:publish

# Clean up
npm config delete //registry.npmjs.org/:_authToken
```

#### Option 2: Use Session Token

```bash
# Login (creates 2-hour session)
npm login

# Publish (will require 2FA)
task npm:publish
```

**Note:** This requires 2FA authentication during publish.

### Version Synchronization

**Important:** NPM version must match GitHub release version.

- Git tag: `v0.1.0`
- NPM version: `0.1.0` (without `v` prefix)

The GitHub Actions workflow automatically handles this.

---

## Troubleshooting

### Error: 403 Two-factor authentication required

**Error Message:**
```
npm error 403 Two-factor authentication or granular access token with bypass 2fa enabled is required
```

**Cause:** Token doesn't have "Bypass 2FA" enabled

**Solution:**
1. Generate new token with "Bypass 2FA" enabled
2. Update GitHub secret `NPM_TOKEN`
3. Try publishing again

**Visual Guide:**
```
┌─────────────────────────────────────────┐
│  Packages and scopes                     │
│  ● Read and write                        │
│  ☑ Bypass 2FA for automation  ← CHECK!  │
└─────────────────────────────────────────┘
```

### Error: sed command not found (Windows)

**Error Message:**
```
"sed": executable file not found in $PATH
```

**Cause:** Windows doesn't have `sed` by default

**Solution:** Already fixed in Taskfile (uses PowerShell on Windows)

### Error: Token expired

**Error Message:**
```
npm error 401 Unauthorized
```

**Cause:** Token expired (90 days maximum)

**Solution:**
1. Generate new token
2. Update GitHub secret
3. Set reminder for 85 days

### Error: Package already exists

**Error Message:**
```
npm error 403 You cannot publish over previously published versions
```

**Cause:** Version already published

**Solution:**
```bash
# Create new tag with incremented version
git tag -a v0.0.4 -m "Release v0.0.4"
git push origin v0.0.4
```

### Error: Binary not found during install

**Cause:** GitHub release doesn't exist or binary naming mismatch

**Solution:**
1. Verify GitHub release exists
2. Check binary naming in `.goreleaser.yaml`
3. Test download URL manually

### Error: Not logged in

**Error Message:**
```
npm error code ENEEDAUTH
```

**Cause:** Not logged in locally

**Solution:**
```bash
npm login
```

### GitHub Actions Workflow Not Starting

**Cause:** Workflow configuration issue

**Solution:**
1. Check `.github/workflows/release.yml` exists
2. Verify it triggers on `tags: ['v*']`
3. Check GitHub Actions is enabled

### NPM Publish Fails in GitHub Actions

**Causes:**
- `NPM_TOKEN` secret not set
- Token doesn't have "Bypass 2FA"
- Token expired
- Package name taken

**Solution:**
1. Verify secret at: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Check token has "Bypass 2FA" enabled
3. Verify token hasn't expired
4. Check package ownership

---

## Maintenance

### Token Rotation Schedule

Since write tokens expire after 90 days:

**Day 1:** Generate token, add to GitHub
**Day 75:** Set reminder notification
**Day 85:** Generate new token, update GitHub secret
**Day 90:** Old token expires (no disruption if rotated)

**Pro tip:** Use calendar app to track expiration dates.

### Updating Package Metadata

To update description, keywords, etc.:

```bash
# Edit npm/package.json
# Commit changes
git add npm/package.json
git commit -m "chore: update package metadata"
git push origin main

# Create new release
git tag -a v0.0.5 -m "Release v0.0.5"
git push origin v0.0.5
```

### Deprecating Old Versions

```bash
npm deprecate create-better-next-app@0.0.1 "Please upgrade to 0.1.0"
```

### Unpublishing (Use Carefully!)

You can only unpublish within 72 hours:

```bash
npm unpublish create-better-next-app@0.0.1
```

### Monitoring

After publishing, monitor:

1. **NPM downloads:** https://www.npmjs.com/package/create-better-next-app
2. **GitHub releases:** https://github.com/yeasin2002/better-next-app/releases
3. **Issues:** User-reported installation problems
4. **GitHub Actions:** Workflow success/failure

---

## Testing

### Local Testing

Before releasing:

```bash
# Preview package contents
task npm:pack

# Test installation locally
task npm:test

# Check version synchronization
task npm:version
```

### Integration Testing

After publishing:

```bash
# Test with npx
npx create-better-next-app@latest my-test-app

# Test with different package managers
pnpm dlx create-better-next-app@latest my-test-app
bunx create-better-next-app@latest my-test-app
```

---

## Security Best Practices

1. **Never commit tokens** to repository
2. **Use 90-day expiration** for write tokens
3. **Set rotation reminders** before expiration
4. **Enable "Bypass 2FA"** for CI/CD only
5. **Use minimal permissions** (only packages)
6. **Store securely** in password manager
7. **Monitor usage** via NPM audit logs
8. **Revoke if compromised** immediately

---

## Quick Reference

### Available Tasks

```bash
task npm:setup      # Verify NPM setup
task npm:test       # Test package locally
task npm:pack       # Preview package contents
task npm:version    # Sync version with git tag
task npm:publish    # Publish manually (fallback)
```

### Common Commands

```bash
# Check NPM login
npm whoami

# Login to NPM
npm login

# Create release
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

# Monitor GitHub Actions
# https://github.com/yeasin2002/better-next-app/actions

# Check NPM package
# https://www.npmjs.com/package/create-better-next-app
```

### Important URLs

- **NPM Tokens:** https://www.npmjs.com/settings/YOUR_USERNAME/tokens
- **GitHub Secrets:** https://github.com/yeasin2002/better-next-app/settings/secrets/actions
- **GitHub Actions:** https://github.com/yeasin2002/better-next-app/actions
- **GitHub Releases:** https://github.com/yeasin2002/better-next-app/releases
- **NPM Package:** https://www.npmjs.com/package/create-better-next-app

---

## Future Enhancements

Consider adding:

1. **OIDC Trusted Publishing** - No tokens to manage
2. **Scoped package** - `@yeasin2002/create-better-next-app`
3. **NPM Provenance** - Supply chain security
4. **Binary caching** - Faster reinstalls
5. **Fallback mirrors** - Alternative download sources
6. **Version checking** - Warn about outdated versions

---

## Resources

- **NPM Token Documentation:** https://docs.npmjs.com/about-access-tokens
- **NPM CLI Token Management:** https://docs.npmjs.com/cli/
- **NPM Security Update:** https://github.blog/changelog/2025-12-09-npm-classic-tokens-revoked-session-based-auth-and-cli-token-management-now-available/
- **GitHub Secrets Documentation:** https://docs.github.com/en/actions/security-guides/encrypted-secrets
- **OIDC Trusted Publishing:** https://docs.npmjs.com/trusted-publishers
- **GoReleaser Documentation:** https://goreleaser.com/

---

## Support

If you encounter issues:

1. Check this guide's troubleshooting section
2. Review GitHub Actions logs
3. Verify token settings on NPM
4. Check GitHub secret configuration
5. Open an issue on GitHub

---

**Last Updated:** February 4, 2026
**NPM Security Update:** December 9, 2025
