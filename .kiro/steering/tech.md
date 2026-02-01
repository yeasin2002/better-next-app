# Technology Stack

## Language & Runtime

- Go 1.21 or higher
- No external runtime dependencies for the compiled binary

## Core Dependencies

- **Cobra** - CLI framework for argument parsing and command structure
- **Viper** - Configuration management and user preference persistence
- **Huh** - Interactive terminal prompts (replaces `prompts` from TypeScript)
- **Lipgloss** - Terminal styling and colors (replaces `picocolors`)
- **Doublestar** - File globbing patterns (replaces `fast-glob`)

## Development Tools

- **Task** - Cross-platform task runner (replaces Make) - https://github.com/go-task/task
- **golangci-lint** - Fast, parallel Go linter
- **goimports** - Import formatter
- **Lefthook** - Git hooks manager
- **GoReleaser** - Automated release tool for Go applications
- **GitHub Actions** - CI/CD for release automation and workflow management

## Build System

### Development Commands

```bash
# Install dependencies
task deps

# Run locally
task dev

# Build for current platform
task build

# Build for all platforms
task build:all
```

### Cross-Platform Compilation

```bash
# All platforms at once
task build:all

# Or individually with GOOS/GOARCH
GOOS=linux GOARCH=amd64 go build -o better-next-app-linux .
GOOS=darwin GOARCH=amd64 go build -o better-next-app-darwin .
GOOS=darwin GOARCH=arm64 go build -o better-next-app-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -o better-next-app.exe .
```

## Template System

Templates are embedded at compile time using Go's `embed` directive, allowing single-binary distribution with no external file dependencies.

## Code Quality

### Linting

Minimal golangci-lint configuration with essential linters:
- errcheck, gosimple, govet, ineffassign, staticcheck, unused
- gofmt, goimports

### Git Hooks

Lefthook manages pre-commit and pre-push hooks:
- Pre-commit: format, vet, lint, test (runs in parallel)
- Pre-push: full check suite

## Testing

Use standard library `testing` package:
- Table-driven tests (Go idiom)
- Test files with `_test.go` suffix
- Run with `task test`
