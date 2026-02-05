# Testing Strategy

## Test Categories

### Unit Tests (primary focus)
- **Parser tests** (`internal/netstat/parser_test.go`): Table-driven tests covering all ss output formats. Target: 90%+ coverage.
- **Style tests**: Verify correct style mapping for each connection state.

### Integration Tests
- **SS runner test** (`internal/netstat/ss_test.go`): Verifies ss command execution. Skipped on non-Linux platforms.

### Manual Testing
- Run `sudo mise run run` and verify the TUI renders correctly with live data.

## Coverage

Run coverage report:
```sh
mise run test       # Prints coverage to stdout
mise run coverage   # Generates coverage.html
```

Coverage expectations:
- `internal/netstat/parser.go` — 90%+ (pure function, table-driven tests)
- `internal/netstat/ss.go` — basic execution test only
- `internal/tui/` — not unit tested initially (UI rendering)
- `internal/styles/` — trivial, low priority

## Testing Patterns

### Table-driven tests
All parser tests use the Go table-driven pattern:
```go
tests := []struct {
    name     string
    input    string
    expected []Connection
}{...}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { ... })
}
```

### Mock injection
The `Runner` interface in `ss.go` enables test doubles:
```go
type mockRunner struct {
    output string
    err    error
}
func (m *mockRunner) Run(ctx context.Context) (string, error) {
    return m.output, m.err
}
```

## Adding Tests

When adding new functionality:
1. Write parser tests first (pure functions are easiest to test)
2. Use interfaces for any external command or I/O
3. Skip platform-specific tests with `runtime.GOOS` checks
4. Keep test data as hardcoded strings, not external files
