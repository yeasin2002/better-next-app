package validate

import (
	"fmt"
	"os"
	"path/filepath"
)

// DirectoryError represents directory validation errors
type DirectoryError struct {
	Path             string
	ConflictingFiles []string
}

func (e *DirectoryError) Error() string {
	return fmt.Sprintf("directory %s contains conflicting files", e.Path)
}

// allowedFiles are files that can exist in a "empty" directory
var allowedFiles = map[string]bool{
	".git":       true,
	".gitignore": true,
	".gitkeep":   true,
	"LICENSE":    true,
	"license":    true,
	"README.md":  true,
	"readme.md":  true,
}

// ValidateDirectory checks if a directory is safe to use for project creation
func ValidateDirectory(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		// Directory doesn't exist, that's fine
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to check directory: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path exists but is not a directory")
	}

	// Check if directory is writable
	testFile := filepath.Join(absPath, ".write-test")
	f, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("directory is not writable")
	}
	f.Close()
	os.Remove(testFile)

	return nil
}

// IsFolderEmpty checks if a directory is empty or contains only allowed files
func IsFolderEmpty(path string) (bool, []string, error) {
	entries, err := os.ReadDir(path)
	if os.IsNotExist(err) {
		return true, nil, nil
	}
	if err != nil {
		return false, nil, err
	}

	var conflicting []string
	for _, entry := range entries {
		name := entry.Name()
		if !allowedFiles[name] {
			conflicting = append(conflicting, name)
		}
	}

	return len(conflicting) == 0, conflicting, nil
}

// EnsureDirectory creates a directory if it doesn't exist
func EnsureDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}
