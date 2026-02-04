# NPM Documentation Summary

All NPM-related documentation has been consolidated into comprehensive guides.

## 📚 Available Documentation

### 1. Complete Guide (Recommended)
**File:** [npm-complete-guide.md](./npm-complete-guide.md)

**What it covers:**
- ✅ Quick start (5 minutes)
- ✅ Detailed setup instructions
- ✅ Token generation guide
- ✅ Release process (automated & manual)
- ✅ Troubleshooting all common issues
- ✅ Maintenance and best practices
- ✅ Testing procedures
- ✅ Security guidelines

**Use this when:** You need comprehensive information about NPM publishing.

### 2. Token Generation Guide
**File:** [npm-token-guide.md](./npm-token-guide.md)

**What it covers:**
- ✅ Visual step-by-step token creation
- ✅ NPM security updates (December 2025)
- ✅ "Bypass 2FA" requirement explanation
- ✅ Token types comparison
- ✅ Common mistakes to avoid

**Use this when:** You need to generate or regenerate your NPM token.

## 🚀 Quick Reference

### First-Time Setup

```bash
# 1. Login to NPM
npm login

# 2. Verify setup
task npm:setup

# 3. Generate token at:
# https://www.npmjs.com/settings/YOUR_USERNAME/tokens
# - Enable "Bypass 2FA for automation"
# - Set 90 days expiration

# 4. Add token to GitHub Secrets:
# https://github.com/yeasin2002/better-next-app/settings/secrets/actions

# 5. Create release
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

### Common Tasks

```bash
task npm:setup      # Verify NPM setup
task npm:test       # Test package locally
task npm:pack       # Preview package contents
task npm:version    # Sync version with git tag
task npm:publish    # Publish manually (fallback)
```

### Important URLs

- **NPM Package:** https://www.npmjs.com/package/create-better-next-app
- **GitHub Actions:** https://github.com/yeasin2002/better-next-app/actions
- **GitHub Releases:** https://github.com/yeasin2002/better-next-app/releases
- **NPM Tokens:** https://www.npmjs.com/settings/YOUR_USERNAME/tokens
- **GitHub Secrets:** https://github.com/yeasin2002/better-next-app/settings/secrets/actions

## 📝 What Changed

### Consolidated Files

The following files have been merged into `npm-complete-guide.md`:

- ❌ `npm-publishing.md` (deleted)
- ❌ `npm-quick-start.md` (deleted)
- ❌ `TROUBLESHOOTING-NPM.md` (deleted)
- ❌ `RELEASE-STATUS.md` (deleted)
- ❌ `PUBLISH-NOW.md` (deleted)
- ❌ `QUICK-FIX.md` (deleted)
- ❌ `FIX-NOW.md` (deleted)

### Kept Files

- ✅ `npm-complete-guide.md` - Comprehensive guide (NEW)
- ✅ `npm-token-guide.md` - Visual token generation guide

## 🎯 When to Use Each Guide

### Use npm-complete-guide.md when:
- Setting up NPM publishing for the first time
- Need to understand the full workflow
- Troubleshooting any NPM-related issues
- Learning about release process
- Need security best practices

### Use npm-token-guide.md when:
- Generating a new NPM token
- Token expired and needs rotation
- Confused about token settings
- Need visual step-by-step instructions
- Want to understand "Bypass 2FA" requirement

## ⚠️ Critical Information

### NPM Security Update (December 9, 2025)

- ❌ Classic tokens permanently revoked
- ✅ Granular Access Tokens required
- 🔒 Must enable "Bypass 2FA for automation"
- ⏳ Write tokens limited to 90 days maximum

### Token Requirements

For CI/CD automation, your token MUST have:
1. **Type:** Granular Access Token
2. **Packages permission:** Read and write
3. **Bypass 2FA:** Enabled ✅ (CRITICAL!)
4. **Expiration:** 90 days (maximum)
5. **Organizations:** No access

## 🔄 Release Workflow

### Automated (Recommended)

```bash
# Make changes
git add .
git commit -m "feat: your changes"
git push origin main

# Create release
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

# GitHub Actions automatically:
# - Builds binaries
# - Creates GitHub release
# - Publishes to NPM
```

### Manual (Fallback)

Only use if GitHub Actions fails. See complete guide for instructions.

## 📊 Monitoring

After release, check:

1. **GitHub Actions:** https://github.com/yeasin2002/better-next-app/actions
2. **GitHub Release:** https://github.com/yeasin2002/better-next-app/releases
3. **NPM Package:** https://www.npmjs.com/package/create-better-next-app
4. **Test Installation:** `npx create-better-next-app@latest --help`

## 🆘 Getting Help

1. Check [npm-complete-guide.md](./npm-complete-guide.md) troubleshooting section
2. Review [npm-token-guide.md](./npm-token-guide.md) for token issues
3. Check GitHub Actions logs for workflow errors
4. Verify token settings on NPM website
5. Open an issue on GitHub

## 📅 Maintenance Reminders

- **Day 85:** Rotate NPM token (before 90-day expiration)
- **After each release:** Verify NPM package updated
- **Monthly:** Check download statistics
- **Quarterly:** Review and update documentation

---

**Last Updated:** February 4, 2026
**Documentation Version:** 2.0 (Consolidated)
