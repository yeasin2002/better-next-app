# Product Overview

Better Next App is a high-performance CLI tool for scaffolding Next.js projects, written in Go. It's a complete rewrite of `create-next-app` that provides faster startup times, single binary distribution, and feature parity with the original TypeScript implementation.

## Key Value Propositions

- Single binary distribution with no Node.js required to run the CLI
- Faster startup through native compilation
- Cross-platform support (Windows, macOS, Linux)
- Built-in concurrency using Go's goroutines
- Template embedding via Go's `embed` directive

## Core Features

- Interactive and non-interactive modes
- Multiple templates (App Router, with/without Tailwind)
- this CLI dosen't support nextjs page router, only work with nextjs app router
- Smart package manager detection (npm, pnpm, yarn, bun)
- Flexible configuration (TypeScript, ESLint, Biome, React Compiler)
- GitHub example downloading
- Preference persistence
- Full feature parity with the original TypeScript version


## How people will use it: 

##### npm
```shell
npx create better-next-app@@latest 
```

##### bun
```shell
bunx  create better-next-app@@latest 
```

##### pnpm
```shell
pnpm dlx create better-next-app@@latest 
```


#### Process 
On installation, you'll see the following prompts:

```terminal
What is your project named? my-app
Would you like to use the recommended Next.js defaults?
    Yes, use recommended defaults - TypeScript, ESLint, Tailwind CSS, App Router, Turbopack
    No, reuse previous settings
    No, customize settings - Choose your own preferences
```

If you choose to customize settings, you'll see the following prompts:

```
Would you like to use TypeScript? No / Yes
Which linter would you like to use? ESLint / Biome / None
Would you like to use React Compiler? No / Yes
Would you like to use Tailwind CSS? No / Yes
Would you like your code inside a `src/` directory? No / Yes
Would you like to customize the import alias (`@/*` by default)? No / Yes
What import alias would you like configured? @/*
```

Note: will will be expended, write now we will create this way but in feature will be add new feature like adding database, OTM etc. 