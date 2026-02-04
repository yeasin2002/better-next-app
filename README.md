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
- 🎨 **App Router Only** - Focused on Next.js App Router with/without Tailwind
- 📦 **Smart Package Manager Detection** - Auto-detects npm, pnpm, yarn, or bun
- 🔧 **Flexible Configuration** - TypeScript, ESLint, Biome, React Compiler support
- 🌐 **GitHub Examples** - Download official examples or custom repositories
- 💾 **Preference Persistence** - Save your choices for future projects
- 🚀 **Future Extensibility** - Planned support for database integration, ORM, and more

## Installation

### Using Package Managers (Recommended)

```bash
# npm
npx create-better-next-app@latest

# pnpm
pnpm dlx create-better-next-app@latest

# bun
bunx create-better-next-app@latest

# yarn
yarn create better-next-app
```

### Using Go Install

```bash
go install github.com/yeasin2002/better-next-app@latest
```

### Direct Download

Download the latest binary for your platform from the [releases page](https://github.com/yeasin2002/better-next-app/releases).

### Build from Source

```bash
git clone https://github.com/yeasin2002/better-next-app.git
cd better-next-app
task build
```

## Usage

### Interactive Mode

```bash
npx create better-next-app@latest
```

On installation, you'll see the following prompts:

```
What is your project named? my-app
Would you like to use the recommended Next.js defaults?
    Yes, use recommended defaults - TypeScript, ESLint, Tailwind CSS, App Router, Turbopack
    No, reuse previous settings
    No, customize settings - Choose your own preferences
```

If you choose to customize settings, you'll see:

```
Would you like to use TypeScript? No / Yes
Which linter would you like to use? ESLint / Biome / None
Would you like to use React Compiler? No / Yes
Would you like to use Tailwind CSS? No / Yes
Would you like your code inside a `src/` directory? No / Yes
Would you like to customize the import alias (`@/*` by default)? No / Yes
What import alias would you like configured? @/*
```

### Non-Interactive Mode

```bash
better-next-app my-app --typescript --tailwind --eslint
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
- `--tailwind` / `--no-tailwind` - Include Tailwind CSS
- `--src-dir` / `--no-src-dir` - Use `src/` directory
- `--import-alias <string>` - Custom import alias (default: `@/*`)
- `--empty` - Minimal template with no boilerplate

> Note: This CLI only supports Next.js App Router. Pages Router is not supported.

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
- [Task](https://taskfile.dev) - Task runner (optional but recommended)

### Setup

```bash
# Clone the repository
git clone https://github.com/yeasin2002/better-next-app.git
cd better-next-app

# Setup development environment (installs tools and git hooks)
task setup

# Or manually
go mod download
```

### Common Tasks

```bash
# Run locally
task dev

# Build for current platform
task build

# Build for all platforms
task build:all

# Run tests
task test

# Run all checks (format, lint, test)
task check

# Test release locally (no git tag required)
task release:snapshot
```

### NPM Publishing

```bash
# Setup NPM publishing (first time)
task npm:setup

# Test NPM package locally
task npm:test

# Publish to NPM manually
task npm:publish
```

**Documentation:**
- Quick Start (5 min): [docs/npm-quick-start.md](./docs/npm-quick-start.md)
- Token Generation Guide: [docs/npm-token-guide.md](./docs/npm-token-guide.md)
- Full Guide: [docs/npm-publishing.md](./docs/npm-publishing.md)
- Troubleshooting: [docs/TROUBLESHOOTING-NPM.md](./docs/TROUBLESHOOTING-NPM.md)

### Releasing

This project uses [GoReleaser](https://goreleaser.com) for automated releases:

```bash
# Create and push a tag
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

# GitHub Actions will automatically build and release
```

See [docs/releasing.md](./docs/releasing.md) for detailed release instructions.

## Documentation

- [PROJECT-GUIDE.md](./PROJECT-GUIDE.md) - Comprehensive implementation guide
- [docs/releasing.md](./docs/releasing.md) - Release process and versioning
- [Steering Rules](./.kiro/steering/) - Development guidelines and patterns

## Contributing

Contributions are welcome! Please read the [PROJECT-GUIDE.md](./PROJECT-GUIDE.md) for architecture details and implementation guidelines.

## License

MIT

## Roadmap

Future enhancements planned:
- Database integration options
- ORM setup (Prisma, Drizzle, etc.)
- Authentication scaffolding
- Additional tooling integrations

## Acknowledgments

This project is inspired by and maintains feature parity with [create-next-app](https://github.com/vercel/next.js/tree/canary/packages/create-next-app) from the Next.js team, with a focus on App Router and modern Next.js development.
