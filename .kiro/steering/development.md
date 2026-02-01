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

# Run with hot reload
task dev:watch

# Build for current platform
task build

# Run all checks (format, lint, test)
task check
```

## Common Tasks

### Development

- `task dev` - Run the application
- `task dev:watch` - Run with hot reload (requires air)

### Building

- `task build` - Build for current platform
- `task build:all` - Build for Linux, macOS, Windows
- `task build:linux` - Build for Linux only
- `task build:darwin` - Build for macOS (Intel)
- `task build:darwin-arm64` - Build for macOS (Apple Silicon)
- `task build:windows` - Build for Windows

### Testing

- `task test` - Run all tests
- `task test:unit` - Run unit tests only
- `task test:coverage` - Generate coverage report
- `task test:bench` - Run benchmarks

### Code Quality

- `task lint` - Run golangci-lint
- `task lint:fix` - Auto-fix linting issues
- `task fmt` - Format code
- `task fmt:check` - Check if code is formatted
- `task vet` - Run go vet
- `task check` - Run all checks (fmt, vet, lint, test)

### Dependencies

- `task deps` - Download dependencies
- `task deps:tidy` - Tidy dependencies
- `task deps:update` - Update all dependencies

### Cleaning

- `task clean` - Clean all artifacts and caches
- `task clean:build` - Clean only build artifacts

### CI/CD

- `task ci` - Run CI pipeline
- `task ci:full` - Run full CI with cross-platform builds

## Code Quality Standards

### Linting

The project uses golangci-lint with a balanced configuration:

- **Enabled linters**: errcheck, gosimple, govet, ineffassign, staticcheck, unused, gofmt, goimports, misspell, revive, gosec, gocritic, gocyclo, dupl, and more
- **Cyclomatic complexity limit**: 15
- **Security checks**: Enabled with reasonable exceptions for CLI tools

### Formatting

- Use `gofmt` for formatting
- Use `goimports` for import organization
- Run `task fmt` before committing

### Testing

- Write tests for all new features
- Maintain test coverage above 70%
- Use table-driven tests for multiple scenarios
- Run `task test:coverage` to check coverage

## Pre-commit Checklist

Before committing code, run:

```bash
task check
```

This will:

1. Check code formatting
2. Run go vet
3. Run golangci-lint
4. Run all tests

## Hot Reload Development

Install air for hot reload:

```bash
go install github.com/air-verse/air@latest
```

Then run:

```bash
task dev:watch
```

## Building for Release

```bash
# Prepare release (runs all checks and builds for all platforms)
task release:prepare
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
task setup:tools
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
