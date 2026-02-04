# Create Better Next App

A high-performance CLI tool for scaffolding Next.js projects, written in Go.

## Quick Start

```bash
# Using npx (recommended)
npx create-better-next-app@latest

# Using pnpm
pnpm dlx create-better-next-app@latest

# Using bun
bunx create-better-next-app@latest

# Using yarn
yarn create better-next-app
```

## Features

- ✨ **Interactive or Non-Interactive** - Use prompts or CLI flags
- 🎨 **App Router Only** - Focused on Next.js App Router
- 📦 **Smart Package Manager Detection** - Auto-detects npm, pnpm, yarn, or bun
- 🔧 **Flexible Configuration** - TypeScript, ESLint, Biome, React Compiler support
- 🌐 **GitHub Examples** - Download official examples or custom repositories
- 💾 **Preference Persistence** - Save your choices for future projects
- ⚡ **Fast** - Written in Go for instant startup

## Usage

### Interactive Mode

```bash
npx create-better-next-app@latest
```

You'll be prompted for:
- Project name
- TypeScript or JavaScript
- Linter choice (ESLint, Biome, or None)
- Tailwind CSS
- React Compiler
- `src/` directory
- Custom import alias

### Non-Interactive Mode

```bash
npx create-better-next-app@latest my-app --typescript --tailwind --eslint
```

### Using Examples

```bash
# Official Next.js example
npx create-better-next-app@latest my-app --example with-tailwindcss

# Custom GitHub repository
npx create-better-next-app@latest my-app --example https://github.com/user/repo
```

## CLI Options

- `--typescript` / `--javascript` - Language choice
- `--tailwind` / `--no-tailwind` - Include Tailwind CSS
- `--eslint` / `--biome` / `--no-lint` - Linter choice
- `--react-compiler` - Enable React Compiler
- `--src-dir` - Use `src/` directory
- `--import-alias <string>` - Custom import alias
- `--empty` - Minimal template
- `--use-npm` / `--use-pnpm` / `--use-yarn` / `--use-bun` - Package manager
- `--skip-install` - Skip dependency installation
- `--skip-git` - Skip git initialization
- `--yes` - Skip all prompts and use defaults

## Documentation

For more information, visit the [GitHub repository](https://github.com/yeasin2002/better-next-app).

## License

MIT
