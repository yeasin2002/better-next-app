package prompt

import (
	"github.com/charmbracelet/huh"
)

const (
	SetupRecommended = "recommended"
	SetupReuse       = "reuse"
	SetupCustomize   = "customize"
)

// AskSetupChoice prompts for initial setup choice
func AskSetupChoice(hasSavedPrefs bool) (string, error) {
	var choice string

	options := []huh.Option[string]{
		huh.NewOption("Yes, use recommended defaults", SetupRecommended),
	}

	if hasSavedPrefs {
		options = append(options,
			huh.NewOption("No, reuse previous settings", SetupReuse))
	}

	options = append(options,
		huh.NewOption("No, customize settings", SetupCustomize))

	err := huh.NewSelect[string]().
		Title("Would you like to use the recommended Next.js defaults?").
		Options(options...).
		Value(&choice).
		Run()

	return choice, err
}
