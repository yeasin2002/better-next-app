# NPM Publishing - Quick Start

This is a condensed guide for setting up NPM publishing. For detailed information, see [npm-publishing.md](./npm-publishing.md).

## ⚠️ Important Update (December 9, 2025)

NPM has permanently revoked all classic tokens. You **must** use Granular Access Tokens.

**Key Changes:**
- ❌ Classic tokens no longer work
- ✅ Granular Access Tokens are now required
- ⏰ Session tokens (from `npm login`) expire after 2 hours
- 🔒 Must enable "Bypass 2FA" for CI/CD automation
- ⏳ Write tokens limited to 90 days maximum

## Prerequisites

- ✅ Node.js and npm installed
- ✅ NPM account created at https://www.npmjs.com/signup
- ✅ Repository pushed to GitHub

## Setup (5 Minutes)

### Step 1: Login to NPM

```bash
npm login
```

Enter your NPM username, password, and email when prompted.

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

**Need help?** See the detailed visual guide: [npm-token-guide.md](./npm-token-guide.md)

#### Quick Steps:

1. **Go to:** https://www.npmjs.com/settings/YOUR_USERNAME/tokens
2. **Click:** "Generate New Token"
3. **Select:** "Granular Access Token" (only option available)
4. **Configure:**
   - Token name: `better-next-app-ci`
   - Expiration: `90 days` (maximum for write tokens)
   - Packages permission: `Read and write` ✅
   - **Bypass 2FA:** `Enabled` ✅ (critical for CI/CD!)
   - Organizations permission: `No access`
5. **Click:** "Generate token"
6. **Copy:** The token immediately (starts with `npm_...`)

**Important:** You MUST enable "Bypass 2FA for automation" or GitHub Actions will fail!

### Step 4: Add Token to GitHub

#### 4.1 Navigate to Secrets

Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions

#### 4.2 Create Secret

1. Click **"New repository secret"**
2. **Name:** `NPM_TOKEN` (exactly this, case-sensitive)
3. **Secret:** Paste your NPM token
4. Click **"Add secret"**

### Step 5: Claim Package Name

```bash
task npm:publish
```

This publishes v0.0.2 to NPM and reserves the package name.

**Expected output:**
```
Updated package.json to version 0.0.2
+ create-better-next-app@0.0.2
✅ Published to NPM successfully!
```

## Done! 🎉

From now on, just push a git tag and GitHub Actions will automatically publish to NPM:

```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

## What Happens Next?

When you push a tag:
1. ⚙️ GitHub Actions triggers
2. 🔨 GoReleaser builds binaries
3. 📦 Creates GitHub release
4. 🚀 Publishes to NPM automatically
5. ✅ Users can install: `npx create-better-next-app@latest`

## Testing Before Release

```bash
# Preview package contents
task npm:pack

# Test installation locally
task npm:test

# Check version sync
task npm:version
```

## Common Commands

```bash
task npm:setup      # Verify NPM setup
task npm:test       # Test package locally
task npm:pack       # Preview what will be published
task npm:version    # Sync version with git tag
task npm:publish    # Publish manually
```

## Troubleshooting

### "Not logged in to NPM"

```bash
npm login
```

### "Package name already taken"

If you don't own it, update `npm/package.json`:

```json
{
  "name": "@yeasin2002/create-better-next-app"
}
```

### "Version mismatch"

```bash
task npm:version
```

### "Binary not found" during install

Make sure you've created a GitHub release with binaries first:

```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

## What Happens When Users Install

1. User runs: `npx create-better-next-app@latest`
2. NPM downloads your package
3. Post-install script detects their OS/architecture
4. Downloads matching binary from GitHub releases
5. User can create Next.js projects!

## Resources

- Full documentation: [npm-publishing.md](./npm-publishing.md)
- NPM package: https://www.npmjs.com/package/create-better-next-app
- GitHub releases: https://github.com/yeasin2002/better-next-app/releases
