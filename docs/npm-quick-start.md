# NPM Publishing - Quick Start

This is a condensed guide for setting up NPM publishing. For detailed information, see [npm-publishing.md](./npm-publishing.md).

## Prerequisites

- Node.js and npm installed
- NPM account created at https://www.npmjs.com/signup
- Repository pushed to GitHub

## Setup (5 Minutes)

### 1. Login to NPM

```bash
npm login
```

### 2. Verify Setup

```bash
task npm:setup
```

This checks:
- ✓ NPM is installed
- ✓ You're logged in
- ✓ Package name is available (or you own it)

### 3. Get NPM Token

1. Go to https://www.npmjs.com/settings/YOUR_USERNAME/tokens
2. Click "Generate New Token" → "Classic Token"
3. Select "Automation" type
4. Copy the token (starts with `npm_...`)

### 4. Add Token to GitHub

1. Go to https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Click "New repository secret"
3. Name: `NPM_TOKEN`
4. Value: Paste your token
5. Click "Add secret"

### 5. Claim Package Name

```bash
task npm:publish
```

This publishes v0.0.2 to NPM and claims the package name.

## Done! 🎉

From now on, just push a git tag and GitHub Actions will automatically publish to NPM:

```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

## Testing

Before releasing, test locally:

```bash
# Preview package contents
task npm:pack

# Test installation
task npm:test
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
