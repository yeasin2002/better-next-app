---
inclusion: always
---

# Go Architecture & Project Structure

## Core Principles

- **Feature Parity**: All TypeScript features must work identically
- **Go Idioms**: Leverage Go's strengths (concurrency, single binary, embed)
- **Extensibility**: Design for future plugin architecture
- **Template Preservation**: Keep templates unchanged from TypeScript version

## Project Structure

Use the `internal/` package pattern to keep implementation details private:

```
create-next-app-go/
├── main.go                    # Entry point with signal handlers
├── cmd/                       # Cobra command definitions
├── internal/                  # All business logic (private)
│   ├── config/               # Configuration management
│   ├── prompt/               # Interactive prompts (Huh)
│   ├── validate/             # Validation logic
│   ├── template/             # Template installation
│   ├── example/              # GitHub example downloading
│   ├── install/              # Post-installation tasks
│   └── util/                 # Shared utilities
└── templates/                # Embedded template files
```

## Module Responsibilities

**cmd/** - CLI argument parsing and flag registration using Cobra

**internal/config/** - Config struct, defaults, preference persistence (Viper)

**internal/prompt/** - Interactive prompts using Huh (isolated for testability)

**internal/validate/** - Validation logic (npm names, directories, network)

**internal/template/** - Template copying, transformations, package.json generation

**internal/example/** - GitHub API integration, tar extraction

**internal/install/** - Package manager execution, TypeScript typegen, git init

**internal/util/** - Colors (Lipgloss), command execution, filesystem helpers

## Application Flow

Linear pipeline with one branching point:

```
CLI Entry → Config Builder → Validation →
[Example Download OR Template Install] →
Post-Install → Success Output
```

## Error Handling

Define custom error types for better messages:

```go
type ValidationError struct {
    Field    string
    Problems []string
}

type DirectoryError struct {
    Path             string
    ConflictingFiles []string
}
```

Always provide user-friendly error messages with actionable solutions.
