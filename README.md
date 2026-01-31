# Better Next App

A modern, high-performance CLI tool for scaffolding Next.js projects, written in Go. This is a complete rewrite of `create-next-app` that provides faster startup times, single binary distribution, and feature parity with the original TypeScript implementation.

## Why Go?

- **Single Binary Distribution** - No Node.js required to run the CLI
- **Faster Startup** - Native compilation means instant execution
- **Cross-Platform** - Compile once for Windows, macOS, and Linux
- **Built-in Concurrency** - Leverage Go's goroutines for parallel operations
- **Template Embedding** - All templates bundled in the binary using `embed` directive

## Features

- ✨ **Interactive or Non-Interactive** - Use prompts or CLI flags
- 🎨 **Multiple Templates** - App Router, Pages Router, with/without Tailwind
- 📦 **Smart Package Manager Detection** - Auto-detects npm, pnpm, yarn, or bun
- 🔧 **Flexible Configuration** - TypeScript, ESLint, Biome, React Compiler support
- 🌐 **GitHub Examples** - Download official examples or custom repositories
- 💾 **Preference Persistence** - Save your choices for future projects
- 🎯 **Full Feature Parity** - All features from the original TypeScript version

## Installation

### Using Go Install

```bash
go install github.com/yourusername/better-next-app@latest
```

### Direct Download

Download the latest binary for your platform from the [releases page](https://github.com/yourusername/better-next-app/releases).

### Build from Source

```bash
git clone https://github.com/yourusername/better-next-app.git
cd better-next-app
go build -o better-next-app
```

## Usage

### Interactive Mode

```bash
better-next-app my-app
```

### Non-Interactive Mode

```bash
better-next-app my-app --typescript --tailwind --app --eslint
```

### Using Examples

```bash
# Official Next.js example
better-next-app my-app --example with-tailwindcss

# Custom GitHub repository
better-next-app my-app --example https://github.com/user/repo
```

## CLI Options

### Project Configuration

- `--typescript` / `--javascript` - Language choice
- `--app` / `--pages` - Router type (App Router or Pages Router)
- `--tailwind` / `--no-tailwind` - Include Tailwind CSS
- `--src-dir` / `--no-src-dir` - Use `src/` directory
- `--import-alias <string>` - Custom import alias (default: `@/*`)
- `--empty` - Minimal template with no boilerplate

### Linting & Tools

- `--eslint` - Use ESLint
- `--biome` - Use Biome
- `--no-lint` - Skip linter setup
- `--react-compiler` / `--no-react-compiler` - Enable React Compiler

### Bundler Options

- `--turbo` - Use Turbopack (default)
- `--webpack` - Use Webpack
- `--rspack` - Use Rspack

### Package Manager

- `--use-npm` - Use npm
- `--use-pnpm` - Use pnpm
- `--use-yarn` - Use Yarn
- `--use-bun` - Use Bun
- `--skip-install` - Skip dependency installation

### Git Options

- `--skip-git` - Skip git initialization

### Example Mode

- `--example <name-or-url>` - Use an example template
- `--example-path <path>` - Path within repository (for monorepos)

### Automation

- `--yes` - Skip all prompts and use defaults
- `--reset-preferences` - Clear saved preferences

## Project Structure

```
better-next-app/
├── main.go                    # Entry point with signal handlers
├── cmd/                       # Cobra command definitions
├── internal/                  # Business logic (private)
│   ├── config/               # Configuration management
│   ├── prompt/               # Interactive prompts (Huh)
│   ├── validate/             # Validation logic
│   ├── template/             # Template installation
│   ├── example/              # GitHub example downloading
│   ├── install/              # Post-installation tasks
│   └── util/                 # Shared utilities
└── templates/                # Embedded template files
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Huh](https://github.com/charmbracelet/huh) - Interactive prompts
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Doublestar](https://github.com/bmatcuk/doublestar) - File globbing

## Development

### Prerequisites

- Go 1.21 or higher

### Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/better-next-app.git
cd better-next-app

# Install dependencies
go mod download

# Run locally
go run main.go my-test-app
```

### Building

```bash
# Build for current platform
go build -o better-next-app

# Build with optimizations
go build -ldflags="-s -w" -o better-next-app

# Cross-compile for multiple platforms
GOOS=linux GOARCH=amd64 go build -o better-next-app-linux
GOOS=darwin GOARCH=amd64 go build -o better-next-app-macos
GOOS=windows GOARCH=amd64 go build -o better-next-app.exe
```

## Documentation

- [PROJECT-GUIDE.md](./PROJECT-GUIDE.md) - Comprehensive implementation guide
- [Steering Rules](./.kiro/steering/) - Development guidelines and patterns

## Contributing

Contributions are welcome! Please read the [PROJECT-GUIDE.md](./PROJECT-GUIDE.md) for architecture details and implementation guidelines.

## License

MIT

## Acknowledgments

This project is inspired by and maintains feature parity with [create-next-app](https://github.com/vercel/next.js/tree/canary/packages/create-next-app) from the Next.js team.
