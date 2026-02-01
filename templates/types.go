package templates

// TemplateType represents the type of Next.js template to use
// Note: Only App Router templates are supported (no Pages Router)
type TemplateType string

const (
	// App is the standard App Router template
	App TemplateType = "app"
	// AppAPI is the API-only template (no React components)
	AppAPI TemplateType = "app-api"
	// AppEmpty is the minimal App Router template
	AppEmpty TemplateType = "app-empty"
	// AppTW is the App Router template with Tailwind CSS
	AppTW TemplateType = "app-tw"
	// AppTWEmpty is the minimal App Router template with Tailwind CSS
	AppTWEmpty TemplateType = "app-tw-empty"
)

// TemplateMode represents the language mode (JavaScript or TypeScript)
type TemplateMode string

const (
	// JS represents JavaScript mode
	JS TemplateMode = "js"
	// TS represents TypeScript mode
	TS TemplateMode = "ts"
)

// Bundler represents the bundler choice for Next.js
type Bundler string

const (
	// Turbopack is the default Next.js bundler
	Turbopack Bundler = "turbopack"
	// Webpack is the traditional Next.js bundler
	Webpack Bundler = "webpack"
	// Rspack is the Rust-based bundler
	Rspack Bundler = "rspack"
)

// PackageManager represents the package manager to use
type PackageManager string

const (
	// NPM is the Node Package Manager
	NPM PackageManager = "npm"
	// PNPM is the performant npm
	PNPM PackageManager = "pnpm"
	// Yarn is the Yarn package manager
	Yarn PackageManager = "yarn"
	// Bun is the Bun package manager
	Bun PackageManager = "bun"
)

// GetTemplateFileArgs contains arguments for getting a template file path
type GetTemplateFileArgs struct {
	Template TemplateType
	Mode     TemplateMode
	File     string
}

// InstallTemplateArgs contains all arguments needed to install a template
type InstallTemplateArgs struct {
	AppName        string
	Root           string
	PackageManager PackageManager
	IsOnline       bool
	Template       TemplateType
	Mode           TemplateMode
	Eslint         bool
	Biome          bool
	Tailwind       bool
	SrcDir         bool
	ImportAlias    string
	SkipInstall    bool
	Bundler        Bundler
	ReactCompiler  bool
}
