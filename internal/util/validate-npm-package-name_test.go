package util

import (
	"strings"
	"testing"
)

func TestValidateNpmPackageName_ValidNames(t *testing.T) {
	validNames := []string{
		"some-package",
		"example.com",
		"under_score",
		"123numeric",
		"@npm/thingy",
		"@jane/foo.js",
		"my-app",
		"react-dom",
		"lodash",
	}

	for _, name := range validNames {
		t.Run(name, func(t *testing.T) {
			result := ValidateNpmPackageName(name)
			if !result.Valid {
				t.Errorf("Expected %q to be valid, got errors: %v", name, result.Errors)
			}
			if len(result.Errors) > 0 {
				t.Errorf("Expected no errors for %q, got: %v", name, result.Errors)
			}
		})
	}
}

func TestValidateNpmPackageName_InvalidNames(t *testing.T) {
	tests := []struct {
		name          string
		expectedError string
	}{
		{"", "name length must be greater than zero"},
		{"excited!", "name can only contain URL-friendly characters"},
		{" leading-space", "name cannot contain leading or trailing spaces"},
		{"trailing-space ", "name cannot contain leading or trailing spaces"},
		{".start-with-dot", "name cannot start with a period or underscore"},
		{"_start-with-underscore", "name cannot start with a period or underscore"},
		{"has spaces", "name cannot contain spaces"},
		{"has~tilde", "name can only contain URL-friendly characters"},
		{"has(parens)", "name can only contain URL-friendly characters"},
		{"has'quote", "name can only contain URL-friendly characters"},
		{"has!exclaim", "name can only contain URL-friendly characters"},
		{"has*asterisk", "name can only contain URL-friendly characters"},
		{"http", "name cannot be a core module or reserved name"},
		{"stream", "name cannot be a core module or reserved name"},
		{"node_modules", "name cannot be a core module or reserved name"},
		{"favicon.ico", "name cannot be a core module or reserved name"},
		{"UpperCase", "name cannot contain uppercase letters"},
		{"MixedCase-Package", "name cannot contain uppercase letters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateNpmPackageName(tt.name)
			if result.Valid {
				t.Errorf("Expected %q to be invalid", tt.name)
			}
			if len(result.Errors) == 0 {
				t.Errorf("Expected errors for %q, got none", tt.name)
			}
		})
	}
}

func TestValidateNpmPackageName_TooLong(t *testing.T) {
	// Create a name longer than 214 characters
	longName := strings.Repeat("a", 215)
	result := ValidateNpmPackageName(longName)
	
	if result.Valid {
		t.Error("Expected long name to be invalid")
	}
	if len(result.Errors) == 0 {
		t.Error("Expected error about length")
	}
}

func TestValidateNpmPackageName_ScopedPackages(t *testing.T) {
	tests := []struct {
		name     string
		valid    bool
		hasError bool
	}{
		{"@scope/package", true, false},
		{"@scope/package-name", true, false},
		{"@my-org/my-package", true, false},
		{"@", false, true},
		{"@scope", false, true},
		{"@scope/", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateNpmPackageName(tt.name)
			if result.Valid != tt.valid {
				t.Errorf("Expected %q valid=%v, got valid=%v, errors: %v",
					tt.name, tt.valid, result.Valid, result.Errors)
			}
			if tt.hasError && len(result.Errors) == 0 {
				t.Errorf("Expected errors for %q, got none", tt.name)
			}
		})
	}
}

func TestIsValidPackageName(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
	}{
		{"valid-package", true},
		{"@scope/package", true},
		{"Invalid-Package", false},
		{"has spaces", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPackageName(tt.name)
			if result != tt.valid {
				t.Errorf("IsValidPackageName(%q) = %v, want %v", tt.name, result, tt.valid)
			}
		})
	}
}
