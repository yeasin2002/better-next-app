# Technology Stack

## Language & Runtime

- Go 1.21 or higher
- No external runtime dependencies for the compiled binary

## Core Dependencies

- **Cobra** - CLI framework for argument parsing and command structure
- **Huh** - Interactive terminal prompts (replaces `prompts` from TypeScript)
- **Lipgloss** - Terminal styling and colors (replaces `picocolors`)
- **Doublestar** - File globbing patterns (replaces `fast-glob`)
- **Go-git** (optional) - Git operations

## Dev Dependencies
- **task** - A fast, cross-platform build tool inspired by Make, designed for modern workflows. (github: https://github.com/go-task/task)
- **goreleaser** - GoReleaser is a tool that automatically builds, packages, and publishes your Go application releases so you do not have to handle the process manually
- **Github CI/CD** - for release automation and workflow management and creating PR and issue templates 

## Build System

### Development Commands

```bash
# Install dependencies
go mod download

# Run locally
go run main.go my-test-app

# Build for current platform
go build -o better-next-app

# Build with optimizations (smaller binary)
go build -ldflags="-s -w" -o better-next-app
```

### Cross-Platform Compilation

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o better-next-app-linux

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o better-next-app-macos

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o better-next-app-macos-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o better-next-app.exe
```

## Template System

Templates are embedded at compile time using Go's `embed` directive, allowing single-binary distribution with no external file dependencies.

## Testing

No test framework is currently configured. When adding tests, consider:
- Standard library `testing` package
- Table-driven tests (Go idiom)
- Test files with `_test.go` suffix
