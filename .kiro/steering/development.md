---
inclusion: always
---

# Development Workflow

## Task Runner

This project uses [Task](https://taskfile.dev) instead of Make for cross-platform compatibility.

### Quick Start

```bash
# Show all available tasks
task --list

# Setup development environment
task setup

# Run in development mode
task dev

# Build for current platform
task build

# Run all checks (format, lint, test)
task check
```

## Common Tasks

### Development

- `task dev` - Run the application

### Building

- `task build` - Build for current platform
- `task build:all` - Build for Linux, macOS, Windows

### Testing

- `task test` - Run all tests

### Code Quality

- `task fmt` - Format code
- `task lint` - Run golangci-lint
- `task vet` - Run go vet
- `task check` - Run all checks (fmt, vet, lint, test)

### Dependencies

- `task deps` - Download dependencies

### Cleaning

- `task clean` - Clean build artifacts

### Setup

- `task setup` - Setup development environment (installs tools and git hooks)

## Code Quality Standards

### Linting

The project uses golangci-lint with a minimal configuration:

- **Enabled linters**: errcheck, gosimple, govet, ineffassign, staticcheck, unused, gofmt, goimports
- **Cyclomatic complexity limit**: 15

### Formatting

- Use `gofmt` for formatting
- Use `goimports` for import organization
- Run `task fmt` before committing (or let git hooks do it)

### Testing

- Write tests for all new features
- Use table-driven tests for multiple scenarios
- Run `task test` to verify

## Pre-commit Checklist

Git hooks automatically run before each commit:

1. Format code (`task fmt`)
2. Run go vet (`task vet`)
3. Run golangci-lint (`task lint`)
4. Run all tests (`task test`)

You can also run manually:

```bash
task check
```

## Git Hooks

This project uses Lefthook for automatic code quality checks.

### Pre-commit (runs automatically)
- Formats code
- Runs vet
- Runs linter
- Runs tests

### Pre-push (runs automatically)
- Runs full check suite

### Skip hooks if needed
```bash
git commit --no-verify
git push --no-verify
```

## Troubleshooting

### Task not found

Install Task:

```bash
# macOS
brew install go-task/tap/go-task

# Linux/macOS (script)
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

# Windows (Scoop)
scoop install task

# NPM (cross-platform)
npm install -g @go-task/cli
```

### golangci-lint not found

Install golangci-lint:

```bash
# macOS
brew install golangci-lint

# Go install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Or run setup task
task setup
```

### Lefthook not running

```bash
# Reinstall hooks
lefthook install

# Or skip temporarily
git commit --no-verify
```

### Build fails

1. Ensure Go 1.21+ is installed
2. Run `task deps` to download dependencies
3. Run `task clean` to clear caches
4. Try building again

## IDE Integration

### VS Code

Install the Go extension and configure:

```json
{
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "go.formatTool": "goimports",
  "editor.formatOnSave": true
}
```

### GoLand/IntelliJ

1. Go to Settings → Tools → File Watchers
2. Add golangci-lint file watcher
3. Enable "Format on save"
