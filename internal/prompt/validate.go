package prompt

import (
	"fmt"
	"strings"

	"github.com/yeasin2002/better-next-app/internal/util"
)

// ValidateProjectName validates npm package name rules using the util validator
func ValidateProjectName(name string) error {
	if name == "" {
		return nil // Empty is allowed, will use default
	}

	result := util.ValidateNpmPackageName(name)
	if !result.Valid {
		// Return the first error message
		if len(result.Errors) > 0 {
			return fmt.Errorf("%s", result.Errors[0])
		}
		return fmt.Errorf("invalid package name")
	}

	return nil
}

// ValidateImportAlias validates import alias format
func ValidateImportAlias(alias string) error {
	if alias == "" {
		return nil
	}

	if !strings.HasSuffix(alias, "/*") {
		return fmt.Errorf("import alias must end with '/*'")
	}

	return nil
}
