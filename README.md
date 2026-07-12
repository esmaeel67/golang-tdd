# Testing

This project uses Go's built-in testing toolkit. Below are the available test commands for different scenarios.

## Test Commands

### Basic Testing

| Command | Description |
|---------|-------------|
| `go test ./...` | Run all tests in all subdirectories |
| `go test -v ./...` | Run all tests with verbose output |
| `go test -cover ./...` | Run all tests and display coverage report |
| `go test -run TestName ./...` | Run specific test by name (supports regex) |
| `go test -bench=. ./...` | Run all benchmarks |

### Advanced Testing

| Command | Description |
|---------|-------------|
| `go test -race ./...` | Run tests with race detection enabled |
| `go test -count=5 ./...` | Run tests N times (useful for flaky tests) |
| `go test -timeout 30s ./...` | Set a custom timeout for tests |
| `go test -short ./...` | Skip long-running tests (uses `-short` flag) |
| `go test -parallel 8 ./...` | Run tests in parallel with specified number of workers |

## Usage Examples

```bash
# Basic testing
go test ./...                    # Run all tests
go test -v ./...                 # Verbose output
go test -cover ./...             # Coverage report

# Targeted testing
go test -run TestLogin ./...     # Run tests matching "TestLogin"
go test -run "TestUser|TestAuth" # Run multiple test patterns
go test -bench=. ./...           # Run benchmarks

# Advanced scenarios
go test -race ./...              # Check for race conditions
go test -count=5 ./...           # Run tests 5 times
go test -timeout 30s ./...       # 30-second timeout
go test -short ./...             # Skip long tests
go test -parallel 8 ./...        # Run with 8 parallel workers

# Combined examples
go test -v -race -cover ./...    # Verbose + race + coverage
go test -v -run TestAuth -count=3 ./... # Specific test, 3 times
