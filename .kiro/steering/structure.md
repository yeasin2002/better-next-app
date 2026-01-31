# Project Structure

## Current State

The project is in early development with minimal implementation:

```
better-next-app/
├── .git/                  # Git repository
├── .kiro/                 # Kiro configuration
│   └── steering/          # AI assistant steering rules
├── go.mod                 # Go module definition
├── main.go                # Entry point (minimal placeholder)
├── PROJECT-GUIDE.md       # Comprehensive implementation guide
└── README.md              # Project documentation
```

## Planned Architecture

Based on PROJECT-GUIDE.md, the intended structure follows Go best practices:

```
better-next-app/
├── main.go                # Entry point with signal handlers
├── cmd/                   # Cobra command definitions
│   └── root.go            # Root command and flag registration
├── internal/              # Private business logic
│   ├── config/            # Configuration management
│   ├── prompt/            # Interactive prompts (Huh)
│   ├── validate/          # Validation logic
│   ├── template/          # Template installation
│   ├── example/           # GitHub example downloading
│   ├── install/           # Post-installation tasks
│   └── util/              # Shared utilities
└── templates/             # Embedded template files
    ├── app/               # App Router templates
    ├── app-tw/            # App Router + Tailwind
```

## Key Architectural Principles

### Internal Package

The `internal/` directory prevents external imports, keeping implementation details private. This is a Go convention for encapsulation.

### Module Organization

Each subdirectory under `internal/` represents a distinct concern:

- **config** - Configuration struct, defaults, and preference persistence
- **prompt** - User interaction and input collection
- **validate** - Input validation separated from prompts
- **template** - Template copying, transformations, and package.json generation
- **example** - GitHub API integration and tar extraction
- **install** - Package manager execution, TypeScript generation, git init
- **util** - Cross-cutting concerns (colors, filesystem, command execution)

### Template Embedding

Templates are embedded using Go's `embed` directive, allowing the entire application to be distributed as a single binary with no external file dependencies.

## File Naming Conventions

- Go source files: lowercase with underscores (e.g., `package_json.go`)
- Test files: `_test.go` suffix
- Package names: single lowercase word matching directory name
- Exported identifiers: PascalCase
- Unexported identifiers: camelCase
