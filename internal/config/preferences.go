package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Preferences stores user's saved preferences
type Preferences struct {
	TypeScript     bool   `json:"typescript" mapstructure:"typescript"`
	Linter         string `json:"linter" mapstructure:"linter"`
	Tailwind       bool   `json:"tailwind" mapstructure:"tailwind"`
	AppRouter      bool   `json:"appRouter" mapstructure:"appRouter"`
	SrcDir         bool   `json:"srcDir" mapstructure:"srcDir"`
	ImportAlias    string `json:"importAlias" mapstructure:"importAlias"`
	CustomizeAlias bool   `json:"customizeAlias" mapstructure:"customizeAlias"`
	EmptyTemplate  bool   `json:"emptyTemplate" mapstructure:"emptyTemplate"`
	DisableGit     bool   `json:"disableGit" mapstructure:"disableGit"`
	ReactCompiler  bool   `json:"reactCompiler" mapstructure:"reactCompiler"`
}

// getConfigDir returns the config directory path
func getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "create-next-app"), nil
}

// LoadPreferences loads user preferences from disk
func LoadPreferences() (*Preferences, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	viper.SetConfigName("preferences")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, nil
		}
		return nil, err
	}

	var prefs Preferences
	if err := viper.Unmarshal(&prefs); err != nil {
		return nil, err
	}

	return &prefs, nil
}

// SavePreferences saves user preferences to disk
func SavePreferences(prefs *Preferences) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	viper.Set("typescript", prefs.TypeScript)
	viper.Set("linter", prefs.Linter)
	viper.Set("tailwind", prefs.Tailwind)
	viper.Set("appRouter", prefs.AppRouter)
	viper.Set("srcDir", prefs.SrcDir)
	viper.Set("importAlias", prefs.ImportAlias)
	viper.Set("customizeAlias", prefs.CustomizeAlias)
	viper.Set("emptyTemplate", prefs.EmptyTemplate)
	viper.Set("disableGit", prefs.DisableGit)
	viper.Set("reactCompiler", prefs.ReactCompiler)

	return viper.WriteConfigAs(filepath.Join(configDir, "preferences.json"))
}

// HasPreferences checks if preferences file exists
func HasPreferences() bool {
	configDir, err := getConfigDir()
	if err != nil {
		return false
	}

	_, err = os.Stat(filepath.Join(configDir, "preferences.json"))
	return err == nil
}

// ClearPreferences removes the preferences file
func ClearPreferences() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	return os.Remove(filepath.Join(configDir, "preferences.json"))
}

// MergeConfig merges CLI flags with preferences and defaults
func MergeConfig(flags *Config, prefs *Preferences) *Config {
	defaults := DefaultConfig()

	if flags == nil {
		flags = &Config{}
	}

	// Apply preferences over defaults, then flags over preferences
	result := defaults

	if prefs != nil {
		result.TypeScript = prefs.TypeScript
		result.Linter = prefs.Linter
		result.Tailwind = prefs.Tailwind
		result.AppRouter = prefs.AppRouter
		result.SrcDir = prefs.SrcDir
		result.ImportAlias = prefs.ImportAlias
		result.EmptyTemplate = prefs.EmptyTemplate
		result.SkipGit = prefs.DisableGit
		result.ReactCompiler = prefs.ReactCompiler
	}

	// CLI flags take highest priority
	if flags.ProjectName != "" {
		result.ProjectName = flags.ProjectName
	}
	if flags.ProjectPath != "" {
		result.ProjectPath = flags.ProjectPath
	}

	return result
}
