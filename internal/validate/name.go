package validate

import (
	"fmt"
	"regexp"
	"strings"
)

// npm built-in module names that cannot be used
var npmBuiltins = map[string]bool{
	"assert": true, "buffer": true, "child_process": true, "cluster": true,
	"console": true, "constants": true, "crypto": true, "dgram": true,
	"dns": true, "domain": true, "events": true, "fs": true, "http": true,
	"https": true, "module": true, "net": true, "os": true, "path": true,
	"punycode": true, "querystring": true, "readline": true, "repl": true,
	"stream": true, "string_decoder": true, "sys": true, "timers": true,
	"tls": true, "tty": true, "url": true, "util": true, "vm": true, "zlib": true,
}

// ValidationError represents npm name validation errors
type ValidationError struct {
	Field    string
	Problems []string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, strings.Join(e.Problems, ", "))
}

// ValidateNpmName validates a string against npm package naming rules
func ValidateNpmName(name string) error {
	var problems []string

	if name == "" {
		problems = append(problems, "name cannot be empty")
		return &ValidationError{Field: "name", Problems: problems}
	}

	if len(name) > 214 {
		problems = append(problems, "name must be less than 214 characters")
	}

	if strings.HasPrefix(name, ".") {
		problems = append(problems, "name cannot start with a period")
	}

	if strings.HasPrefix(name, "_") {
		problems = append(problems, "name cannot start with an underscore")
	}

	if strings.ToLower(name) != name {
		problems = append(problems, "name cannot contain capital letters")
	}

	if strings.TrimSpace(name) != name {
		problems = append(problems, "name cannot contain leading or trailing spaces")
	}

	// Check for URL-safe characters
	validChars := regexp.MustCompile(`^[a-z0-9._-]+$`)
	if !validChars.MatchString(name) {
		problems = append(problems, "name can only contain URL-safe characters (lowercase letters, numbers, hyphens, underscores, periods)")
	}

	// Check for npm built-in modules
	if npmBuiltins[name] {
		problems = append(problems, "name cannot be a Node.js built-in module name")
	}

	if len(problems) > 0 {
		return &ValidationError{Field: "name", Problems: problems}
	}

	return nil
}
