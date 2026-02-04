---
inclusion: always
---

# Project Structure

## Current State

The project has basic scaffolding with core modules implemented:

```
better-next-app/
├── .git/                  # Git repository
├── .github/               # GitHub Actions workflows
├── .kiro/                 # Kiro configuration
│   └── steering/          # AI assistant steering rules
├── cmd/                   # Cobra command definitions
│   └── root.go            # Root command (basic structure)
├── internal/              # Private business logic
│   ├── config/            # Configuration management (implemented)
│   │   ├── config.go      # Main Config struct
│   │   ├── defaults.go    # Default values
│   │   └── preferences.go # Viper-based preferences
│   ├── prompt/            # Interactive prompts (implemented)
│   │   ├── prompt.go      # Huh-based prompts
│   │   ├── setup.go       # Initial setup choice
│   │   └── validate.go    # Input validation
│   └── util/              # Shared utilities (partial)
│       ├── colors.go      # Lipgloss color definitions
│       └── validate_npm_package_name.go # NPM name validation
├── templates/             # Embedded template files (complete)
│   ├── app/               # App Router templates
│   ├── app-tw/            # App Router + Tailwind
│   ├── app-empty/         # Minimal App Router
│   ├── app-tw-empty/      # Minimal App Router + Tailwind
│   └── app-api/           # API-only templates
├── docs/                  # Documentation
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
├── main.go                # Entry point with signal handlers
├── Taskfile.yml           # Task runner configuration
├── lefthook.yml           # Git hooks configuration
├── .goreleaser.yaml       # GoReleaser configuration
├── .golangci.yml          # Linter configuration
├── PROJECT-GUIDE.md       # Comprehensive implementation guide
└── README.md              # Project documentation
```

## Implementation Status

### Completed Modules
- ✅ **main.go** - Entry point with signal handlers and template embedding
- ✅ **cmd/root.go** - Basic Cobra command structure
- ✅ **internal/config/** - Complete configuration management with Viper
- ✅ **internal/prompt/** - Interactive prompts using Huh
- ✅ **internal/util/colors.go** - Terminal styling with Lipgloss
- ✅ **internal/util/validate_npm_package_name.go** - NPM name validation
- ✅ **templates/** - All template files (app, app-tw, app-empty, app-tw-empty, app-api)

### Pending Modules
- ⏳ **cmd/root.go** - Flag registration and main application logic
- ⏳ **internal/validate/** - Directory and network validation
- ⏳ **internal/template/** - Template installation and transformations
- ⏳ **internal/example/** - GitHub example downloading
- ⏳ **internal/install/** - Post-installation tasks (deps, typegen, git)
- ⏳ **internal/util/** - Additional utilities (filesystem, command execution)

## Target Architecture

The final structure follows Go best practices:

```
better-next-app/
├── main.go                # Entry point with signal handlers
├── cmd/                   # Cobra command definitions
│   └── root.go            # Root command and flag registration
├── internal/              # Private business logic
│   ├── config/            # Configuration management ✅
│   ├── prompt/            # Interactive prompts (Huh) ✅
│   ├── validate/          # Validation logic ⏳
│   ├── template/          # Template installation ⏳
│   ├── example/           # GitHub example downloading ⏳
│   ├── install/           # Post-installation tasks ⏳
│   └── util/              # Shared utilities (partial) ⏳
└── templates/             # Embedded template files ✅
    ├── app/               # App Router templates
    ├── app-tw/            # App Router + Tailwind
    ├── app-empty/         # Minimal App Router
    ├── app-tw-empty/      # Minimal App Router + Tailwind
    └── app-api/           # API-only templates
```

## Key Architectural Principles

### Internal Package

The `internal/` directory prevents external imports, keeping implementation details private. This is a Go convention for encapsulation.

### Module Organization

Each subdirectory under `internal/` represents a distinct concern:

- **config** ✅ - Configuration struct, defaults, and preference persistence using Viper
- **prompt** ✅ - User interaction and input collection using Huh
- **validate** ⏳ - Input validation separated from prompts (directory, network checks)
- **template** ⏳ - Template copying, transformations, and package.json generation
- **example** ⏳ - GitHub API integration and tar extraction
- **install** ⏳ - Package manager execution, TypeScript generation, git init
- **util** (partial) - Cross-cutting concerns (colors ✅, npm validation ✅, filesystem ⏳, command execution ⏳)

### Template Embedding

Templates are embedded using Go's `embed` directive, allowing the entire application to be distributed as a single binary with no external file dependencies.

## File Naming Conventions

- Go source files: lowercase with underscores (e.g., `package_json.go`)
- Test files: `_test.go` suffix
- Package names: single lowercase word matching directory name
- Exported identifiers: PascalCase
- Unexported identifiers: camelCase
