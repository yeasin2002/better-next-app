# Product Overview

Better Next App is a high-performance CLI tool for scaffolding Next.js projects, written in Go. It's a complete rewrite of `create-next-app` that provides faster startup times, single binary distribution, and feature parity with the original TypeScript implementation.

## Key Value Propositions

- Single binary distribution with no Node.js required to run the CLI
- Faster startup through native compilation
- Cross-platform support (Windows, macOS, Linux)
- Built-in concurrency using Go's goroutines
- Template embedding via Go's `embed` directive

## Core Features

- Interactive and non-interactive modes
- Multiple templates (App Router, Pages Router, with/without Tailwind)
- Smart package manager detection (npm, pnpm, yarn, bun)
- Flexible configuration (TypeScript, ESLint, Biome, React Compiler)
- GitHub example downloading
- Preference persistence
- Full feature parity with the original TypeScript version
