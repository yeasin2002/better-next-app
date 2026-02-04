# NPM Token Generation Guide

A visual guide to help you generate the correct NPM token for your project.

## Which Token Type Should I Use?

### For This Project: **Granular Access Token** ✅

**Why?**
- ✅ Modern and recommended by NPM
- ✅ Fine-grained permissions (only what you need)
- ✅ Better security
- ✅ Works with GitHub Actions
- ✅ Can be scoped to specific packages

### Classic Token (Legacy)

**Don't use unless:**
- ❌ You're maintaining an old project
- ❌ You need backward compatibility

## Step-by-Step Token Generation

### Step 1: Access Token Page

**URL:** https://www.npmjs.com/settings/YOUR_USERNAME/tokens

Replace `YOUR_USERNAME` with your actual NPM username.

**Visual Guide:**
```
NPM Website → Profile Picture (top right) → Access Tokens
```

### Step 2: Click "Generate New Token"

You'll see two options:

```
┌─────────────────────────────────────┐
│  Granular Access Token              │  ← Choose this one! ✅
│  (Recommended)                       │
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│  Classic Token                       │  ← Don't use
│  (Legacy)                            │
└─────────────────────────────────────┘
```

**Select:** Granular Access Token

### Step 3: Configure Token

#### General Section

**Token name:** (Required)
```
better-next-app-ci
```
Or any descriptive name like:
- `github-actions-publish`
- `ci-automation`
- `better-next-app-automation`

**Description:** (Optional but recommended)
```
GitHub Actions automation for publishing create-better-next-app
```

**Expiration:** (Required)

Choose one:
- ✅ **90 days** - Good for testing
- ✅ **1 year** - Good for production (set calendar reminder)
- ⚠️ **Custom** - Choose your own duration
- ❌ **No expiration** - Not recommended for security

**Recommendation:** Use **90 days** for initial setup, then switch to **1 year** for production.

#### Packages and Scopes Section

**Permissions:** (Required)

```
┌─────────────────────────────────────┐
│  Permissions                         │
│                                      │
│  ○ No access                         │
│  ● Read and write  ← Select this! ✅ │
│  ○ Read-only                         │
└─────────────────────────────────────┘
```

**Select:** Read and write

**Why?** This allows GitHub Actions to publish new versions of your package.

**Select organizations:** (Optional)

Leave empty unless you're publishing under an organization.

#### Organizations Section

**Permissions:** (Required)

```
┌─────────────────────────────────────┐
│  Permissions                         │
│                                      │
│  ● No access  ← Keep this! ✅        │
│  ○ Read and write                    │
│  ○ Read-only                         │
└─────────────────────────────────────┘
```

**Select:** No access

**Why?** You're publishing as an individual user, not an organization.

### Step 4: Review Summary

Before generating, verify the summary shows:

```
This token will:
✓ Provide read and write access to packages and scopes
✓ Provide read and write access to 0 organizations
✓ Expires on [your selected date]
```

### Step 5: Generate Token

1. Click **"Generate token"** button
2. Token appears on screen (starts with `npm_...`)
3. **CRITICAL:** Copy it immediately!
4. Store it securely (you won't see it again)

**Token format:**
```
npm_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

### Step 6: Add to GitHub Secrets

#### 6.1 Navigate to Repository Secrets

**URL:** https://github.com/yeasin2002/better-next-app/settings/secrets/actions

**Visual Path:**
```
Repository → Settings → Secrets and variables → Actions
```

#### 6.2 Create New Secret

Click **"New repository secret"** (green button)

**Form:**
```
┌─────────────────────────────────────┐
│  Name *                              │
│  ┌─────────────────────────────┐   │
│  │ NPM_TOKEN                    │   │  ← Exactly this!
│  └─────────────────────────────┘   │
│                                      │
│  Secret *                            │
│  ┌─────────────────────────────┐   │
│  │ npm_xxxxxxxxxxxxxxxx         │   │  ← Paste your token
│  └─────────────────────────────┘   │
│                                      │
│  [Add secret]                        │
└─────────────────────────────────────┘
```

**Important:**
- Name must be **exactly** `NPM_TOKEN` (case-sensitive)
- Paste the full token including `npm_` prefix

#### 6.3 Verify Secret

After adding, you should see:

```
Repository secrets
┌─────────────────────────────────────┐
│ NPM_TOKEN                            │
│ Updated just now                     │
│ [Update] [Remove]                    │
└─────────────────────────────────────┘
```

## Verification Checklist

Before proceeding, verify:

- [ ] Token type: Granular Access Token
- [ ] Token name: Descriptive (e.g., `better-next-app-ci`)
- [ ] Expiration: Set (90 days or 1 year)
- [ ] Packages permission: Read and write
- [ ] Organizations permission: No access
- [ ] Token copied and saved securely
- [ ] GitHub secret created with name `NPM_TOKEN`
- [ ] GitHub secret contains full token (starts with `npm_`)

## Testing Your Setup

Run this command to verify everything is configured:

```bash
task npm:setup
```

**Expected output:**
```
✓ npm is installed
✓ Logged in as: your-username
✓ Package name 'create-better-next-app' is available

==========================================
✅ Setup Complete!

Next steps:

1. Add NPM_TOKEN to GitHub Secrets:
   - Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
   - Create secret: NPM_TOKEN
   - Get token from: https://www.npmjs.com/settings/your-username/tokens

2. Create your first release:
   git tag -a v0.1.0 -m 'Release v0.1.0'
   git push origin v0.1.0

3. Or publish manually now:
   task npm:publish
```

## Common Mistakes to Avoid

### ❌ Wrong Token Type

**Mistake:** Selecting "Classic Token"

**Fix:** Use "Granular Access Token" instead

### ❌ Wrong Permissions

**Mistake:** Selecting "Read-only" or "No access"

**Fix:** Select "Read and write" for Packages and scopes

### ❌ Wrong Secret Name

**Mistake:** Using `NPM_AUTH_TOKEN`, `npm_token`, or other variations

**Fix:** Use exactly `NPM_TOKEN` (case-sensitive)

### ❌ Token Not Copied

**Mistake:** Closing the page without copying the token

**Fix:** Generate a new token (you can't retrieve the old one)

### ❌ Token Expired

**Mistake:** Token expires and publishing fails

**Fix:** Generate a new token and update GitHub secret

## Troubleshooting

### "Token not found" in GitHub Actions

**Cause:** Secret name is incorrect

**Fix:**
1. Go to repository secrets
2. Verify secret is named exactly `NPM_TOKEN`
3. If not, delete and recreate with correct name

### "Insufficient permissions" when publishing

**Cause:** Token doesn't have write permissions

**Fix:**
1. Generate a new token
2. Select "Read and write" for Packages and scopes
3. Update GitHub secret with new token

### "Token expired"

**Cause:** Token expiration date passed

**Fix:**
1. Generate a new token with longer expiration
2. Update GitHub secret with new token
3. Set calendar reminder for next rotation

### "Package not found" during publish

**Cause:** Package name not claimed yet

**Fix:**
```bash
task npm:publish
```

This claims the package name on NPM.

## Security Best Practices

1. **Never commit tokens** to your repository
2. **Use expiration dates** - Rotate tokens regularly
3. **Use minimal permissions** - Only "Read and write" for packages
4. **Store securely** - Use password manager for backup
5. **Rotate regularly** - Set reminders to update tokens
6. **Monitor usage** - Check NPM audit logs periodically
7. **Revoke if compromised** - Delete token immediately if exposed

## Next Steps

After setting up your token:

1. ✅ Verify setup: `task npm:setup`
2. ✅ Test locally: `task npm:test`
3. ✅ Publish first version: `task npm:publish`
4. ✅ Create release: `git tag -a v0.1.0 -m "Release v0.1.0" && git push origin v0.1.0`
5. ✅ Verify automated publishing works

## Resources

- NPM Token Documentation: https://docs.npmjs.com/about-access-tokens
- GitHub Secrets Documentation: https://docs.github.com/en/actions/security-guides/encrypted-secrets
- Full Publishing Guide: [npm-publishing.md](./npm-publishing.md)
- Quick Start: [npm-quick-start.md](./npm-quick-start.md)
