---
inclusion: always
---

# Git Hooks

This project uses [Lefthook](https://github.com/evilmartians/lefthook) for managing git hooks and [commitlint](https://github.com/conventionalcommit/commitlint) for commit message validation.

## Setup

Hooks are automatically installed when you run:

```bash
task setup
```

Or manually:

```bash
# Install commitlint
go install github.com/conventionalcommit/commitlint@latest

# Install and setup Lefthook
lefthook install
```

## Commit Message Hook

Enforces [Conventional Commits](https://www.conventionalcommits.org/) format using commitlint:

```
<type>(<scope>): <subject>
```

### Allowed Types

- **feat** - A new feature
- **fix** - A bug fix
- **docs** - Documentation changes
- **style** - Code style changes (formatting, etc.)
- **refactor** - Code refactoring
- **test** - Adding or updating tests
- **chore** - Maintenance tasks
- **perf** - Performance improvements
- **ci** - CI/CD changes
- **build** - Build system changes
- **revert** - Reverts a previous commit

### Allowed Scopes (optional)

- **cli** - CLI command changes
- **config** - Configuration changes
- **prompt** - Interactive prompt changes
- **template** - Template changes
- **util** - Utility function changes
- **deps** - Dependency updates
- **release** - Release-related changes

### Rules

- Header must be between 10-100 characters
- Description must be at least 3 characters
- Body lines max 100 characters (if present)
- Footer lines max 100 characters (if present)

### Examples

```bash
# Good commits
git commit -m "feat: add user authentication"
git commit -m "fix(api): resolve null pointer error"
git commit -m "docs: update README with installation steps"
git commit -m "test(util): add unit tests for validation"
git commit -m "chore(deps): update dependencies"

# Bad commits (will be rejected)
git commit -m "added new feature"  # Wrong format
git commit -m "Fixed bug"          # Wrong format
git commit -m "WIP"                # Too short
git commit -m "feat: x"            # Description too short
```

## Configuration

Commit message rules are configured in `.commitlint.yaml`. You can customize:
- Allowed types and scopes
- Length constraints
- Character sets
- Footer requirements

See [commitlint documentation](https://github.com/conventionalcommit/commitlint#available-rules) for all available rules.

## Pre-commit Hooks

Runs automatically before each commit:

1. **Format** - Formats Go code (`task fmt`)
2. **Vet** - Runs go vet (`task vet`)
3. **Lint** - Runs golangci-lint (`task lint`)
4. **Test** - Runs all tests (`task test`)

All hooks run in parallel for speed.

## Pre-push Hooks

Runs automatically before pushing:

- **Check** - Runs full check suite (`task check`)

## Skipping Hooks

If you need to skip hooks temporarily:

```bash
# Skip commit-msg validation
git commit --no-verify -m "WIP: work in progress"

# Skip pre-commit hooks
git commit --no-verify

# Skip pre-push hooks
git push --no-verify
```

## Manual Testing

Test your commit message before committing:

```bash
# Test a message
echo "feat: add new feature" | commitlint lint

# Test from file
commitlint lint --message-file .git/COMMIT_EDITMSG

# Debug commitlint
commitlint debug
```

## Troubleshooting

If hooks aren't running:

```bash
# Reinstall hooks
lefthook install

# Check hook status
lefthook run commit-msg
lefthook run pre-commit

# Verify commitlint is installed
commitlint --version
```

If commitlint is not found:

```bash
# Install commitlint
go install github.com/conventionalcommit/commitlint@latest

# Ensure $GOPATH/bin is in your PATH
```
