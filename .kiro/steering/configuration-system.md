---
inclusion: always
---

# Configuration System

## Config Struct

The main configuration struct must include all project options:

```go
type Config struct {
    // Project Identity
    ProjectName  string
    ProjectPath  string // Absolute path

    // Language & Framework
    TypeScript   bool
    AppRouter    bool   // App Router vs Pages Router
    APIOnly      bool   // API-only project (no React)

    // Styling
    Tailwind     bool

    // Linting & Formatting
    Linter       string // "eslint", "biome", or "none"

    // Project Structure
    SrcDir       bool
    ImportAlias  string // Default: "@/*"
    EmptyTemplate bool  // Minimal template

    // Bundler
    Bundler      string // "turbopack", "webpack", or "rspack"

    // Features
    ReactCompiler bool

    // Package Manager
    PackageManager string // "npm", "pnpm", "yarn", or "bun"
    SkipInstall    bool

    // Git
    SkipGit      bool

    // Example Mode
    Example      string // Example name or GitHub URL
    ExamplePath  string // Path within repo (for subdirectories)
}
```

## User Preferences

Preferences are saved between sessions:

```go
type Preferences struct {
    TypeScript         bool   `json:"typescript"`
    Linter            string `json:"linter"`
    Tailwind          bool   `json:"tailwind"`
    AppRouter         bool   `json:"appRouter"`
    SrcDir            bool   `json:"srcDir"`
    ImportAlias       string `json:"importAlias"`
    CustomizeAlias    bool   `json:"customizeAlias"`
    EmptyTemplate     bool   `json:"emptyTemplate"`
    DisableGit        bool   `json:"disableGit"`
    ReactCompiler     bool   `json:"reactCompiler"`
}
```

## Configuration Priority

Merge config sources in this order (highest to lowest priority):

1. **CLI Flags** - Explicit user input via command line
2. **User Preferences** - Saved from previous sessions
3. **Defaults** - Recommended Next.js defaults

```go
func MergeConfig(flags *Config, prefs *Preferences) *Config {
    // CLI flags override preferences
    // Preferences override defaults
    // Return merged config
}
```

## Required Functions

```go
// Load preferences from disk
func LoadPreferences() (*Preferences, error)

// Save preferences to disk
func SavePreferences(prefs *Preferences) error

// Check if preferences exist
func HasPreferences() bool

// Clear all preferences
func ClearPreferences() error

// Get default config
func DefaultConfig() *Config
```

## CLI Flags

All flags must be registered in Cobra:

**Project Configuration:**

- `--typescript` / `--javascript`
- `--tailwind` / `--no-tailwind`
- `--app` / `--pages`
- `--src-dir` / `--no-src-dir`
- `--import-alias <string>`

**Linting & Tools:**

- `--eslint` / `--biome` / `--no-lint`
- `--react-compiler` / `--no-react-compiler`

**Bundler Options:**

- `--turbo` / `--webpack` / `--rspack`

**Package Manager:**

- `--use-npm` / `--use-pnpm` / `--use-yarn` / `--use-bun`
- `--skip-install`

**Git Options:**

- `--skip-git`

**Example Mode:**

- `--example <name-or-url>`
- `--example-path <path>`

**Automation:**

- `--yes` (skip prompts, use defaults)
- `--reset-preferences`
- `--empty` (minimal template)

## Package Manager Detection

Detect package manager from environment:

```go
func DetectPackageManager() string {
    userAgent := os.Getenv("npm_config_user_agent")

    if strings.Contains(userAgent, "pnpm") {
        return "pnpm"
    } else if strings.Contains(userAgent, "yarn") {
        return "yarn"
    } else if strings.Contains(userAgent, "bun") {
        return "bun"
    }

    return "npm" // default
}
```
