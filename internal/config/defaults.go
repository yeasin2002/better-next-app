package config

// DefaultConfig returns configuration with recommended defaults
func DefaultConfig() *Config {
	return &Config{
		TypeScript:     true,
		AppRouter:      true,
		Tailwind:       true,
		Linter:         "eslint",
		SrcDir:         false,
		ImportAlias:    "@/*",
		EmptyTemplate:  false,
		Bundler:        "turbopack",
		ReactCompiler:  false,
		PackageManager: "npm",
		SkipInstall:    false,
		SkipGit:        false,
	}
}
