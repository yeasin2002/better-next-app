# Create Next App - Go CLI Build Guide

A comprehensive guide for building the create-next-app CLI prototype in Go, covering architecture, folder structure, package selection, and implementation strategy.

---

## Table of Contents

1. [Overview](#overview)
2. [Project Architecture](#project-architecture)
3. [Folder Structure](#folder-structure)
4. [Required Dependencies](#required-dependencies)
5. [Core Data Structures](#core-data-structures)
6. [Module Implementation Guide](#module-implementation-guide)
7. [Template System](#template-system)
8. [Application Flow](#application-flow)
9. [Error Handling](#error-handling)
10. [Building and Distribution](#building-and-distribution)

---

## Overview

### What This CLI Does

A scaffolding tool that creates new Next.js projects with:

- Interactive or non-interactive configuration
- Multiple template options (App Router, Pages Router, Tailwind, etc.)
- Support for official examples or custom GitHub repositories
- Automatic dependency installation
- Git repository initialization
- User preference persistence

### Why Go?

The Go rewrite provides several advantages:

- Single binary distribution (no Node.js needed for the CLI itself)
- Faster startup time
- Cross-platform compilation
- Built-in concurrency support
- Native template embedding via `embed` directive

---

## Project Architecture

### High-Level Design

The application follows a linear pipeline with one key branching point:

```
CLI Entry → Config Builder → Validation → [Example Download OR Template Install] → Post-Install → Success Output
```

**Flow Details:**

1. **CLI Entry Point** - Receives user input through command-line arguments
2. **Configuration Builder** - Combines CLI flags, prompts, preferences, and defaults
3. **Validation Layer** - Checks project name, directory permissions, network connectivity
4. **Branching Point** - Either downloads example from GitHub OR copies built-in templates
5. **Post-Installation** - Runs dependency installation, TypeScript generation, git init
6. **Success Output** - Displays instructions and next steps

### Core Principles

- **Feature Parity**: All TypeScript features must work identically
- **Go Idioms**: Leverage Go's strengths (concurrency, single binary, embed)
- **Extensibility**: Design for future plugin architecture
- **Template Preservation**: Keep templates unchanged from TypeScript version

---

## Folder Structure (Just for example)

```
create-next-app-go/
├── main.go                           # Entry point
├── go.mod                            # Go module definition
├── go.sum                            # Dependency checksums
│
├── cmd/
│   └── root.go                       # Cobra root command definition
│
├── internal/
│   ├── config/
│   │   ├── config.go                 # Main configuration struct
│   │   ├── defaults.go               # Default values
│   │   └── preferences.go            # Viper-based user preferences
│   │
│   ├── prompt/
│   │   ├── prompt.go                 # Huh-based prompts
│   │   ├── setup.go                  # Initial setup choice
│   │   └── validate.go               # Input validation functions
│   │
│   ├── validate/
│   │   ├── name.go                   # NPM name validation
│   │   ├── directory.go              # Directory checks
│   │   └── network.go                # Network connectivity detection
│   │
│   ├── template/
│   │   ├── install.go                # Template installation orchestration
│   │   ├── copy.go                   # File copying with transformations
│   │   ├── transform.go              # File content transformations
│   │   └── packagejson.go            # Package.json generation
│   │
│   ├── example/
│   │   ├── download.go               # GitHub API integration
│   │   ├── extract.go                # Tar extraction
│   │   └── github.go                 # Repository info parsing
│   │
│   ├── install/
│   │   ├── deps.go                   # Package manager execution
│   │   ├── typegen.go                # TypeScript type generation
│   │   └── git.go                    # Git initialization
│   │
│   └── util/
│       ├── colors.go                 # Lipgloss color definitions
│       ├── exec.go                   # Command execution helpers
│       ├── fs.go                     # Filesystem utilities
│       └── spinner.go                # Optional spinner (if using Bubbles)
│
└── templates/                        # Embedded template files
    ├── app/                          # App Router templates
    ├── app-tw/                       # App Router + Tailwind
    ├── default/                      # Pages Router templates
    └── default-tw/                   # Pages Router + Tailwind
```

### Key Directories Explained

**`cmd/`** - Contains Cobra command definitions. This is where CLI arguments and flags are registered.

**`internal/`** - All business logic lives here. The `internal` package prevents external imports, keeping implementation details private.

**`internal/config/`** - Configuration management including the main Config struct, default values, and preference persistence using Viper.

**`internal/prompt/`** - Interactive prompts using Huh. Each prompt function is isolated for testability.

**`internal/validate/`** - Validation logic separated from prompts. Includes npm name validation, directory safety checks, and network detection.

**`internal/template/`** - Template installation logic including file copying, transformations (src directory, import aliases), and package.json generation.

**`internal/example/`** - GitHub example downloading including API calls, tar extraction, and repository URL parsing.

**`internal/install/`** - Post-installation tasks including running package managers, TypeScript type generation, and git initialization.

**`internal/util/`** - Shared utilities for colors, command execution, and filesystem operations.

**`templates/`** - Embedded template files using Go's `embed` directive. Identical to TypeScript version.

---

## Required Dependencies

### Essential Packages

```go
// go.mod
module github.com/yourusername/create-next-app-go

go 1.21

require (
    // CLI Framework
    github.com/spf13/cobra v1.8.0
    github.com/spf13/viper v1.18.0

    // Charm Packages (UI/UX)
    github.com/charmbracelet/huh v0.6.0
    github.com/charmbracelet/lipgloss v1.0.0

    // File Operations
    github.com/bmatcuk/doublestar/v4 v4.6.0

    // Optional: Git Operations (can use os/exec instead)
    github.com/go-git/go-git/v5 v5.11.0
)
```

### Package Purpose Mapping

| Package        | Purpose                   | Replaces (TypeScript) |
| -------------- | ------------------------- | --------------------- |
| **Cobra**      | CLI argument parsing      | `commander`           |
| **Viper**      | Configuration/preferences | `conf`                |
| **Huh**        | Interactive prompts       | `prompts`             |
| **Lip Gloss**  | Terminal colors/styling   | `picocolors`          |
| **Doublestar** | File globbing             | `fast-glob`           |
| **Go-git**     | Git operations (optional) | `child_process`       |

### What You DON'T Need

- **Bubble Tea** - Overkill for simple prompts; Huh handles all needs
- **Bubbles** - Only needed if you want fancy spinners (optional)
- **Glamour** - No markdown rendering needed
- **Gum** - Shell script tool, not needed in Go
- **Wish** - SSH server, not needed

---

## Core Data Structures

### Main Configuration Struct

```go
// internal/config/config.go
package config

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

### User Preferences Struct

```go
// internal/config/preferences.go
package config

type Preferences struct {
    TypeScript         bool   `json:"typescript"`
    Linter            string `json:"linter"` // "eslint", "biome", "none"
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

### GitHub Repository Info

```go
// internal/example/github.go
package example

type RepoInfo struct {
    Username string
    Repo     string
    Branch   string
    FilePath string // For subdirectories
}
```

---

## Module Implementation Guide

### 1. Main Entry Point

**File:** `main.go`

**Responsibilities:**

- Set up signal handlers (SIGINT, SIGTERM)
- Initialize Cobra root command
- Execute CLI and handle top-level errors

**Key Implementation:**

```go
package main

import (
    "os"
    "os/signal"
    "syscall"
    "github.com/yourusername/create-next-app-go/cmd"
)

func main() {
    // Setup signal handlers for graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-sigChan
        // Restore cursor and exit gracefully
        os.Exit(0)
    }()

    // Execute CLI
    if err := cmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

---

### 2. CLI Command Definition

**File:** `cmd/root.go`

**Responsibilities:**

- Define root command and all flags
- Map flags to Config struct
- Detect package manager from environment
- Call main application logic

**Flag Categories:**

1. **Project Configuration**
   - `--typescript` / `--javascript`
   - `--tailwind` / `--no-tailwind`
   - `--app` / `--pages`
   - `--src-dir` / `--no-src-dir`
   - `--import-alias <string>`

2. **Linting & Tools**
   - `--eslint` / `--biome` / `--no-lint`
   - `--react-compiler` / `--no-react-compiler`

3. **Bundler Options**
   - `--turbo` / `--webpack` / `--rspack`

4. **Package Manager**
   - `--use-npm` / `--use-pnpm` / `--use-yarn` / `--use-bun`
   - `--skip-install`

5. **Git Options**
   - `--skip-git`

6. **Example Mode**
   - `--example <name-or-url>`
   - `--example-path <path>`

7. **Automation & Preferences**
   - `--yes` (skip prompts, use defaults)
   - `--reset-preferences`
   - `--empty` (minimal template)

**Key Implementation Pattern:**

```go
package cmd

import (
    "github.com/spf13/cobra"
    "github.com/yourusername/create-next-app-go/internal/config"
)

var rootCmd = &cobra.Command{
    Use:   "create-next-app [directory]",
    Short: "Create a new Next.js application",
    Args:  cobra.MaximumNArgs(1),
    RunE:  runCreate,
}

func init() {
    // Boolean flags
    rootCmd.Flags().Bool("typescript", false, "Initialize as TypeScript project")
    rootCmd.Flags().Bool("javascript", false, "Initialize as JavaScript project")
    rootCmd.Flags().Bool("tailwind", false, "Initialize with Tailwind CSS")
    rootCmd.Flags().Bool("eslint", false, "Initialize with ESLint")
    // ... add all other flags
}

func runCreate(cmd *cobra.Command, args []string) error {
    cfg := config.New()

    // Parse flags into config
    // ... flag parsing logic

    // Run main application
    return run(cfg)
}
```

---

### 3. Configuration Module

**Files:** `internal/config/config.go`, `preferences.go`, `defaults.go`

**Responsibilities:**

- Define Config struct
- Load/save user preferences using Viper
- Provide default values
- Merge CLI flags, preferences, and defaults

**Preference Storage Location:**

- Unix/Linux: `~/.config/create-next-app/preferences.json`
- macOS: `~/Library/Application Support/create-next-app/preferences.json`
- Windows: `%APPDATA%\create-next-app\preferences.json`

**Key Functions:**

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

// Merge config sources (CLI > Preferences > Defaults)
func MergeConfig(flags *Config, prefs *Preferences) *Config
```

---

### 4. Interactive Prompts Module

**Files:** `internal/prompt/prompt.go`, `setup.go`, `validate.go`

**Responsibilities:**

- Show interactive prompts using Huh
- Handle initial setup choice (recommended/reuse/customize)
- Collect missing configuration
- Validate user input

**Prompt Types Needed:**

1. **Text Input** - Project name with validation
2. **Confirm** - Yes/No toggles (TypeScript, Tailwind, etc.)
3. **Select** - Single choice from list (Linter selection)

**Implementation Pattern:**

```go
package prompt

import "github.com/charmbracelet/huh"

// Ask for project name
func AskProjectName(defaultName string) (string, error) {
    var name string

    err := huh.NewInput().
        Title("What is your project named?").
        Value(&name).
        Placeholder(defaultName).
        Validate(validateProjectName).
        Run()

    if name == "" {
        name = defaultName
    }

    return name, err
}

// Ask for TypeScript preference
func AskTypeScript() (bool, error) {
    var useTS bool

    err := huh.NewConfirm().
        Title("Would you like to use TypeScript?").
        Value(&useTS).
        Run()

    return useTS, err
}

// Ask for linter choice
func AskLinter() (string, error) {
    var linter string

    err := huh.NewSelect[string]().
        Title("Which linter would you like to use?").
        Options(
            huh.NewOption("ESLint", "eslint"),
            huh.NewOption("Biome", "biome"),
            huh.NewOption("None", "none"),
        ).
        Value(&linter).
        Run()

    return linter, err
}

// Group multiple prompts together
func AskConfigOptions() (*config.Config, error) {
    cfg := &config.Config{}

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewConfirm().
                Title("Would you like to use TypeScript?").
                Value(&cfg.TypeScript),
            huh.NewConfirm().
                Title("Would you like to use Tailwind CSS?").
                Value(&cfg.Tailwind),
            // ... more prompts
        ),
    )

    return cfg, form.Run()
}
```

**Initial Setup Choice:**

```go
// Ask if user wants recommended defaults, reuse settings, or customize
func AskSetupChoice(hasSavedPrefs bool) (string, error) {
    var choice string

    options := []huh.Option[string]{
        huh.NewOption("Yes, use recommended defaults", "recommended"),
    }

    if hasSavedPrefs {
        options = append(options,
            huh.NewOption("No, reuse previous settings", "reuse"))
    }

    options = append(options,
        huh.NewOption("No, customize settings", "customize"))

    err := huh.NewSelect[string]().
        Title("Would you like to use the recommended Next.js defaults?").
        Options(options...).
        Value(&choice).
        Run()

    return choice, err
}
```

---

### 5. Validation Module

**Files:** `internal/validate/name.go`, `directory.go`, `network.go`

**Responsibilities:**

- Validate npm package names
- Check directory permissions and emptiness
- Detect CI environment
- Check network connectivity

**NPM Name Validation:**

Must follow npm naming rules:

- Length: 1-214 characters
- Cannot start with `.` or `_`
- No uppercase letters
- No non-URL-safe characters
- Cannot be npm built-in modules

**Directory Validation:**

Check if directory:

- Exists and is empty (or contains only allowed files)
- Has write permissions
- Allowed files: `.git`, `.gitignore`, `LICENSE`, `README.md`

**Implementation Helpers:**

```go
// Validate npm package name
func ValidateNpmName(name string) error

// Check if directory is safe to use
func ValidateDirectory(path string) error

// Check if directory is empty (allowing certain files)
func IsFolderEmpty(path string) (bool, []string, error)

// Detect if running in CI environment
func IsCI() bool

// Check network connectivity
func IsOnline() bool
```

---

### 6. Template Installation Module

**Files:** `internal/template/install.go`, `copy.go`, `transform.go`, `packagejson.go`

**Responsibilities:**

- Copy built-in templates from embedded filesystem
- Apply file transformations (renames, content modifications)
- Handle src directory restructuring
- Replace import aliases
- Generate package.json

**Template Selection Logic:**

Nine template types exist:

- `app` - App Router
- `app-tw` - App Router + Tailwind
- `app-empty` - App Router (minimal)
- `app-tw-empty` - App Router + Tailwind (minimal)
- `app-api` - API-only (no React)
- `default` - Pages Router
- `default-tw` - Pages Router + Tailwind
- `default-empty` - Pages Router (minimal)
- `default-tw-empty` - Pages Router + Tailwind (minimal)

Each has both JS and TS variants.

**File Transformations:**

1. **Renames During Copy:**
   - `gitignore` → `.gitignore`
   - `README-template.md` → `README.md`

2. **Config Modifications:**
   - For Rspack: Add import and wrap export
   - For React Compiler: Add `reactCompiler: true`

3. **Path Config Updates:**
   - For src directory: Update tsconfig/jsconfig paths
   - For custom alias: Update import alias key

4. **Import Alias Replacement:**
   - Replace `@/` with custom alias in all source files
   - Exclude: configs, .git, fonts, favicon

5. **Src Directory Restructuring:**
   - Create `src/` folder
   - Move `app/`, `pages/`, `styles/` into `src/`
   - Update page references to include `src/` prefix

**Embedding Templates:**

```go
package template

import "embed"

//go:embed templates/*
var templatesFS embed.FS

func CopyTemplate(templateName string, targetPath string) error {
    // Walk embedded filesystem
    // Copy files with transformations
    // Apply renames
}
```

---

### 7. Example Download Module

**Files:** `internal/example/download.go`, `extract.go`, `github.go`

**Responsibilities:**

- Parse GitHub URLs or example names
- Fetch from GitHub API
- Download tar archives
- Extract to target directory
- Retry on network failures

**GitHub URL Formats Supported:**

- `https://github.com/user/repo`
- `https://github.com/user/repo/tree/branch`
- `https://github.com/user/repo/tree/branch/path/to/dir`

**Download Flow:**

1. Parse URL or look up example name in official examples list
2. Fetch repo info from GitHub API
3. Check if path exists (for subdirectories)
4. Download tar archive
5. Extract to temporary directory
6. Copy relevant subdirectory to target
7. Retry up to 3 times on failure

**Error Handling:**

- Invalid URL: Show error and exit
- Repo not found: Show error and exit
- Download failure: Retry 3 times, then offer to use default template
- Network issues: Detect and prompt for offline mode

---

### 8. Package Installation Module

**Files:** `internal/install/deps.go`, `typegen.go`, `git.go`

**Responsibilities:**

- Run package manager install command
- Generate TypeScript types (if TypeScript enabled)
- Initialize git repository
- Display progress/errors

**Package Manager Commands:**

- npm: `npm install`
- pnpm: `pnpm install`
- yarn: `yarn install`
- bun: `bun install`

**TypeScript Type Generation:**

Run after installation for TypeScript projects:

```bash
npm run next telemetry disable
npm run next typegen
```

This is best-effort; failures are logged but not fatal.

**Git Initialization:**

```bash
git init
git add -A
git commit -m "Initial commit from Create Next App"
```

Silent failures are acceptable; git init is optional.

---

### 9. Terminal Styling Module

**File:** `internal/util/colors.go`

**Responsibilities:**

- Define color styles using Lip Gloss
- Provide helper functions for styled output
- Match TypeScript CLI aesthetics

**Common Styles Needed:**

```go
package util

import "github.com/charmbracelet/lipgloss"

var (
    Green  = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
    Blue   = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
    Cyan   = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
    Yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
    Red    = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
    Bold   = lipgloss.NewStyle().Bold(true)
)

// Helper functions
func Success(s string) string { return Green.Render(s) }
func Info(s string) string    { return Cyan.Render(s) }
func Warning(s string) string { return Yellow.Render(s) }
func Error(s string) string   { return Red.Render(s) }
```

---

## Template System

### Template Organization

Templates are organized by router type and configuration:

```
templates/
├── app/                    # App Router (JavaScript)
├── app-ts/                 # App Router (TypeScript)
├── app-tw/                 # App Router + Tailwind (JavaScript)
├── app-tw-ts/              # App Router + Tailwind (TypeScript)
├── app-empty/              # App Router minimal (JavaScript)
├── app-empty-ts/           # App Router minimal (TypeScript)
├── default/                # Pages Router (JavaScript)
├── default-ts/             # Pages Router (TypeScript)
├── default-tw/             # Pages Router + Tailwind (JavaScript)
└── default-tw-ts/          # Pages Router + Tailwind (TypeScript)
```

### Template Selection Algorithm

```
if config.APIOnly:
    template = "app-api"
elif config.AppRouter:
    if config.EmptyTemplate:
        template = "app-empty" + ("-tw" if config.Tailwind else "")
    else:
        template = "app" + ("-tw" if config.Tailwind else "")
else:
    if config.EmptyTemplate:
        template = "default-empty" + ("-tw" if config.Tailwind else "")
    else:
        template = "default" + ("-tw" if config.Tailwind else "")

if config.TypeScript:
    template += "-ts"
```

### Package.json Generation

The package.json is generated dynamically based on configuration:

**Base Structure:**

```json
{
  "name": "project-name",
  "version": "0.1.0",
  "private": true
}
```

**Scripts (vary by config):**

```json
{
  "scripts": {
    "dev": "next dev --turbo", // or --webpack
    "build": "next build",
    "start": "next start",
    "lint": "eslint ." // if ESLint
  }
}
```

**Dependencies:**

- Always: `next`
- If not API-only: `react`, `react-dom`
- If Rspack: `next-rspack`

**DevDependencies:**

- If TypeScript: `typescript`, `@types/node`, `@types/react`, `@types/react-dom`
- If Tailwind: `tailwindcss`, `postcss`, `autoprefixer`
- If ESLint: `eslint`, `eslint-config-next`
- If Biome: `@biomejs/biome`
- If React Compiler: `babel-plugin-react-compiler`

---

## Application Flow

### Complete Flow Diagram

```
1. Parse CLI Args (Cobra)
   ├─> Detect package manager
   ├─> Parse all flags
   └─> Get project path argument

2. Load Preferences (Viper)
   ├─> Check if preferences exist
   └─> Load from config directory

3. Handle --reset-preferences
   └─> If set: Clear prefs, exit

4. Validate Project Path
   ├─> Resolve to absolute path
   ├─> Validate npm name
   └─> Check directory safety

5. Determine Prompt Mode
   ├─> Skip if --yes flag
   ├─> Skip if CI environment
   └─> Otherwise: interactive mode

6. Interactive Configuration (if needed)
   ├─> Ask setup choice (recommended/reuse/customize)
   ├─> For each option:
   │   ├─> Use CLI flag if provided
   │   ├─> Use preference if available
   │   └─> Prompt if neither
   └─> Save new preferences

7. Create Project Directory
   └─> mkdir -p {projectPath}

8. Branch: Example vs Template

   IF --example PROVIDED:
   ├─> Parse GitHub URL or example name
   ├─> Download from GitHub API
   ├─> Extract tar to project directory
   └─> Skip to step 10

   ELSE:
   └─> Continue to step 9

9. Install Template
   ├─> Select template based on config
   ├─> Copy from embedded filesystem
   ├─> Apply file transformations
   ├─> Apply import alias replacements
   ├─> Restructure for src directory (if enabled)
   └─> Generate package.json

10. Post-Installation Tasks
    ├─> Run package manager install (unless --skip-install)
    ├─> Run TypeScript typegen (if TypeScript)
    └─> Initialize git repo (unless --skip-git)

11. Display Success Message
    └─> Show "cd {projectName}" and "npm run dev" instructions
```

### Detailed Step Breakdown

**Step 1-2: Initialization**

- Set up signal handlers
- Parse CLI arguments
- Detect package manager from `npm_config_user_agent`
- Load user preferences from config directory

**Step 3: Reset Preferences (Optional)**

- If `--reset-preferences` flag is set
- Prompt for confirmation
- Clear preferences file
- Exit successfully

**Step 4: Path Validation**

- Get project path from CLI arg or prompt
- Resolve to absolute path
- Extract project name from basename
- Validate against npm naming rules
- Check if directory exists and is empty
- Verify write permissions

**Step 5: Determine Interaction Mode**

- Skip prompts if `--yes` flag
- Skip prompts if CI environment detected
- Otherwise, enter interactive mode

**Step 6: Configuration Collection**

- Show initial setup choice if no CLI flags
- For each configuration option:
  - Check CLI flag first (highest priority)
  - Check user preference second
  - Prompt user third
  - Use default last
- Save user responses to preferences

**Step 7: Project Directory**

- Create project directory
- Handle permission errors

**Step 8: Template vs Example Branch**

- If `--example` provided: Download from GitHub
- Otherwise: Install from built-in templates

**Step 9: Template Installation**

- Select appropriate template
- Copy files from embedded filesystem
- Rename special files (gitignore, README)
- Modify configs (Rspack, React Compiler)
- Replace import aliases if custom
- Restructure for src directory if enabled
- Generate package.json dynamically

**Step 10: Post-Installation**

- Run package manager install
- Run TypeScript type generation (best effort)
- Initialize git repository (silent failure OK)

**Step 11: Success Output**

- Display success message with project path
- Show next steps (cd, npm run dev)

---

## Error Handling

### Error Types

Define custom error types for better error messages:

```go
type ValidationError struct {
    Field    string
    Problems []string
}

type DirectoryError struct {
    Path           string
    ConflictingFiles []string
}

type DownloadError struct {
    URL string
    Err error
}

type InstallError struct {
    Command string
    Err     error
}
```

### Error Scenarios and Responses

**Invalid Project Name:**

```
Could not create project because of npm naming restrictions:
  • Name cannot contain capital letters
  • Name length must be less than 214 characters
```

**Directory Not Empty:**

```
The directory my-app contains files that could conflict:
  • index.html
  • package.json

Either use a new directory name, or remove the files listed above.
```

**No Write Permission:**

```
The application path is not writable, please check folder permissions.
It is likely you do not have write permissions for this folder.
```

**Example Not Found:**

```
Could not locate an example named "next-auth".
It could be due to:
  1. Your spelling might be incorrect.
  2. You might not be connected to the internet or behind a proxy.
```

**Repository Not Found:**

```
Could not locate the repository for "https://github.com/user/repo".
Please check that it exists and try again.
```

**Download Failure with Retry:**

```
Retrying download... (Attempt 2 of 3)
...
Could not download because of a connectivity issue.
Would you like to use the default template instead? (Y/n)
```

**Installation Failure:**

```
Aborting installation.
npm install has failed.
```

### Signal Handling

**SIGINT/SIGTERM Handler:**

- Restore terminal cursor (may be hidden by prompts)
- Print newline for clean exit
- Exit with code 0

**Prompt Cancellation (Ctrl+C):**

- Detect `huh.ErrUserAborted`
- Print "Exiting."
- Exit with code 1

---

## Building and Distribution

### Build Commands

**Build for current platform:**

```bash
go build -o create-next-app
```

**Build for multiple platforms:**

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o create-next-app-linux

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o create-next-app-macos

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o create-next-app-macos-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o create-next-app.exe
```

**Build with optimizations:**

```bash
go build -ldflags="-s -w" -o create-next-app
```

- `-s`: Omit symbol table
- `-w`: Omit DWARF debugging info
- Results in smaller binary size

### Embedding Templates

Templates are embedded at compile time using Go's `embed` directive:

```go
package template

import "embed"

//go:embed templates/*
var templatesFS embed.FS
```

This allows distribution as a single binary with no external dependencies.

### Distribution

**Advantages of Go Binary:**

- Single file, no dependencies
- No Node.js required to run the CLI
- Fast startup (no interpreter)
- Cross-platform support

**Installation Methods:**

1. **Direct download:**

   ```bash
   curl -L https://example.com/create-next-app -o /usr/local/bin/create-next-app
   chmod +x /usr/local/bin/create-next-app
   ```

2. **Go install:**

   ```bash
   go install github.com/yourusername/create-next-app-go@latest
   ```

3. **Package managers:**
   - Homebrew (macOS)
   - Scoop (Windows)
   - APT/YUM (Linux)

---

## Next Steps

### Implementation Order

1. **Set up project structure** - Create folders, initialize go.mod
2. **Implement Config module** - Define structs, defaults, preferences
3. **Implement CLI commands** - Set up Cobra, define flags
4. **Implement Validation** - NPM name, directory checks
5. **Implement Prompts** - Huh-based interactive prompts
6. **Implement Template system** - Embed templates, copy logic
7. **Implement Package.json generation** - Dynamic generation
8. **Implement Transformations** - Import alias, src directory
9. **Implement Example download** - GitHub API, tar extraction
10. **Implement Post-install** - Package manager, git init
11. **Add error handling** - Custom errors, user-friendly messages
12. **Add styling** - Lip Gloss colors, success messages
13. **Testing** - Unit tests, integration tests
14. **Documentation** - README, usage examples

### Testing Strategy

**Unit Tests:**

- Validation functions
- Configuration merging
- Template selection logic
- Package.json generation

**Integration Tests:**

- End-to-end CLI execution
- Template installation
- File transformations

**Manual Testing:**

- Test all flag combinations
- Test interactive prompts
- Test on multiple platforms

---

## Summary

This guide provides a complete blueprint for building the create-next-app CLI in Go. Key takeaways:

- **Architecture**: Linear pipeline with template/example branching
- **Packages**: Cobra for CLI, Huh for prompts, Lip Gloss for styling
- **Structure**: Clean separation of concerns with internal packages
- **Templates**: Embedded at compile time for single-binary distribution
- **Flow**: Parse → Configure → Validate → Install/Download → Post-install → Success

The Go version will be faster, easier to distribute, and ready for future backend feature additions while maintaining complete feature parity with the TypeScript implementation.
