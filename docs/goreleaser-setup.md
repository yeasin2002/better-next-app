# GoReleaser Setup Summary

GoReleaser has been successfully integrated into the Better Next App project.

## What Was Added

### Configuration Files

1. **`.goreleaser.yaml`** - Main GoReleaser configuration
   - Builds for Linux, macOS, Windows (amd64 and arm64)
   - Creates compressed archives (tar.gz for Unix, zip for Windows)
   - Generates checksums and changelogs
   - Configured for GitHub releases

2. **`.github/workflows/release.yml`** - GitHub Actions workflow
   - Automatically triggers on git tag push
   - Builds and publishes releases to GitHub

3. **`docs/releasing.md`** - Comprehensive release documentation
   - Step-by-step release process
   - Versioning guidelines
   - Troubleshooting tips

### Updated Files

1. **`Taskfile.yml`** - Added release tasks:
   - `task release` - Create a production release
   - `task release:snapshot` - Test build locally
   - `task release:check` - Validate configuration
   - `task setup` - Now installs GoReleaser

2. **`.gitignore`** - Added GoReleaser artifacts

3. **`README.md`** - Updated with release information

4. **`CHANGELOG.md`** - Created for tracking releases

## Quick Start

### Test Locally

```bash
# Validate configuration
task release:check

# Build snapshot (no git tag required)
task release:snapshot

# Check the dist/ directory
ls dist/
```

### Create a Release

```bash
# Commit all changes
git add .
git commit -m "feat: ready for first release"
git push

# Create and push a tag
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

# GitHub Actions will automatically:
# - Build binaries for all platforms
# - Create checksums
# - Generate changelog
# - Publish GitHub release
```

## Build Artifacts

Each release includes:

- **Linux** (amd64, arm64)
- **macOS** (amd64, arm64)
- **Windows** (amd64)
- **Checksums** (SHA256)
- **Auto-generated changelog**

## Features

### Current

- ✅ Cross-platform builds
- ✅ Automated GitHub releases
- ✅ Changelog generation from commits
- ✅ Checksum generation
- ✅ Archive creation (tar.gz/zip)
- ✅ GitHub Actions integration

### Future (Commented Out)

- 🔜 Homebrew tap (for macOS)
- 🔜 Scoop bucket (for Windows)
- 🔜 Docker images
- 🔜 NPM package wrapper

## Versioning

Follow [Semantic Versioning](https://semver.org/):

- **v1.0.0** - Major (breaking changes)
- **v0.1.0** - Minor (new features)
- **v0.0.1** - Patch (bug fixes)

## Commit Message Format

Use [Conventional Commits](https://www.conventionalcommits.org/) for automatic changelog:

- `feat:` - New features
- `fix:` - Bug fixes
- `perf:` - Performance improvements
- `refactor:` - Code refactoring
- `docs:` - Documentation (excluded from changelog)
- `test:` - Tests (excluded from changelog)
- `chore:` - Maintenance (excluded from changelog)

## Verification

Configuration is valid and tested:

```bash
$ task release:check
✓ Configuration is valid

$ task release:snapshot
✓ Built successfully for all platforms
```

## Next Steps

1. Complete the core CLI implementation
2. Write comprehensive tests
3. Create your first release: `v0.1.0`
4. Consider enabling Homebrew/Scoop when stable

## Resources

- [GoReleaser Documentation](https://goreleaser.com)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Semantic Versioning](https://semver.org/)
- [Conventional Commits](https://www.conventionalcommits.org/)
