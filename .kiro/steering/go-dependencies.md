---
inclusion: always
---

# Go Dependencies & Package Selection

## Required Packages

```go
require (
    // CLI Framework
    github.com/spf13/cobra v1.8.0
    github.com/spf13/viper v1.18.0

    // Charm Packages (UI/UX)
    github.com/charmbracelet/huh v0.6.0
    github.com/charmbracelet/lipgloss v1.0.0

    // File Operations
    github.com/bmatcuk/doublestar/v4 v4.6.0

    // Optional: Git Operations
    github.com/go-git/go-git/v5 v5.11.0
)
```

## Package Purpose Mapping

| Package        | Purpose                   | Replaces (TypeScript) |
| -------------- | ------------------------- | --------------------- |
| **Cobra**      | CLI argument parsing      | `commander`           |
| **Viper**      | Configuration/preferences | `conf`                |
| **Huh**        | Interactive prompts       | `prompts`             |
| **Lip Gloss**  | Terminal colors/styling   | `picocolors`          |
| **Doublestar** | File globbing             | `fast-glob`           |
| **Go-git**     | Git operations (optional) | `child_process`       |

## What NOT to Use

- **Bubble Tea** - Overkill for simple prompts; Huh handles all needs
- **Bubbles** - Only needed for fancy spinners (optional)
- **Glamour** - No markdown rendering needed
- **Gum** - Shell script tool, not needed in Go
- **Wish** - SSH server, not needed

## Cobra Usage

Use Cobra for all CLI flag definitions and command structure:

```go
var rootCmd = &cobra.Command{
    Use:   "create-next-app [directory]",
    Short: "Create a new Next.js application",
    Args:  cobra.MaximumNArgs(1),
    RunE:  runCreate,
}
```

## Viper Usage

Use Viper for preference persistence:

- Unix/Linux: `~/.config/create-next-app/preferences.json`
- macOS: `~/Library/Application Support/create-next-app/preferences.json`
- Windows: `%APPDATA%\create-next-app\preferences.json`

## Huh Usage

Use Huh for all interactive prompts:

- `huh.NewInput()` - Text input with validation
- `huh.NewConfirm()` - Yes/No toggles
- `huh.NewSelect()` - Single choice from list
- `huh.NewForm()` - Group multiple prompts together

## Lipgloss Usage

Define color styles for consistent terminal output:

```go
var (
    Green  = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
    Blue   = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
    Cyan   = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
    Yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
    Red    = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
    Bold   = lipgloss.NewStyle().Bold(true)
)
```
