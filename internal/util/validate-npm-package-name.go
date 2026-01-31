package util

import (
	"net/url"
	"strings"
	"unicode"
)

// ValidationResult represents the result of npm package name validation
type ValidationResult struct {
	ValidForNewPackages bool
	ValidForOldPackages bool
	Errors              []string
	Warnings            []string
}

// Node.js/io.js core modules and reserved names
var builtinModules = map[string]bool{
	"http":           true,
	"https":          true,
	"stream":         true,
	"events":         true,
	"fs":             true,
	"path":           true,
	"url":            true,
	"util":           true,
	"os":             true,
	"crypto":         true,
	"buffer":         true,
	"querystring":    true,
	"string_decoder": true,
	"timers":         true,
	"tty":            true,
	"net":            true,
	"dgram":          true,
	"dns":            true,
	"tls":            true,
	"child_process":  true,
	"cluster":        true,
	"zlib":           true,
	"assert":         true,
	"punycode":       true,
	"domain":         true,
	"constants":      true,
	"process":        true,
	"console":        true,
	"vm":             true,
	"repl":           true,
	"readline":       true,
	"module":         true,
	"v8":             true,
	"async_hooks":    true,
	"perf_hooks":     true,
	"worker_threads": true,
	"inspector":      true,
	"trace_events":   true,
	"node_modules":   true,
	"favicon.ico":    true,
}

// ValidateNpmPackageName validates an npm package name according to npm rules
// Returns a ValidationResult with validation status, errors, and warnings
func ValidateNpmPackageName(name string) ValidationResult {
	result := ValidationResult{
		ValidForNewPackages: true,
		ValidForOldPackages: true,
		Errors:              []string{},
		Warnings:            []string{},
	}

	// Rule: package name length should be greater than zero
	if len(name) == 0 {
		result.ValidForNewPackages = false
		result.ValidForOldPackages = false
		result.Errors = append(result.Errors, "name length must be greater than zero")
		return result
	}

	// Rule: package name length cannot exceed 214
	if len(name) > 214 {
		result.ValidForNewPackages = false
		result.Warnings = append(result.Warnings, "name can no longer contain more than 214 characters")
	}

	// Rule: name cannot contain leading or trailing spaces
	if strings.TrimSpace(name) != name {
		result.ValidForNewPackages = false
		result.ValidForOldPackages = false
		result.Errors = append(result.Errors, "name cannot contain leading or trailing spaces")
	}

	// Extract the package name without scope for validation
	packageName := name
	if strings.HasPrefix(name, "@") {
		parts := strings.SplitN(name, "/", 2)
		if len(parts) == 2 {
			packageName = parts[1]
			// Validate scope name
			scopeName := parts[0]
			if len(scopeName) <= 1 {
				result.ValidForNewPackages = false
				result.ValidForOldPackages = false
				result.Errors = append(result.Errors, "scope name cannot be empty")
			}
			// Validate package name after scope is not empty
			if len(packageName) == 0 {
				result.ValidForNewPackages = false
				result.ValidForOldPackages = false
				result.Errors = append(result.Errors, "package name after scope cannot be empty")
			}
		} else {
			result.ValidForNewPackages = false
			result.ValidForOldPackages = false
			result.Errors = append(result.Errors, "scoped package name must include a package name after the scope")
		}
	}

	// Rule: package name should not start with . or _ (unless scoped)
	if !strings.HasPrefix(name, "@") {
		if strings.HasPrefix(packageName, ".") || strings.HasPrefix(packageName, "_") {
			result.ValidForNewPackages = false
			result.ValidForOldPackages = false
			result.Errors = append(result.Errors, "name cannot start with a period or underscore")
		}
	}

	// Rule: all characters must be lowercase
	hasUpperCase := false
	for _, r := range name {
		if unicode.IsUpper(r) {
			hasUpperCase = true
			break
		}
	}
	if hasUpperCase {
		result.ValidForNewPackages = false
		result.Warnings = append(result.Warnings, "name can no longer contain capital letters")
	}

	// Rule: name should not contain spaces
	if strings.Contains(name, " ") {
		result.ValidForNewPackages = false
		result.ValidForOldPackages = false
		result.Errors = append(result.Errors, "name cannot contain spaces")
	}

	// Rule: name should not contain special characters: ~)('!*
	specialChars := []string{"~", ")", "(", "'", "!", "*"}
	for _, char := range specialChars {
		if strings.Contains(name, char) {
			result.ValidForNewPackages = false
			result.ValidForOldPackages = false
			result.Errors = append(result.Errors, "name can only contain URL-friendly characters")
			break
		}
	}

	// Rule: name must not contain any non-url-safe characters
	// URL encode and check if it changes
	encoded := url.PathEscape(name)
	if encoded != name {
		// Check if it's just because of @ or / (which are allowed in scoped packages)
		testName := strings.ReplaceAll(name, "@", "")
		testName = strings.ReplaceAll(testName, "/", "")
		testEncoded := url.PathEscape(testName)
		if testEncoded != testName {
			result.ValidForNewPackages = false
			result.ValidForOldPackages = false
			result.Errors = append(result.Errors, "name can only contain URL-friendly characters")
		}
	}

	// Rule: name cannot be a Node.js core module or reserved name
	checkName := strings.ToLower(packageName)
	if builtinModules[checkName] {
		result.ValidForNewPackages = false
		result.ValidForOldPackages = false
		result.Errors = append(result.Errors, "name cannot be a core module or reserved name")
	}

	// Update ValidForOldPackages based on errors
	if len(result.Errors) > 0 {
		result.ValidForOldPackages = false
	}

	return result
}

// IsValidPackageName is a convenience function that returns true if the name is valid for new packages
func IsValidPackageName(name string) bool {
	result := ValidateNpmPackageName(name)
	return result.ValidForNewPackages
}
