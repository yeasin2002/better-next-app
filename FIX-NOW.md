# Fix Your NPM Token NOW

## The Problem

Your NPM token is missing "Bypass 2FA for automation" setting. This is required as of December 9, 2025.

## Follow These Steps EXACTLY

### Step 1: Open NPM Token Page

Click this link: https://www.npmjs.com/settings/mdkawsarislam2002/tokens

### Step 2: Generate New Token

1. Click the green **"Generate New Token"** button
2. Select **"Granular Access Token"**

### Step 3: Fill in These EXACT Settings

Copy and paste these values:

**Token name:**
```
better-next-app-ci-v2
```

**Description:**
```
GitHub Actions automation with 2FA bypass
```

**Expiration:**
- Select: **90 days**

### Step 4: Configure Permissions

Scroll down to **"Packages and scopes"** section:

**Permissions:**
- Select: **● Read and write**

**⚠️ CRITICAL - Check this box:**
```
☑ Bypass 2FA for automation
```

**THIS IS THE MOST IMPORTANT STEP!** If you don't check this box, publishing will fail.

**Organizations:**
- Select: **○ No access**

### Step 5: Generate Token

1. Scroll to bottom
2. Click **"Generate token"** button
3. **IMMEDIATELY COPY THE TOKEN** (starts with `npm_...`)
4. Save it somewhere safe (you won't see it again!)

### Step 6: Update GitHub Secret

1. Go to: https://github.com/yeasin2002/better-next-app/settings/secrets/actions
2. Find **NPM_TOKEN** in the list
3. Click **"Update"** button
4. Paste your new token
5. Click **"Update secret"**

### Step 7: Try Publishing Again

Open your terminal and run:

```bash
task npm:publish
```

## Expected Success Output

You should see:

```
npm notice Publishing to https://registry.npmjs.org/
+ create-better-next-app@0.0.2
✅ Published to NPM successfully!
```

## If It Still Fails

1. Double-check you enabled "Bypass 2FA for automation"
2. Make sure you updated the GitHub secret (not created a new one)
3. Wait 1 minute for GitHub to refresh the secret
4. Try again: `task npm:publish`

## Visual Checklist

When creating the token, you should see:

```
┌─────────────────────────────────────────┐
│  Packages and scopes                     │
│                                          │
│  Permissions:                            │
│  ● Read and write                        │
│                                          │
│  ☑ Bypass 2FA for automation  ← CHECK!  │
│                                          │
└─────────────────────────────────────────┘
```

## After Publishing Successfully

1. Verify at: https://www.npmjs.com/package/create-better-next-app
2. Test installation: `npx create-better-next-app@latest --help`
3. Set calendar reminder for 85 days (to rotate token)

## Need Help?

See detailed guide: [docs/npm-token-guide.md](./docs/npm-token-guide.md)
