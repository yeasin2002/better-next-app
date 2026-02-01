# Contributing to Create Next App (Go)

Thanks for your interest in contributing! This guide will help you get started.

## Prerequisites

- Go 1.21 or later
- Git
- Node.js (for testing generated projects)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yeasin2002/create-next-app-go_prototype.git
cd create-next-app-go_prototype
```

### 2. Install Dependencies

```bash
make deps
# Or: go mod tidy
```

### 3. Build the CLI

```bash
# Development build
make build

# Optimized build (smaller binary)
make build-prod

# Or directly with Go
go build -o create-next-go-app
go build -ldflags="-s -w" -o create-next-go-app
```

### 4. Run the CLI

```bash
# Using Make
make run

# Or directly
./create-next-go-app --help

# Create a test project
./create-next-go-app test-app --yes --skip-install
```

## Project Structure

```
/
├── main.go                    # Entry point, template embedding
├── cmd/
│   └── root.go                # Cobra setup, dependency wiring (120 lines)
├── internal/
│   ├── core/                  # Core infrastructure
│   │   ├── constants/         # Typed constants (Linter, Bundler, etc.)
│   │   ├── errors/            # Structured error types
│   │   └── interfaces/        # Interface definitions
│   ├── cli/                   # CLI layer
│   │   ├── runner.go          # Main orchestration logic
│   │   └── flags.go           # Flag parsing
│   ├── config/                # Configuration structs
│   │   ├── config.go          # Config struct definition
│   │   ├── defaults.go        # Default values
│   │   └── preferences.go     # Preferences struct (legacy)
│   ├── validation/            # Unified validation
│   │   └── validator.go       # NPM name, directory validation
│   ├── template/              # Template operations
│   │   ├── installer.go       # Installation coordinator
│   │   ├── copier.go          # File copying (no global state)
│   │   ├── transformer.go     # File transformations
│   │   └── pkgjson.go         # package.json generation
│   ├── installation/          # Post-installation
│   │   ├── deps.go            # Dependency installation
│   │   ├── git.go             # Git initialization
│   │   └── typegen.go         # TypeScript type generation
│   ├── preferences/           # User preferences management
│   │   └── manager.go         # Preferences persistence
│   ├── prompting/             # Interactive prompts wrapper
│   │   └── prompter.go        # Prompt interface implementation
│   ├── downloading/           # Example downloader wrapper
│   │   └── downloader.go      # GitHub example downloading
│   ├── prompt/                # Prompt implementations (legacy)
│   ├── example/               # Example download (legacy)
│   ├── validate/              # Validation (legacy, wrapped)
│   └── util/                  # Utilities
│       ├── colors.go          # Terminal styling
│       ├── exec.go            # Command execution
│       └── fs.go              # Filesystem helpers
└── templates/                 # Next.js templates (embedded)
    ├── app/                   # App Router
    ├── app-tw/                # App Router + Tailwind
    ├── app-empty/             # Minimal App Router
    ├── app-tw-empty/          # Minimal App Router + Tailwind
    └── app-api/               # API-only (no React)
```

### Architecture Overview

The codebase follows a **clean architecture** with dependency injection:

- **Core Layer** (`internal/core/`) - Shared types, interfaces, constants, errors
- **CLI Layer** (`internal/cli/`) - Command-line interface logic
- **Business Logic** - Domain-specific implementations (validation, template, installation)
- **Infrastructure** - External dependencies (filesystem, git, npm)

All major components implement interfaces defined in `internal/core/interfaces/`, enabling:
- Easy testing with mocks
- Flexible implementations
- Clear contracts between components

## Development Workflow

### Making Changes

1. Create a feature branch:
   ```bash
   git checkout -b feature/my-feature
   ```

2. Make your changes

3. Build and test:
   ```bash
   go build -o create-next-app
   ./create-next-app test-project --yes --skip-install
   ```

4. Clean up test projects:
   ```bash
   rm -rf test-project
   ```

### Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Keep functions focused and small
- Add comments for exported functions

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Or directly with Go
go test ./...
go test -cover ./...
```

## Available Make Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the binary |
| `make build-prod` | Build optimized binary (smaller) |
| `make run` | Build and run |
| `make run-args ARGS="..."` | Run with arguments |
| `make clean` | Remove build artifacts |
| `make test` | Run tests |
| `make test-coverage` | Run tests with coverage report |
| `make deps` | Install dependencies |
| `make install` | Install to GOPATH/bin |
| `make dev` | Development build and test |
| `make release` | Build for all platforms |
| `make fmt` | Format code |
| `make lint` | Lint code |
| `make help` | Show all commands |

## Key Components

### Understanding the Architecture

The codebase uses **dependency injection** and **interface-based design**:

```go
// All dependencies are injected via constructors
runner := cli.NewRunner(
    validator,
    templateInstaller,
    depInstaller,
    gitInitializer,
    typeGenerator,
    exampleDownloader,
    prompter,
    prefsManager,
)
```

This makes the code:
- **Testable** - Components can be mocked
- **Maintainable** - Clear dependencies
- **Flexible** - Easy to swap implementations

### Adding a New CLI Flag

1. Register the flag in `cmd/root.go` `init()`:
   ```go
   rootCmd.Flags().Bool("my-flag", false, "Description")
   ```

2. Handle it in `internal/cli/flags.go` `Parse()`:
   ```go
   if cmd.Flags().Changed("my-flag") {
       cfg.MyField, _ = cmd.Flags().GetBool("my-flag")
   }
   ```

3. If it's a new constant value, add to `internal/core/constants/constants.go`:
   ```go
   type MyOption string
   const (
       MyOptionA MyOption = "option-a"
       MyOptionB MyOption = "option-b"
   )
   ```

### Adding a New Template

1. Create the template directory:
   ```
   templates/my-template/
   ├── js/
   │   └── ...
   └── ts/
       └── ...
   ```

2. Add constant in `internal/core/constants/constants.go`:
   ```go
   const (
       TemplateMyTemplate TemplateType = "my-template"
   )
   ```

3. Update `SelectTemplate()` in `internal/template/installer.go`:
   ```go
   if cfg.MyCondition {
       return string(constants.TemplateMyTemplate), nil
   }
   ```

4. Templates are automatically embedded via `go:embed` in `main.go`

### Modifying package.json Generation

Edit `internal/template/pkgjson.go`:
- `buildScripts()` — npm scripts
- `buildDeps()` — dependencies
- `buildDevDeps()` — devDependencies

Use typed constants:
```go
if cfg.Linter == string(constants.LinterBiome) {
    deps["@biomejs/biome"] = "2.2.0"
}
```

### Adding a New Component

1. Define interface in `internal/core/interfaces/interfaces.go`:
   ```go
   type MyComponent interface {
       DoSomething(ctx context.Context, input string) error
   }
   ```

2. Create implementation:
   ```go
   // internal/mypackage/component.go
   type MyComponentImpl struct {
       dependency SomeDependency
   }
   
   func NewMyComponent(dep SomeDependency) *MyComponentImpl {
       return &MyComponentImpl{dependency: dep}
   }
   
   func (m *MyComponentImpl) DoSomething(ctx context.Context, input string) error {
       // Implementation
       return nil
   }
   ```

3. Wire in `cmd/root.go`:
   ```go
   myComponent := mypackage.NewMyComponent(someDep)
   runner := cli.NewRunner(..., myComponent)
   ```

## Cross-Compilation

Build for all platforms at once:

```bash
make release
```

This creates binaries in `dist/` for:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

Or build individually:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o create-next-go-app-linux

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o create-next-go-app-macos

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o create-next-go-app-macos-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o create-next-go-app.exe
```

## Debugging

### Verbose Output

Add print statements or use a debugger:

```go
import "github.com/yeasin2002/create-next-app-go_prototype/internal/core/errors"

// Use structured errors for better debugging
return errors.NewWithPath("operation name", filePath, err)
```

### Common Issues

**Template not found:**
- Ensure templates are in the `templates/` directory
- Check that `go:embed templates` is in `main.go`
- Rebuild after adding new templates
- Verify template name matches constants in `internal/core/constants/`

**Flags not working:**
- Check flag is registered in `cmd/root.go` `init()`
- Verify parsing in `internal/cli/flags.go`
- Use `cmd.Flags().Changed()` to detect explicit flags
- Boolean flags need explicit `--flag` or `--no-flag`

**Dependency injection errors:**
- Ensure all dependencies are created in `cmd/root.go` `initializeRunner()`
- Check that interfaces match implementations
- Verify constructor signatures

**Build errors:**
- Run `go mod tidy` to sync dependencies
- Check for import cycles
- Ensure all interfaces are properly implemented

### Testing Your Changes

```bash
# Build
go build -o create-next-app.exe

# Test basic functionality
.\create-next-app.exe test-app --yes --skip-install --skip-git

# Test with different configurations
.\create-next-app.exe test-js --javascript --no-tailwind --yes --skip-install
.\create-next-app.exe test-biome --biome --yes --skip-install

# Clean up
Remove-Item -Recurse -Force test-app, test-js, test-biome
```

## Pull Request Guidelines

1. Keep PRs focused on a single change
2. Update documentation if needed
3. Add tests for new functionality (see `TESTING_GUIDE.md`)
4. Ensure `go build` succeeds
5. Test the CLI manually with various flag combinations
6. Follow the existing architecture patterns:
   - Use interfaces for new components
   - Inject dependencies via constructors
   - Use typed constants instead of strings
   - Wrap errors with context
7. Run `go fmt` before committing
8. Update `CHANGELOG.md` if applicable

### Code Review Checklist

Before submitting:
- [ ] Code follows Go conventions
- [ ] New components have interfaces
- [ ] Dependencies are injected
- [ ] Errors are properly wrapped
- [ ] Constants are used instead of magic strings
- [ ] Documentation is updated
- [ ] Manual testing completed
- [ ] No global mutable state introduced

## Reporting Issues

When reporting bugs, include:
- Go version (`go version`)
- Operating system
- Command that caused the issue
- Expected vs actual behavior
- Error messages (if any)

## Questions?

Open an issue for questions or discussions about the project.

For architecture questions, see:
- `REFACTORING_SUMMARY.md` - Overview of the refactored architecture
- `MIGRATION_GUIDE.md` - Guide for working with the new codebase
- `TESTING_GUIDE.md` - Testing strategy and examples
