package templates

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// NextjsReactPeerVersion is the React version used by Next.js
// Do not rename or format. sync-react script relies on this line.
const NextjsReactPeerVersion = "19.2.3"

// SrcDirNames are the directories that should be moved into src/ when srcDir is enabled
var SrcDirNames = []string{"app", "pages", "styles"}

//go:embed app app-tw app-empty app-tw-empty app-api
var templatesFS embed.FS

// GetTemplateFile returns the file path for a given file in a template
func GetTemplateFile(args GetTemplateFileArgs) string {
	return filepath.Join(string(args.Template), string(args.Mode), args.File)
}

// sorted returns a map with keys sorted alphabetically
func sorted(m map[string]string) map[string]string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := make(map[string]string, len(m))
	for _, k := range keys {
		result[k] = m[k]
	}
	return result
}

// InstallTemplate installs a Next.js template to the given root directory
func InstallTemplate(args InstallTemplateArgs) error {
	fmt.Printf("\nInitializing project with template: %s\n\n", args.Template)

	isAPI := args.Template == AppAPI
	templatePath := filepath.Join(string(args.Template), string(args.Mode))

	// Determine which files to copy
	copySource := []string{"**"}
	if !args.Eslint {
		copySource = append(copySource, "!eslint.config.mjs")
	}
	if !args.Biome {
		copySource = append(copySource, "!biome.json")
	}
	if !args.Tailwind {
		copySource = append(copySource, "!postcss.config.mjs")
	}

	// Copy template files
	if err := copyTemplateFiles(templatesFS, templatePath, args.Root, copySource); err != nil {
		return fmt.Errorf("failed to copy template files: %w", err)
	}

	// Handle Rspack bundler configuration
	if args.Bundler == Rspack {
		if err := modifyNextConfigForRspack(args.Root, args.Mode); err != nil {
			return fmt.Errorf("failed to modify next.config for Rspack: %w", err)
		}
	}

	// Handle React Compiler configuration
	if args.ReactCompiler {
		if err := modifyNextConfigForReactCompiler(args.Root, args.Mode); err != nil {
			return fmt.Errorf("failed to modify next.config for React Compiler: %w", err)
		}
	}

	// Update tsconfig/jsconfig paths
	if err := updateConfigPaths(args.Root, args.Mode, args.SrcDir, args.ImportAlias); err != nil {
		return fmt.Errorf("failed to update config paths: %w", err)
	}

	// Update import aliases in source files
	if args.ImportAlias != "@/*" {
		if err := updateImportAliases(args.Root, args.ImportAlias); err != nil {
			return fmt.Errorf("failed to update import aliases: %w", err)
		}
	}

	// Handle src directory restructuring
	if args.SrcDir {
		if err := moveDirsToSrc(args.Root, isAPI, args.Template, args.Mode); err != nil {
			return fmt.Errorf("failed to move directories to src: %w", err)
		}
	}

	// Generate package.json
	if err := generatePackageJSON(args); err != nil {
		return fmt.Errorf("failed to generate package.json: %w", err)
	}

	return nil
}

// copyTemplateFiles copies files from embedded filesystem to target directory
func copyTemplateFiles(fsys embed.FS, templatePath, targetDir string, patterns []string) error {
	// Get all files matching patterns
	entries, err := fsys.ReadDir(templatePath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(templatePath, entry.Name())
		targetPath := filepath.Join(targetDir, renameFile(entry.Name()))

		if entry.IsDir() {
			// Recursively copy directory
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
			if err := copyTemplateFiles(fsys, sourcePath, targetPath, patterns); err != nil {
				return err
			}
		} else {
			// Check if file should be copied based on patterns
			if shouldCopyFile(entry.Name(), patterns) {
				data, err := fsys.ReadFile(sourcePath)
				if err != nil {
					return err
				}
				if err := os.WriteFile(targetPath, data, 0644); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// renameFile handles special file renames during copy
func renameFile(name string) string {
	switch name {
	case "gitignore":
		return ".gitignore"
	case "README-template.md":
		return "README.md"
	default:
		return name
	}
}

// shouldCopyFile checks if a file matches the copy patterns
func shouldCopyFile(filename string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.HasPrefix(pattern, "!") {
			// Exclusion pattern
			exclude := strings.TrimPrefix(pattern, "!")
			if matched, _ := doublestar.Match(exclude, filename); matched {
				return false
			}
		}
	}
	return true
}

// modifyNextConfigForRspack adds Rspack wrapper to next.config
func modifyNextConfigForRspack(root string, mode TemplateMode) error {
	configFile := "next.config.mjs"
	if mode == TS {
		configFile = "next.config.ts"
	}

	configPath := filepath.Join(root, configFile)
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	newContent := `import withRspack from "next-rspack";

` + strings.Replace(string(content), "export default nextConfig;", "export default withRspack(nextConfig);", 1)

	return os.WriteFile(configPath, []byte(newContent), 0644)
}

// modifyNextConfigForReactCompiler adds React Compiler config to next.config
func modifyNextConfigForReactCompiler(root string, mode TemplateMode) error {
	configFile := "next.config.mjs"
	if mode == TS {
		configFile = "next.config.ts"
	}

	configPath := filepath.Join(root, configFile)
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	newContent := strings.Replace(
		string(content),
		"/* config options here */\n",
		"/* config options here */\n  reactCompiler: true,\n",
		1,
	)

	return os.WriteFile(configPath, []byte(newContent), 0644)
}

// updateConfigPaths updates tsconfig/jsconfig paths for src directory and import alias
func updateConfigPaths(root string, mode TemplateMode, srcDir bool, importAlias string) error {
	configFile := "jsconfig.json"
	if mode == TS {
		configFile = "tsconfig.json"
	}

	configPath := filepath.Join(root, configFile)
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Update path based on srcDir
	if srcDir {
		contentStr = strings.Replace(contentStr, `"@/*": ["./*"]`, `"@/*": ["./src/*"]`, 1)
	}

	// Update import alias
	contentStr = strings.Replace(contentStr, `"@/*":`, fmt.Sprintf(`"%s":`, importAlias), 1)

	return os.WriteFile(configPath, []byte(contentStr), 0644)
}

// updateImportAliases replaces import aliases in all source files
func updateImportAliases(root, importAlias string) error {
	// Files to exclude from alias replacement
	excludePatterns := []string{
		"tsconfig.json",
		"jsconfig.json",
		".git/**/*",
		"**/fonts/**",
		"**/favicon.ico",
	}

	// Walk through all files
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Check if file should be excluded
		relPath, _ := filepath.Rel(root, path)
		for _, pattern := range excludePatterns {
			if matched, _ := doublestar.Match(pattern, relPath); matched {
				return nil
			}
		}

		// Read and replace content
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		newContent := strings.ReplaceAll(
			string(content),
			"@/",
			strings.TrimSuffix(importAlias, "*"),
		)

		if newContent != string(content) {
			return os.WriteFile(path, []byte(newContent), info.Mode())
		}

		return nil
	})
}

// moveDirsToSrc moves app, pages, and styles directories into src/
func moveDirsToSrc(root string, isAPI bool, template TemplateType, mode TemplateMode) error {
	srcDir := filepath.Join(root, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		return err
	}

	// Move directories
	for _, dirName := range SrcDirNames {
		oldPath := filepath.Join(root, dirName)
		newPath := filepath.Join(srcDir, dirName)

		if _, err := os.Stat(oldPath); err == nil {
			if err := os.Rename(oldPath, newPath); err != nil {
				return err
			}
		}
	}

	// Update page references if not API-only
	if !isAPI {
		isAppTemplate := strings.HasPrefix(string(template), "app")
		pageDir := "pages"
		pageFile := "index"
		if isAppTemplate {
			pageDir = "app"
			pageFile = "page"
		}

		ext := "js"
		if mode == TS {
			ext = "tsx"
		}

		indexPageFile := filepath.Join(srcDir, pageDir, fmt.Sprintf("%s.%s", pageFile, ext))
		content, err := os.ReadFile(indexPageFile)
		if err != nil {
			return err
		}

		oldRef := fmt.Sprintf("%s/%s", pageDir, pageFile)
		newRef := fmt.Sprintf("src/%s/%s", pageDir, pageFile)
		newContent := strings.Replace(string(content), oldRef, newRef, 1)

		return os.WriteFile(indexPageFile, []byte(newContent), 0644)
	}

	return nil
}

// generatePackageJSON creates and writes package.json for the new project
func generatePackageJSON(args InstallTemplateArgs) error {
	version := os.Getenv("NEXT_PRIVATE_TEST_VERSION")
	if version == "" {
		version = "latest" // Will be replaced with actual version from build
	}

	bundlerFlags := ""
	if args.Bundler == Webpack {
		bundlerFlags = " --webpack"
	}

	// Base package.json structure
	pkg := map[string]interface{}{
		"name":    args.AppName,
		"version": "0.1.0",
		"private": true,
		"scripts": map[string]string{
			"dev":   fmt.Sprintf("next dev%s", bundlerFlags),
			"build": fmt.Sprintf("next build%s", bundlerFlags),
			"start": "next start",
		},
		"dependencies": map[string]string{
			"react":     NextjsReactPeerVersion,
			"react-dom": NextjsReactPeerVersion,
			"next":      version,
		},
		"devDependencies": map[string]string{},
	}

	scripts := pkg["scripts"].(map[string]string)
	deps := pkg["dependencies"].(map[string]string)
	devDeps := pkg["devDependencies"].(map[string]string)

	// Add linting scripts
	isAPI := args.Template == AppAPI
	if args.Eslint && !isAPI {
		scripts["lint"] = "eslint"
	}
	if args.Biome && !isAPI {
		scripts["lint"] = "biome check"
		scripts["format"] = "biome format --write"
	}

	// Add bundler dependencies
	if args.Bundler == Rspack {
		deps["next-rspack"] = version
	}

	// Add React Compiler
	if args.ReactCompiler {
		devDeps["babel-plugin-react-compiler"] = "1.0.0"
	}

	// Add TypeScript dependencies
	if args.Mode == TS {
		devDeps["typescript"] = "^5"
		devDeps["@types/node"] = "^20"
		devDeps["@types/react"] = "^19"
		devDeps["@types/react-dom"] = "^19"
	}

	// Add Tailwind dependencies
	if args.Tailwind {
		devDeps["@tailwindcss/postcss"] = "^4"
		devDeps["tailwindcss"] = "^4"
	}

	// Add ESLint dependencies
	if args.Eslint {
		devDeps["eslint"] = "^9"
		devDeps["eslint-config-next"] = version
	}

	// Add Biome dependencies
	if args.Biome {
		devDeps["@biomejs/biome"] = "2.2.0"
	}

	// Handle API-only template
	if isAPI {
		delete(deps, "react")
		delete(deps, "react-dom")
		delete(devDeps, "@types/react-dom")
		delete(scripts, "lint")
		delete(scripts, "format")
	}

	// Remove devDependencies if empty
	if len(devDeps) == 0 {
		delete(pkg, "devDependencies")
	}

	// Sort dependencies
	pkg["dependencies"] = sorted(deps)
	if len(devDeps) > 0 {
		pkg["devDependencies"] = sorted(devDeps)
	}

	// Handle package manager specific configurations
	if args.PackageManager == "pnpm" {
		pnpmWorkspace := `packages:
  - .
ignoredBuiltDependencies:
  - sharp
  - unrs-resolver
`
		if err := os.WriteFile(filepath.Join(args.Root, "pnpm-workspace.yaml"), []byte(pnpmWorkspace), 0644); err != nil {
			return err
		}
	}

	if args.PackageManager == "bun" {
		pkg["ignoreScripts"] = []string{"sharp", "unrs-resolver"}
		pkg["trustedDependencies"] = []string{"sharp", "unrs-resolver"}
	}

	// Write package.json
	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return err
	}

	packageJSONPath := filepath.Join(args.Root, "package.json")
	return os.WriteFile(packageJSONPath, append(data, '\n'), 0644)
}
