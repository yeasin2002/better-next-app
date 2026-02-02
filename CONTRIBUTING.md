# Contributing to Better Next App

Thanks for your interest in contributing! This guide will help you get started.

## Prerequisites

- **Go 1.21+** - [Download](https://go.dev/dl/)
- **Task** - [Install](https://taskfile.dev/installation/)
- **Git** - [Download](https://git-scm.com/downloads)

## Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/yourusername/better-next-app.git
cd better-next-app

# Install tools, dependencies, and setup git hooks
task setup
```

This will install:
- golangci-lint (linter)
- goimports (import formatter)
- commitlint (commit message validator)
- Lefthook (git hooks manager)
- GoReleaser (release tool)
- All Go dependencies
- Git hooks (pre-commit, commit-msg, pre-push)

### 2. Development Workflow

```bash
# Run the application
task dev

# Run with hot reload
task dev:watch

# Build for current platform
task build

# Run all checks before committing
task check
```

## Project Structure

```
better-next-app/
├── main.go              # Entry point
├── cmd/                 # Cobra commands
│   └── root.go
├── internal/            # Private code
│   ├── config/         # Configuration
│   ├── prompt/         # Interactive prompts
│   ├── validate/       # Validation logic
│   ├── template/       # Template installation
│   ├── install/        # Post-installation
│   └── util/           # Utilities
└── templates/          # Embedded Next.js templates
    ├── app/           # App Router
    ├── app-tw/        # App Router + Tailwind
    ├── app-empty/     # Minimal templates
    └── app-api/       # API-only
```

## Common Tasks

See all available tasks:
```bash
task --list
```

### Development
- `task dev` - Run the application
- `task dev:watch` - Run with hot reload

### Building
- `task build` - Build for current platform
- `task build:all` - Build for all platforms (Linux, macOS, Windows)

### Testing
- `task test` - Run all tests
- `task test:coverage` - Generate coverage report

### Code Quality
- `task fmt` - Format code
- `task lint` - Run linter
- `task check` - Run all checks (format, lint, test)

### Cleaning
- `task clean` - Clean build artifacts

## Making Changes

1. **Create a branch**
   ```bash
   git checkout -b feature/my-feature
   ```

2. **Make your changes**

3. **Test your changes**
   ```bash
   task dev -- test-app --yes --skip-install
   ```

4. **Commit** (hooks will validate your message and run checks)
   ```bash
   git add .
   git commit -m "feat: add my feature"
   ```
   
   Commit message must follow [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` - New feature
   - `fix:` - Bug fix
   - `docs:` - Documentation
   - `test:` - Tests
   - `chore:` - Maintenance
   - `refactor:` - Code refactoring
   - `perf:` - Performance improvements
   - `ci:` - CI/CD changes
   - `build:` - Build system changes
   
   Git hooks will automatically:
   - **commit-msg**: Validate commit message format with commitlint
   - **pre-commit**: Format code, run vet, run linter, run tests
   - **pre-push**: Run full check suite

5. **Push** (pre-push hook will run full checks)
   ```bash
   git push origin feature/my-feature
   ```

## Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` for formatting (automatic with `task fmt`)
- Add comments for exported functions
- Keep functions focused and small

### Naming Conventions
- **Packages**: lowercase, single word (`config`, `prompt`)
- **Files**: lowercase with underscores (`package_json.go`)
- **Exported**: PascalCase (`InstallTemplate`)
- **Unexported**: camelCase (`copyFiles`)

## Adding Features

### New CLI Flag

1. Register in `cmd/root.go`:
   ```go
   rootCmd.Flags().Bool("my-flag", false, "Description")
   ```

2. Use in your code:
   ```go
   myFlag, _ := cmd.Flags().GetBool("my-flag")
   ```

### New Template

1. Create template directory:
   ```
   templates/my-template/
   ├── js/
   └── ts/
   ```

2. Templates are auto-embedded via `go:embed` in `main.go`

3. Update template selection logic in `templates/index.go`

## Pull Request Guidelines

1. **Run checks**: `task check` before submitting
2. **Keep PRs focused**: One feature/fix per PR
3. **Add tests**: For new functionality
4. **Update docs**: If adding features
5. **Follow conventions**: Use Go best practices

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/) format enforced by commitlint:

```
<type>(<scope>): <subject>

# Examples:
feat: add new feature
fix(template): resolve bug in template copying
docs: update README with installation steps
test(util): add unit tests for validation
chore(deps): update dependencies
```

**Allowed types**: feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert

**Allowed scopes** (optional): cli, config, prompt, template, util, deps, release

**Rules**:
- Header: 10-100 characters
- Description: minimum 3 characters
- Body/footer lines: max 100 characters

See [docs/git-hooks.md](./docs/git-hooks.md) for detailed commit message guidelines.

## Troubleshooting

### Lefthook not running

```bash
# Reinstall hooks
lefthook install

# Or skip hooks temporarily
git commit --no-verify
```

### commitlint not found

```bash
# Install via Task
task setup

# Or manually
go install github.com/conventionalcommit/commitlint@latest

# Ensure $GOPATH/bin is in your PATH
```

### Task not found

Install Task:
```bash
# macOS
brew install go-task/tap/go-task

# Windows
scoop install task

# Cross-platform
npm install -g @go-task/cli
```

### golangci-lint not found

```bash
# Install via Task
task setup:tools

# Or manually
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Build fails

```bash
task clean
task deps
task build
```

## Getting Help

- **Documentation**: See [PROJECT-GUIDE.md](./PROJECT-GUIDE.md)
- **Issues**: Open an issue on GitHub
- **Tasks**: Run `task --list` for all commands
