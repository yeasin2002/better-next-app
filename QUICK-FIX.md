# Quick Fix for Your 403 Error

## The Problem

Your NPM token doesn't have "Bypass 2FA for automation" enabled. This is a NEW requirement as of December 9, 2025.

## The Solution (5 Minutes)

### Step 1: Generate New Token

Go to: https://www.npmjs.com/settings/YOUR_USERNAME/tokens

Click "Generate New Token" → "Granular Access Token"

### Step 2: Fill in These EXACT Settings

```
Token name: better-next-app-ci-v2
Expiration: 90 days

Packages and scopes:
  ● Read and write
  ☑ Bypass 2FA for automation  ← MUST CHECK THIS!

Organizations:
  ○ No access
```

### Step 3: Copy Token

Copy the token (starts with `npm_...`)

### Step 4: Update GitHub Secret

1. Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Click on `NPM_TOKEN`
3. Click "Update"
4. Paste new token
5. Click "Update secret"

### Step 5: Try Again

```bash
task npm:publish
```

## Visual Guide

When creating the token, look for this checkbox in the "Packages and scopes" section:

```
┌─────────────────────────────────────────┐
│  Packages and scopes                     │
│                                          │
│  Permissions:                            │
│  ● Read and write                        │
│                                          │
│  ☑ Bypass 2FA for automation  ← HERE!   │
│                                          │
└─────────────────────────────────────────┘
```

## Why This Happened

NPM changed their security requirements on December 9, 2025:
- Classic tokens were permanently revoked
- Granular tokens now require "Bypass 2FA" for CI/CD
- Write tokens limited to 90 days maximum

## Need More Help?

See: [docs/TROUBLESHOOTING-NPM.md](./docs/TROUBLESHOOTING-NPM.md)
