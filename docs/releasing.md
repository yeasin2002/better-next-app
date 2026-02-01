# Release Process

This document describes how to create releases for Better Next App using GoReleaser.

## Prerequisites

Install GoReleaser:

```bash
# Using Task
task setup

# Or manually
go install github.com/goreleaser/goreleaser/v2@latest
```

## Local Testing

Before creating a release, test the build locally:

```bash
# Build a snapshot (no git tag required)
task release:snapshot

# Check configuration
task release:check
```

This creates builds in the `dist/` directory without publishing anything.

## Creating a Release

### 1. Prepare the Release

Ensure all changes are committed and pushed:

```bash
git add .
git commit -m "feat: prepare for release"
git push
```

### 2. Create and Push a Tag

```bash
# Create a new tag (use semantic versioning)
git tag -a v0.1.0 -m "Release v0.1.0"

# Push the tag
git push origin v0.1.0
```

### 3. Automated Release

Once you push the tag, GitHub Actions will automatically:
- Build binaries for Linux, macOS, and Windows
- Create checksums
- Generate a changelog
- Create a GitHub release with all artifacts

### 4. Manual Release (if needed)

If you prefer to release manually:

```bash
# Ensure you're on the tagged commit
git checkout v0.1.0

# Run GoReleaser
task release

# Or directly
goreleaser release --clean
```

## Release Artifacts

Each release includes:

- **Binaries**: Pre-built executables for:
  - Linux (amd64, arm64)
  - macOS (amd64, arm64)
  - Windows (amd64)
- **Archives**: Compressed archives (tar.gz for Unix, zip for Windows)
- **Checksums**: SHA256 checksums for verification
- **Changelog**: Auto-generated from commit messages

## Versioning

Follow [Semantic Versioning](https://semver.org/):

- **v1.0.0**: Major release (breaking changes)
- **v0.1.0**: Minor release (new features)
- **v0.0.1**: Patch release (bug fixes)

## Changelog Format

Use [Conventional Commits](https://www.conventionalcommits.org/) for automatic changelog generation:

- `feat:` - New features
- `fix:` - Bug fixes
- `perf:` - Performance improvements
- `refactor:` - Code refactoring
- `docs:` - Documentation changes (excluded from changelog)
- `test:` - Test changes (excluded from changelog)
- `chore:` - Maintenance tasks (excluded from changelog)

Examples:
```bash
git commit -m "feat: add TypeScript support"
git commit -m "fix: resolve package installation issue"
git commit -m "perf: optimize template copying"
```

## Troubleshooting

### Release fails with "dirty working tree"

Ensure all changes are committed:
```bash
git status
git add .
git commit -m "chore: prepare release"
```

### Tag already exists

Delete and recreate the tag:
```bash
git tag -d v0.1.0
git push origin :refs/tags/v0.1.0
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

### GitHub Actions fails

Check the workflow logs at:
https://github.com/yeasin2002/better-next-app/actions

Common issues:
- Missing GITHUB_TOKEN (should be automatic)
- Go version mismatch
- Test failures

## Future Enhancements

Once the project is stable, consider adding:

- **Homebrew tap**: For easy installation on macOS
- **Scoop bucket**: For easy installation on Windows
- **Docker images**: For containerized usage
- **NPM package**: For `npx create-better-next-app`

These are commented out in `.goreleaser.yaml` and can be enabled when ready.
