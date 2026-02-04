# NPM Publishing Troubleshooting

## Error: 403 Two-factor authentication required

### Error Message:
```
npm error 403 Two-factor authentication or granular access token with bypass 2fa enabled is required to publish packages.
```

### Cause:
Your NPM token doesn't have "Bypass 2FA for automation" enabled.

### Solution:

#### Step 1: Generate a New Token

1. Go to: https://www.npmjs.com/settings/YOUR_USERNAME/tokens
2. Click "Generate New Token"
3. Select "Granular Access Token"

#### Step 2: Configure Token (CRITICAL SETTINGS)

**Token name:**
```
better-next-app-ci-v2
```

**Expiration:**
```
90 days
```

**Packages and scopes - Permissions:**
```
☑ Read and write
```

**⚠️ CRITICAL: Enable Bypass 2FA**
```
☑ Bypass 2FA for automation  ← YOU MUST CHECK THIS BOX!
```

**Organizations - Permissions:**
```
○ No access
```

#### Step 3: Copy Token

Copy the token immediately (starts with `npm_...`)

#### Step 4: Update GitHub Secret

1. Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Click on `NPM_TOKEN` (existing secret)
3. Click "Update"
4. Paste your new token
5. Click "Update secret"

#### Step 5: Test Publishing

```bash
task npm:publish
```

### Visual Guide

When creating the token, you should see this option:

```
┌─────────────────────────────────────────────┐
│  Packages and scopes                         │
│                                              │
│  Permissions:                                │
│  ● Read and write                            │
│                                              │
│  ☑ Bypass 2FA for automation  ← CHECK THIS! │
│                                              │
└─────────────────────────────────────────────┘
```

**If you don't see this checkbox**, you're looking at the wrong section. Make sure you're in the "Packages and scopes" section, not "Organizations".

## Error: sed command not found (Windows)

### Error Message:
```
"sed": executable file not found in $PATH
```

### Cause:
Windows doesn't have `sed` by default.

### Solution:
This has been fixed in the Taskfile. Update your repository:

```bash
git pull origin main
```

The Taskfile now uses PowerShell for Windows and `sed` for Unix systems.

## Error: Package name invalid

### Error Message:
```
npm warn publish "bin[create-better-next-app]" script name bin/create-better-next-app.js was invalid and removed
```

### Cause:
This is just a warning. NPM auto-corrects it during publish.

### Solution:
You can safely ignore this warning. The package will publish correctly.

To fix it permanently, run:
```bash
cd npm
npm pkg fix
```

## Error: Token expired

### Error Message:
```
npm error 401 Unauthorized - PUT https://registry.npmjs.org/create-better-next-app
```

### Cause:
Your NPM token has expired (90 days maximum).

### Solution:

1. Generate a new token (see steps above)
2. Update GitHub secret
3. Set a calendar reminder for 85 days from now

## Error: Package already exists

### Error Message:
```
npm error 403 You cannot publish over the previously published versions
```

### Cause:
You're trying to publish a version that already exists.

### Solution:

Update the version:
```bash
cd npm
npm version patch  # or minor, or major
git add package.json
git commit -m "chore: bump version"
```

Or create a new git tag:
```bash
git tag -a v0.0.3 -m "Release v0.0.3"
git push origin v0.0.3
```

## Error: Not logged in

### Error Message:
```
npm error code ENEEDAUTH
npm error need auth This command requires you to be logged in to https://registry.npmjs.org/
```

### Cause:
You're not logged in to NPM locally.

### Solution:

```bash
npm login
```

**Note:** This creates a 2-hour session token for local use only. For CI/CD, you still need a Granular Access Token in GitHub Secrets.

## Verification Checklist

Before publishing, verify:

- [ ] NPM token has "Bypass 2FA" enabled
- [ ] Token expiration is set (90 days max)
- [ ] Token has "Read and write" permissions for packages
- [ ] GitHub secret `NPM_TOKEN` is updated with new token
- [ ] You're logged in locally: `npm whoami`
- [ ] Package name is available or you own it
- [ ] Version number is incremented

## Quick Test

Run this to verify your setup:

```bash
# Check if logged in
npm whoami

# Verify setup
task npm:setup

# Test package
task npm:test

# Try publishing
task npm:publish
```

## Still Having Issues?

1. Check the full guide: [npm-token-guide.md](./npm-token-guide.md)
2. Verify token settings at: https://www.npmjs.com/settings/YOUR_USERNAME/tokens
3. Check GitHub secret at: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
4. Review NPM security update: https://github.blog/changelog/2025-12-09-npm-classic-tokens-revoked-session-based-auth-and-cli-token-management-now-available/

## Common Mistakes

1. ❌ **Forgetting to check "Bypass 2FA"** - Most common issue!
2. ❌ **Using session token for CI/CD** - Session tokens expire after 2 hours
3. ❌ **Wrong secret name** - Must be exactly `NPM_TOKEN`
4. ❌ **Token expired** - Tokens expire after 90 days
5. ❌ **Wrong permissions** - Must have "Read and write" for packages
