#!/bin/bash
# Setup script for go-swisseph

set -e

echo "Setting up Go Swiss Ephemeris..."
echo ""

# Check if go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="1.21"
if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then 
    echo "Warning: Go version $GO_VERSION found. Go 1.21 or later is recommended."
fi

# Verify C source files are present
if [ ! -f swisseph/swephexp.h ]; then
    echo "Error: Swiss Ephemeris C source files not found in swisseph/ directory"
    echo "  Please ensure you have cloned the complete repository."
    exit 1
fi
echo "✓ Swiss Ephemeris C source files found"

# Check for C compiler
if command -v gcc &> /dev/null; then
    echo "✓ C compiler found: gcc"
elif command -v clang &> /dev/null; then
    echo "✓ C compiler found: clang"
else
    echo "Warning: No C compiler found. Please install gcc or clang."
    echo "  macOS: xcode-select --install"
    echo "  Linux: sudo apt-get install build-essential"
    exit 1
fi

# Run tests
echo ""
echo "Running tests..."
if go test -run TestVersion > /dev/null 2>&1; then
    echo "✓ Tests passed"
else
    echo "✗ Tests failed. Please check your setup."
    exit 1
fi

echo ""
echo "========================================="
echo "Setup complete! 🎉"
echo "========================================="
echo ""
echo "Next steps:"
echo "  1. Run tests: make test"
echo "  2. Build examples: make examples"
echo "  3. Read the docs: cat README.md"
echo ""
echo "For ephemeris files, see: README.md (Ephemeris Files section)"
echo ""
