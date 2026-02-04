# Publish to NPM NOW

## The Issue

When you run `task npm:publish` locally, it uses your session token (from `npm login`), which requires 2FA. The Granular Access Token you created is only in GitHub CI/CD.

## Solution: Use Your Token Locally

### Step 1: Get Your NPM Token

You need to copy the same token you added to GitHub secrets.

**If you still have it saved:**

- Use that token (starts with `npm_...`)

**If you don't have it:**

1. Go to: https://www.npmjs.com/settings/mdkawsarislam2002/tokens
2. Find your token: `better-next-app-ci-v2`
3. You can't view it again, so you need to generate a new one
4. Follow the same steps from FIX-NOW.md
5. Copy the new token

### Step 2: Set Token in Your Terminal

**For PowerShell (Windows):**

```powershell
$env:NPM_TOKEN="npm_YOUR_TOKEN_HERE"
```

Replace `npm_YOUR_TOKEN_HERE` with your actual token.

### Step 3: Configure NPM to Use Token

Run this command:

```powershell
npm config set //registry.npmjs.org/:_authToken $env:NPM_TOKEN
```

### Step 4: Publish

```bash
task npm:publish
```

### Step 5: Clean Up (After Publishing)

Remove the token from your local config:

```bash
npm config delete //registry.npmjs.org/:_authToken
```

## Alternative: Publish via GitHub Actions

Instead of publishing locally, push a git tag and let GitHub Actions publish automatically:

```bash
# Commit any changes first
git add .
git commit -m "chore: prepare for release"

# Create and push tag
git tag -a v0.0.3 -m "Release v0.0.3"
git push origin v0.0.3
```

GitHub Actions will automatically:

1. Build binaries with GoReleaser
2. Create GitHub release
3. Publish to NPM using the token in secrets

## Which Option Should You Choose?

### Use Local Publishing If:

- ✅ You want to publish right now
- ✅ You have the token available
- ✅ You want to test before creating a release

### Use GitHub Actions If:

- ✅ You want automated publishing
- ✅ You want to create a proper release
- ✅ You don't want to handle tokens locally

## Recommended: Use GitHub Actions

This is safer and more automated:

```bash
# Make sure everything is committed
git status

# Create release tag
git tag -a v0.0.3 -m "Release v0.0.3"
git push origin v0.0.3

# Wait 2-3 minutes for GitHub Actions to complete
# Check: https://github.com/yeasin2002/better-next-app/actions
```

Then verify at: https://www.npmjs.com/package/create-better-next-app
