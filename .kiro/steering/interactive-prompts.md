---
inclusion: always
---

# Interactive Prompts

## Prompt Library: Huh

Use Charmbracelet's Huh for all interactive prompts:

```go
import "github.com/charmbracelet/huh"
```

## Prompt Types

### Text Input

For project name and custom values:

```go
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
```

### Confirm (Yes/No)

For boolean toggles:

```go
func AskTypeScript() (bool, error) {
    var useTS bool

    err := huh.NewConfirm().
        Title("Would you like to use TypeScript?").
        Value(&useTS).
        Affirmative("Yes").
        Negative("No").
        Run()

    return useTS, err
}
```

### Select (Single Choice)

For options with multiple choices:

```go
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
```

## Grouped Prompts

Group related prompts together for better UX:

```go
func AskConfigOptions(cfg *Config) error {
    form := huh.NewForm(
        huh.NewGroup(
            huh.NewConfirm().
                Title("Would you like to use TypeScript?").
                Value(&cfg.TypeScript),

            huh.NewConfirm().
                Title("Would you like to use Tailwind CSS?").
                Value(&cfg.Tailwind),

    huh.NewSelect[string]().
                Title("Which linter would you like to use?").
                Options(
                    huh.NewOption("ESLint", "eslint"),
                    huh.NewOption("Biome", "biome"),
                    huh.NewOption("None", "none"),
                ).
                Value(&cfg.Linter),
        ),
    )

    return form.Run()
}
```

## Initial Setup Choice

Ask user how they want to configure the project:

```go
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

## Prompt Flow Logic

```go
func CollectConfiguration(cfg *Config, prefs *Preferences, hasFlags bool) error {
    // Skip prompts if --yes flag or CI environment
    if cfg.SkipPrompts || IsCI() {
        return nil
    }

    // Ask initial setup choice
    choice, err := AskSetupChoice(HasPreferences())
    if err != nil {
        return err
    }

    switch choice {
    case "recommended":
        // Use defaults, save as preferences
        SavePreferences(DefaultPreferences())
        return nil

    case "reuse":
        // Load and apply saved preferences
        return nil

    case "customize":
        // Ask each question individually
        return AskAllOptions(cfg)
    }

    return nil
}
```

## Conditional Prompts

Only ask for options not provided via CLI flags:

```go
func AskAllOptions(cfg *Config) error {
    // Only ask if not set via flag
    if cfg.TypeScript == nil {
        ts, err := AskTypeScript()
        if err != nil {
            return err
        }
        cfg.TypeScript = ts
    }

    if cfg.Tailwind == nil {
        tw, err := AskTailwind()
        if err != nil {
            return err
        }
        cfg.Tailwind = tw
    }

    // ... continue for all options

    return nil
}
```

## Error Handling

Handle user cancellation (Ctrl+C):

```go
func RunPrompts() error {
    err := AskProjectName()

    if errors.Is(err, huh.ErrUserAborted) {
        fmt.Println("\nExiting.")
        os.Exit(1)
    }

    if err != nil {
        return fmt.Errorf("prompt failed: %w", err)
    }

    return nil
}
```

## Validation

Add real-time validation to inputs:

```go
func validateProjectName(name string) error {
    if name == "" {
        return fmt.Errorf("project name cannot be empty")
    }

    if err := ValidateNpmName(name); err != nil {
        return fmt.Errorf("invalid project name: %w", err)
    }

    return nil
}

func AskProjectName() (string, error) {
    var name string

    err := huh.NewInput().
        Title("What is your project named?").
        Value(&name).
        Validate(validateProjectName). // Real-time validation
        Run()

    return name, err
}
```

## Prompt Skipping Rules

Skip prompts in these scenarios:

1. **--yes flag provided** - Use all defaults
2. **CI environment detected** - Use all defaults
3. **All options provided via flags** - No need to prompt
4. **Example mode (--example)** - Skip template configuration

```go
func ShouldSkipPrompts(cfg *Config) bool {
    return cfg.SkipPrompts ||
           IsCI() ||
           cfg.Example != "" ||
           AllOptionsProvided(cfg)
}
```

```

```
