# Test Coverage Analysis

Analyze the codebase for test coverage gaps and suggest improvements.

## Priority Order
1. **Parser functions** — pure functions, highest value tests
2. **Command execution** — interface-based, mock injection
3. **TUI update logic** — message handling without rendering
4. **Style mapping** — trivial but ensures correctness

## Rules
- Use table-driven tests for all parser functions
- Mock external commands via the Runner interface
- Skip platform-specific tests on unsupported OS
- Keep test data inline (hardcoded strings, not files)
- Test error paths: malformed input, empty output, command failures
- Test edge cases: IPv6 addresses, missing fields, unusual states

## Coverage Targets
- `internal/netstat/parser.go` — 90%+
- `internal/netstat/ss.go` — basic execution only
- `internal/tui/update.go` — message handling paths
- Overall project — 70%+

## Output
For each gap found:
1. File and function name
2. What's not covered
3. Suggested test case (with example input/output)
