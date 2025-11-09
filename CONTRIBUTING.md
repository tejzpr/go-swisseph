# Contributing to Go Swiss Ephemeris

Thank you for your interest in contributing to Go Swiss Ephemeris! This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/tejzpr/go-swisseph.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes: `make test`
6. Commit your changes: `git commit -am 'Add some feature'`
7. Push to the branch: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.21 or later
- A C compiler (gcc, clang, or MSVC)
- Git (for submodule management)
- Make (optional, but recommended)

### Quick Setup

```bash
# Clone and setup
git clone https://github.com/tejzpr/go-swisseph.git
cd go-swisseph

# Run the setup script (recommended)
./setup.sh

# Or manually:
git submodule update --init --recursive
go test -v
```

### Building

```bash
# Build the library
make build

# Run tests
make test

# Run tests with coverage
make test-coverage

# Build examples
make examples
```

### Important: Swiss Ephemeris Submodule

This project uses the Swiss Ephemeris C library as a Git submodule. Always ensure the submodule is initialized:

```bash
git submodule update --init --recursive
```

If you get compilation errors about missing headers, the submodule likely isn't initialized.

## Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and concise

## Testing

- Write tests for new functionality
- Ensure all tests pass before submitting a PR
- Aim for good test coverage
- Include both unit tests and integration tests where appropriate

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

## Documentation

- Update README.md if you add new features
- Add examples for new functionality
- Document all exported functions and types
- Use clear and concise language

## Commit Messages

- Use clear and descriptive commit messages
- Start with a verb in present tense (e.g., "Add", "Fix", "Update")
- Keep the first line under 50 characters
- Add detailed description if necessary

Example:
```
Add support for planetary moons

- Implement calculation functions for planetary satellites
- Add constants for moon IDs
- Update documentation with examples
```

## Pull Request Guidelines

1. **One feature per PR** - Keep PRs focused on a single feature or fix
2. **Update tests** - Add or update tests for your changes
3. **Update documentation** - Update README and add examples if needed
4. **Clean commit history** - Squash commits if necessary
5. **Pass all tests** - Ensure all tests pass before submitting
6. **Follow code style** - Run `go fmt` and follow Go conventions

## Bug Reports

When reporting bugs, please include:

1. Go version
2. Operating system
3. Steps to reproduce
4. Expected behavior
5. Actual behavior
6. Error messages or stack traces
7. Minimal code example that demonstrates the issue

## Feature Requests

When requesting features:

1. Describe the feature and its use case
2. Explain why it would be useful
3. Provide examples of how it would be used
4. Consider if it aligns with the project's goals

## Code Review Process

1. Maintainers will review your PR
2. Address any feedback or requested changes
3. Once approved, your PR will be merged
4. Your contribution will be acknowledged in the release notes

## License

By contributing, you agree that your contributions will be licensed under the same dual license as the project (AGPL-3.0-or-later OR LGPL-3.0-or-later with professional license).

## Questions?

If you have questions, feel free to:

- Open an issue for discussion
- Check existing issues and pull requests
- Review the documentation

## Thank You!

Your contributions help make this project better for everyone. Thank you for taking the time to contribute!

