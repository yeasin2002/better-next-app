package prompt

import (
	"github.com/charmbracelet/huh"
	"github.com/yeasin2002/better-next-app/internal/config"
)

// AskProjectName prompts for project name
func AskProjectName(defaultName string) (string, error) {
	var name string

	err := huh.NewInput().
		Title("What is your project named?").
		Value(&name).
		Placeholder(defaultName).
		Validate(ValidateProjectName).
		Run()

	if name == "" {
		name = defaultName
	}

	return name, err
}

// AskTypeScript prompts for TypeScript preference
func AskTypeScript() (bool, error) {
	var useTS bool

	err := huh.NewConfirm().
		Title("Would you like to use TypeScript?").
		Value(&useTS).
		Run()

	return useTS, err
}

// AskTailwind prompts for Tailwind CSS preference
func AskTailwind() (bool, error) {
	var useTailwind bool

	err := huh.NewConfirm().
		Title("Would you like to use Tailwind CSS?").
		Value(&useTailwind).
		Run()

	return useTailwind, err
}

// AskLinter prompts for linter choice
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

// AskSrcDir prompts for src directory preference
func AskSrcDir() (bool, error) {
	var useSrcDir bool

	err := huh.NewConfirm().
		Title("Would you like to use `src/` directory?").
		Value(&useSrcDir).
		Run()

	return useSrcDir, err
}

// AskAppRouter is deprecated - App Router is now always enabled
// Pages Router templates have been removed
func AskAppRouter() (bool, error) {
	// Always return true - App Router is the only option
	return true, nil
}

// AskImportAlias prompts for custom import alias
func AskImportAlias() (string, error) {
	var alias string

	err := huh.NewInput().
		Title("What import alias would you like configured?").
		Value(&alias).
		Placeholder("@/*").
		Run()

	if alias == "" {
		alias = "@/*"
	}

	return alias, err
}

// AskConfigOptions collects all configuration options in a form
func AskConfigOptions() (*config.Config, error) {
	cfg := config.New()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Would you like to use TypeScript?").
				Value(&cfg.TypeScript),
			huh.NewConfirm().
				Title("Would you like to use Tailwind CSS?").
				Value(&cfg.Tailwind),
			huh.NewConfirm().
				Title("Would you like to use `src/` directory?").
				Value(&cfg.SrcDir),
		),
	)

	// App Router is always enabled
	cfg.AppRouter = true

	err := form.Run()
	return cfg, err
}
