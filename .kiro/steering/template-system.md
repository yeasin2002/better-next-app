---
inclusion: always
---

# Template System

## Template Organization

Templates are embedded at compile time using Go's `embed` directive:

```go
package template

import "embed"

//go:embed templates/*
var templatesFS embed.FS
```

## Template Types

Nine template types exist with JS/TS variants:

- `app` - App Router
- `app-tw` - App Router + Tailwind
- `app-empty` - App Router (minimal)
- `app-tw-empty` - App Router + Tailwind (minimal)
- `app-api` - API-only (no React)
- `default` - Pages Router
- `default-tw` - Pages Router + Tailwind
- `default-empty` - Pages Router (minimal)
- `default-tw-empty` - Pages Router + Tailwind (minimal)

## Template Selection Algorithm

```go
func SelectTemplate(cfg *Config) string {
    var template string

    if cfg.APIOnly {
        template = "app-api"
    } else if cfg.AppRouter {
        if cfg.EmptyTemplate {
            template = "app-empty"
        } else {
            template = "app"
        }
        if cfg.Tailwind {
            template += "-tw"
        }
    } else {
        if cfg.EmptyTemplate {
            template = "default-empty"
        } else {
            template = "default"
        }
        if cfg.Tailwind {
            template += "-tw"
        }
    }

    if cfg.TypeScript {
        template += "-ts"
    }

    return template
}
```

## File Transformations

### 1. Renames During Copy

- `gitignore` → `.gitignore`
- `README-template.md` → `README.md`

### 2. Config Modifications

**For Rspack:**

```js
// Add import and wrap export
import withRspack from "next-rspack";
export default withRspack(nextConfig);
```

**For React Compiler:**

```js
// Add to next.config
experimental: {
  reactCompiler: true;
}
```

### 3. Path Config Updates

**For src directory:**
Update tsconfig.json/jsconfig.json paths:

```json
{
  "compilerOptions": {
    "paths": {
      "@/*": ["./src/*"]
    }
  }
}
```

### 4. Import Alias Replacement

Replace `@/` with custom alias in all source files:

- Include: `.js`, `.jsx`, `.ts`, `.tsx` files
- Exclude: configs, `.git`, fonts, favicon

### 5. Src Directory Restructuring

When `SrcDir` is enabled:

- Create `src/` folder
- Move `app/`, `pages/`, `styles/` into `src/`
- Update page references to include `src/` prefix

## Package.json Generation

Generate dynamically based on configuration:

```go
func GeneratePackageJSON(cfg *Config) map[string]interface{} {
    pkg := map[string]interface{}{
        "name":    cfg.ProjectName,
        "version": "0.1.0",
        "private": true,
    }

    // Add scripts
    scripts := map[string]string{
        "build": "next build",
        "start": "next start",
    }

    // Dev script varies by bundler
    if cfg.Bundler == "turbopack" {
        scripts["dev"] = "next dev --turbo"
    } else if cfg.Bundler == "webpack" {
        scripts["dev"] = "next dev --webpack"
    } else {
        scripts["dev"] = "next dev"
    }

    // Add lint script if linter enabled
    if cfg.Linter == "eslint" {
        scripts["lint"] = "eslint ."
    } else if cfg.Linter == "biome" {
        scripts["lint"] = "biome check ."
    }

    pkg["scripts"] = scripts

    // Add dependencies
    deps := map[string]string{
        "next": "latest",
    }

    if !cfg.APIOnly {
        deps["react"] = "latest"
        deps["react-dom"] = "latest"
    }

    if cfg.Bundler == "rspack" {
        deps["next-rspack"] = "latest"
    }

    pkg["dependencies"] = deps

    // Add devDependencies
    devDeps := map[string]string{}

    if cfg.TypeScript {
        devDeps["typescript"] = "latest"
        devDeps["@types/node"] = "latest"
        if !cfg.APIOnly {
            devDeps["@types/react"] = "latest"
            devDeps["@types/react-dom"] = "latest"
        }
    }

    if cfg.Tailwind {
        devDeps["tailwindcss"] = "latest"
        devDeps["postcss"] = "latest"
        devDeps["autoprefixer"] = "latest"
    }

    if cfg.Linter == "eslint" {
        devDeps["eslint"] = "latest"
        devDeps["eslint-config-next"] = "latest"
    } else if cfg.Linter == "biome" {
        devDeps["@biomejs/biome"] = "latest"
    }

    if cfg.ReactCompiler {
        devDeps["babel-plugin-react-compiler"] = "latest"
    }

    pkg["devDependencies"] = devDeps

    return pkg
}
```

## Template Installation Flow

1. Select template based on config
2. Copy files from embedded filesystem
3. Apply file renames
4. Modify config files (Rspack, React Compiler)
5. Replace import aliases if custom
6. Restructure for src directory if enabled
7. Generate and write package.json
