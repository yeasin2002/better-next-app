package config

// Config holds all configuration for project creation
type Config struct {
	// Project Identity
	ProjectName string
	ProjectPath string // Absolute path

	// Language & Framework
	TypeScript bool
	AppRouter  bool // Always true - Pages Router removed
	APIOnly    bool // API-only project (no React)

	// Styling
	Tailwind bool

	// Linting & Formatting
	Linter string // "eslint", "biome", or "none"

	// Project Structure
	SrcDir        bool
	ImportAlias   string // Default: "@/*"
	EmptyTemplate bool   // Minimal template

	// Bundler
	Bundler string // "turbopack", "webpack", or "rspack"

	// Features
	ReactCompiler bool

	// Package Manager
	PackageManager string // "npm", "pnpm", "yarn", or "bun"
	SkipInstall    bool

	// Git
	SkipGit bool

	// Example Mode
	Example     string // Example name or GitHub URL
	ExamplePath string // Path within repo (for subdirectories)
}

// New creates a new Config with default values
func New() *Config {
	return &Config{
		ImportAlias: "@/*",
		Bundler:     "turbopack",
		AppRouter:   true,
	}
}
