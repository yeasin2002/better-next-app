---
inclusion: always
---

# Git Hooks

This project uses [Lefthook](https://github.com/evilmartians/lefthook) for managing git hooks.

## Setup

Hooks are automatically installed when you run:

```bash
task setup
```

Or manually:

```bash
lefthook install
```

## Pre-commit Hooks

Runs automatically before each commit:

1. **Format** - Formats Go code (`task fmt`)
2. **Vet** - Runs go vet (`task vet`)
3. **Lint** - Runs golangci-lint (`task lint`)
4. **Test** - Runs all tests (`task test`)

All hooks run in parallel for speed.

## Pre-push Hooks

Runs automatically before pushing:

- **Check** - Runs full check suite (`task check`)

## Skipping Hooks

If you need to skip hooks temporarily:

```bash
# Skip pre-commit hooks
git commit --no-verify

# Skip pre-push hooks
git push --no-verify
```

## Configuration

Edit `lefthook.yml` to customize hooks.

## Troubleshooting

If hooks aren't running:

```bash
# Reinstall hooks
lefthook install

# Check hook status
lefthook run pre-commit
```
