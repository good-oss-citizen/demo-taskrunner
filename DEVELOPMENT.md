# Development Guide

## Prerequisites

- Go 1.22+
- golangci-lint v1.57+

## Setup

```bash
go mod download
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
```

## Running Tests

```bash
go test ./...
```

## Running Linter

```bash
golangci-lint run
```

## Test Conventions

- Test files live in `tests/` (not alongside source files)
- Test function names: `Test<FunctionName>_<Scenario>` (e.g., `TestParse_EmptyFile`)
- Use table-driven tests for functions with multiple input/output cases
- Error messages should include the actual value: `t.Errorf("got %v, want %v", got, want)`
