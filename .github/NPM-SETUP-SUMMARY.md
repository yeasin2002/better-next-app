# NPM Publishing Setup - Summary

This document summarizes the NPM publishing setup for `create-better-next-app`.

## What Was Created

### NPM Package Structure (`npm/` directory)
- ✅ `package.json` - Package metadata (name: `create-better-next-app`)
- ✅ `bin/create-better-next-app.js` - Node.js wrapper that executes the Go binary
- ✅ `scripts/install.js` - Post-install script that downloads binaries from GitHub releases
- ✅ `README.md` - NPM-specific documentation
- ✅ `.npmignore` - Controls what gets published

### GitHub Actions Integration
- ✅ Updated `.github/workflows/release.yml` - Automatically publishes to NPM after GoReleaser

### Cross-Platform Task Commands
- ✅ `task npm:setup` - Verify NPM login and package availability
- ✅ `task npm:test` - Test package locally
- ✅ `task npm:pack` - Preview package contents
- ✅ `task npm:version` - Sync version with git tag
- ✅ `task npm:publish` - Publish to NPM manually

### Documentation
- ✅ `docs/npm-publishing.md` - Complete setup and troubleshooting guide
- ✅ `docs/npm-quick-start.md` - 5-minute quick start guide
- ✅ Updated `README.md` - Added NPM installation and publishing sections
- ✅ Updated `.kiro/steering/development.md` - Added NPM task documentation

## How It Works

### User Installation Flow

When users run `npx create-better-next-app@latest`:

1. NPM downloads the package from registry
2. Post-install script (`scripts/install.js`) runs automatically
3. Script detects user's OS and architecture
4. Downloads matching binary from GitHub releases
5. Extracts and makes binary executable
6. User can now create Next.js projects!

### Automated Release Flow

When you push a git tag:

1. GitHub Actions triggers on tag push
2. GoReleaser builds binaries for all platforms
3. Creates GitHub release with binaries
4. NPM publish job runs after GoReleaser
5. Updates package version to match tag
6. Publishes to NPM registry
7. Users can install with `npx create-better-next-app@latest`

## Setup Checklist

- [ ] Install Node.js and npm
- [ ] Run `npm login`
- [ ] Run `task npm:setup` to verify
- [ ] Generate NPM token (Automation type) at https://www.npmjs.com/settings/YOUR_USERNAME/tokens
- [ ] Add `NPM_TOKEN` to GitHub Secrets at https://github.com/yeasin2002/better-next-app/settings/secrets/actions
- [ ] Run `task npm:publish` to claim package name
- [ ] Create first release: `git tag -a v0.1.0 -m "Release v0.1.0" && git push origin v0.1.0`

## Key Features

### Cross-Platform Support
- Works on Linux, macOS (Intel/ARM), and Windows
- Automatic OS/architecture detection
- No platform-specific scripts needed

### Zero Dependencies
- Users don't need Go installed
- Binary is downloaded during npm install
- Single command to get started

### Automated Versioning
- NPM version automatically syncs with git tags
- No manual version updates needed
- Consistent versioning across platforms

### Developer-Friendly
- All commands available via Task
- Clear error messages
- Comprehensive documentation

## Testing

Before releasing to production:

```bash
# Preview what will be published
task npm:pack

# Test installation locally
task npm:test

# Check version synchronization
task npm:version

# Test full release flow (no git tag needed)
task release:snapshot
```

## Troubleshooting

### Common Issues

1. **"Package name already taken"**
   - Check ownership with `task npm:setup`
   - Use scoped package: `@yeasin2002/create-better-next-app`

2. **"Binary not found" during install**
   - Ensure GitHub release exists with binaries
   - Check GoReleaser built successfully
   - Verify binary naming matches install script

3. **"Not logged in to NPM"**
   - Run `npm login`
   - Verify with `npm whoami`

4. **Version mismatch**
   - Run `task npm:version` to sync

## Resources

- Quick Start: [docs/npm-quick-start.md](../docs/npm-quick-start.md)
- Full Guide: [docs/npm-publishing.md](../docs/npm-publishing.md)
- NPM Package: https://www.npmjs.com/package/create-better-next-app
- GitHub Releases: https://github.com/yeasin2002/better-next-app/releases

## Next Steps

1. Complete the setup checklist above
2. Test locally with `task npm:test`
3. Publish first version with `task npm:publish`
4. Create first release tag
5. Verify automated publishing works
6. Share with users: `npx create-better-next-app@latest`

---

**Note:** This setup uses a "binary wrapper" approach where the NPM package downloads the Go binary from GitHub releases. This is a common pattern for distributing native binaries via NPM (used by tools like esbuild, swc, etc.).
