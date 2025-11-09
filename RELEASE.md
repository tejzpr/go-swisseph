# Release Guide

This document describes how to release a new version of the Go Swiss Ephemeris package.

## Versioning

The package uses [Semantic Versioning](https://semver.org/):
- **MAJOR.MINOR.PATCH** (e.g., v1.2.3)
- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

The current version is defined in:
- `swisseph.go`: `PackageVersion` constant
- `go.mod`: Version comment
- `README.md`: Version badge/note

## Release Steps

### 1. Update Version

Update the version in the following files:

**swisseph.go:**
```go
const PackageVersion = "v1.0.0"  // Update to new version
```

**go.mod:**
```go
// Version: v1.0.0  // Update to new version
```

**README.md:**
```markdown
**Version:** v1.0.0  // Update to new version
```

### 2. Update CHANGELOG.md

Add a new section documenting the changes in this release.

### 3. Commit Changes

```bash
git add swisseph.go go.mod README.md CHANGELOG.md
git commit -m "Bump version to v1.0.0"
```

### 4. Create Git Tag

```bash
git tag -a v1.0.0 -m "Release v1.0.0"
```

### 5. Push to GitHub

```bash
git push origin main
git push origin v1.0.0
```

### 6. Verify on pkg.go.dev

After pushing, the package will automatically appear on [pkg.go.dev](https://pkg.go.dev/github.com/tejzpr/go-swisseph) within a few minutes.

Users can then install the specific version:
```bash
go get github.com/tejzpr/go-swisseph@v1.0.0
```

## Pre-Release Checklist

- [ ] All tests pass: `go test ./...`
- [ ] Examples build: `make examples`
- [ ] Version updated in all files
- [ ] CHANGELOG.md updated
- [ ] README.md updated if needed
- [ ] Git tag created and pushed
- [ ] Package appears on pkg.go.dev

## Current Version

**v1.0.0** - Initial stable release

