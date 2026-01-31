---
inclusion: always
---

# Validation Rules

## NPM Package Name Validation

Must follow npm naming rules:

```go
func ValidateNpmName(name string) error {
    // Length: 1-214 characters
    if len(name) == 0 || len(name) > 214 {
        return fmt.Errorf("name length must be between 1 and 214 characters")
    }

    // Cannot start with . or _
    if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
        return fmt.Errorf("name cannot start with . or _")
    }

    // No uppercase letters
    if name != strings.ToLower(name) {
        return fmt.Errorf("name cannot contain uppercase letters")
    }

    // No non-URL-safe characters
    // Only allow: a-z, 0-9, -, _, @, /
    validChars := regexp.MustCompile(`^[@a-z0-9/_-]+$`)
    if !validChars.MatchString(name) {
        return fmt.Errorf("name contains invalid characters")
    }

    // Cannot be npm built-in modules
    builtins := []string{"http", "stream", "path", "fs", "events", ...}
    for _, builtin := range builtins {
        if name == builtin {
            return fmt.Errorf("name cannot be a Node.js built-in module")
        }
    }

    return nil
}
```

## Directory Validation

Check if directory is safe to use:

```go
func ValidateDirectory(path string) error {
    // Check if path exists
    info, err := os.Stat(path)
    if os.IsNotExist(err) {
        // Directory doesn't exist - OK to create
        return nil
    }

    if err != nil {
        return fmt.Errorf("cannot access directory: %w", err)
    }

    // Must be a directory
    if !info.IsDir() {
        return fmt.Errorf("path exists but is not a directory")
    }

    // Check write permissions
    testFile := filepath.Join(path, ".write-test")
    if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
        return fmt.Errorf("directory is not writable")
    }
    os.Remove(testFile)

    // Check if directory is empty (allowing certain files)
    isEmpty, conflicts, err := IsFolderEmpty(path)
    if err != nil {
        return err
    }

    if !isEmpty {
        return &DirectoryError{
            Path:             path,
            ConflictingFiles: conflicts,
        }
    }

    return nil
}
```

## Folder Emptiness Check

Allow certain files to exist:

```go
func IsFolderEmpty(path string) (bool, []string, error) {
    allowedFiles := map[string]bool{
        ".git":       true,
        ".gitignore": true,
        "LICENSE":    true,
        "README.md":  true,
    }

    entries, err := os.ReadDir(path)
    if err != nil {
        return false, nil, err
    }

    var conflicts []string

    for _, entry := range entries {
        if !allowedFiles[entry.Name()] {
            conflicts = append(conflicts, entry.Name())
        }
    }

    return len(conflicts) == 0, conflicts, nil
}
```

## CI Environment Detection

Skip prompts in CI environments:

```go
func IsCI() bool {
    ciEnvVars := []string{
        "CI",
        "CONTINUOUS_INTEGRATION",
        "BUILD_NUMBER",
        "JENKINS_URL",
        "TRAVIS",
        "CIRCLECI",
        "GITHUB_ACTIONS",
        "GITLAB_CI",
    }

    for _, envVar := range ciEnvVars {
        if os.Getenv(envVar) != "" {
            return true
        }
    }

    return false
}
```

## Network Connectivity Check

Detect if online before attempting downloads:

```go
func IsOnline() bool {
    timeout := 5 * time.Second
    client := &http.Client{Timeout: timeout}

    // Try multiple endpoints
    endpoints := []string{
        "https://registry.npmjs.org",
        "https://api.github.com",
        "https://www.google.com",
    }

    for _, endpoint := range endpoints {
        resp, err := client.Head(endpoint)
        if err == nil && resp.StatusCode < 500 {
            return true
        }
    }

    return false
}
```

## Input Validation for Prompts

Validate user input in real-time:

```go
func validateProjectName(name string) error {
    if name == "" {
        return fmt.Errorf("project name cannot be empty")
    }

    return ValidateNpmName(name)
}

func validateImportAlias(alias string) error {
    if alias == "" {
        return fmt.Errorf("import alias cannot be empty")
    }

    if !strings.HasSuffix(alias, "/*") {
        return fmt.Errorf("import alias must end with /*")
    }

    return nil
}
```

## Error Messages

Provide user-friendly error messages:

**Invalid Project Name:**

```
Could not create project because of npm naming restrictions:
  • Name cannot contain capital letters
  • Name length must be less than 214 characters
```

**Directory Not Empty:**

```
The directory my-app contains files that could conflict:
  • index.html
  • package.json

Either use a new directory name, or remove the files listed above.
```

**No Write Permission:**

```
The application path is not writable, please check folder permissions.
It is likely you do not have write permissions for this folder.
```
