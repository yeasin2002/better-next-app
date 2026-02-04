---
inclusion: always
---

# Template Structure

## Overview

Templates are organized in the `/templates` directory and will be embedded into the Go binary using the `embed` directive. All templates are based on Next.js App Router only (Pages Router is not supported).

## Template Organization

```
templates/
├── app/                    # Standard App Router
│   ├── js/                 # JavaScript variant
│   └── ts/                 # TypeScript variant
├── app-tw/                 # App Router + Tailwind CSS
│   ├── js/
│   └── ts/
├── app-empty/              # Minimal App Router (no boilerplate)
│   ├── js/
│   └── ts/
├── app-tw-empty/           # Minimal App Router + Tailwind
│   ├── js/
│   └── ts/
├── app-api/                # API-only (no React components)
│   ├── js/
│   └── ts/
└── .prettierrc.json        # Prettier config for templates
```

Note: Templates are embedded into the Go binary at compile time using the `//go:embed templates` directive in `main.go`.

## Template Variants

Five main template types, each with JavaScript and TypeScript variants:

1. **app** - Full-featured App Router with example components and styling
2. **app-tw** - App Router with Tailwind CSS integration
3. **app-empty** - Minimal App Router template without boilerplate
4. **app-tw-empty** - Minimal App Router with Tailwind CSS
5. **app-api** - API-only template (no React, just API routes)

## Standard Template Structure

Each template variant (js/ts) contains:

### Configuration Files
- `next.config.mjs` or `next.config.ts` - Next.js configuration
- `tsconfig.json` (TypeScript only) - TypeScript configuration
- `jsconfig.json` (JavaScript only) - JavaScript configuration
- `eslint.config.mjs` - ESLint configuration
- `biome.json` - Biome linter configuration
- `postcss.config.mjs` (Tailwind variants) - PostCSS configuration
- `gitignore` - Git ignore rules (renamed to `.gitignore` during installation)
- `.env.example` - Environment variable template
- `README-template.md` - Project README (renamed to `README.md` during installation)

### Source Files
- `app/` directory - Next.js App Router structure
  - `layout.tsx/jsx` - Root layout component
  - `page.tsx/jsx` - Home page component
  - `globals.css` - Global styles
  - `page.module.css` (non-Tailwind variants) - CSS modules
  - `favicon.ico` - Favicon
- `public/` directory (full templates only) - Static assets
  - SVG icons (file.svg, globe.svg, next.svg, vercel.svg, window.svg)

## Template Selection Logic

The CLI selects templates based on user configuration:

```
if API-only:
    → app-api/{js|ts}
else if Tailwind + Empty:
    → app-tw-empty/{js|ts}
else if Tailwind:
    → app-tw/{js|ts}
else if Empty:
    → app-empty/{js|ts}
else:
    → app/{js|ts}
```

## File Transformations During Installation

When copying templates to the user's project:

1. **Renames**
   - `gitignore` → `.gitignore`
   - `README-template.md` → `README.md`

2. **Content Modifications**
   - Replace import alias `@/` with custom alias if specified
   - Update config paths for `src/` directory if enabled
   - Add bundler-specific configuration (Rspack, Webpack)
   - Add React Compiler configuration if enabled

3. **Directory Restructuring**
   - Move `app/` into `src/app/` if `src/` directory is enabled
   - Update all path references in configs

## Package.json Generation

The `package.json` is dynamically generated (not stored in templates) based on:
- Project name
- Language choice (TypeScript/JavaScript)
- Linter choice (ESLint/Biome/None)
- Tailwind CSS inclusion
- React Compiler enablement
- Bundler choice (Turbopack/Webpack/Rspack)

## Notes

- Templates are maintained separately from the Go code
- All templates follow Next.js App Router conventions
- Templates include both linter configurations (ESLint and Biome) by default
- The CLI filters and applies the appropriate configuration based on user choices
