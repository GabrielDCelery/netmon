# Architecture

## Tech Stack

- **Go 1.24** — main language
- **Bubbletea** — terminal UI framework (Model-View-Update pattern)
- **Bubbles** — pre-built TUI components (table)
- **Lipgloss** — terminal styling (colors, borders, padding)

## Package Structure

```
cmd/diagnoose/        CLI entry point, creates tea.Program
internal/tui/         Bubbletea application
  model.go            Model struct, NewModel(), Init()
  update.go           Update() message handler, commands
  view.go             View() rendering
internal/netstat/     Network data layer
  types.go            Connection struct
  ss.go               Runner interface + SSRunner implementation
  parser.go           Parse() — raw ss output → []Connection
  parser_test.go      Table-driven tests for parser
  ss_test.go          Integration test for ss execution
internal/styles/      UI styling
  styles.go           Lipgloss style definitions, StyleForState()
```

## Data Flow

1. `Init()` fires `fetchConnections` command
2. `fetchConnections` calls `Runner.Run()` → raw ss output
3. `Parse()` converts raw output → `[]Connection`
4. `connectionsMsg` delivered to `Update()`
5. `Update()` converts connections to table rows, schedules `tick()`
6. `tick()` fires after 2s → triggers another `fetchConnections`
7. `View()` renders table + status bar + help text

## Design Decisions

### Interface-based command execution
`Runner` interface in `ss.go` allows injecting mock output in tests. Tests never execute real system commands.

### Parser as pure function
`Parse(string) []Connection` has no side effects. This is the primary test target with table-driven tests covering all connection states, IPv6, missing process info, and edge cases.

### internal/ over pkg/
This is a CLI application, not a library. Nothing in this project needs to be importable by external packages.

### Bubbletea MVU pattern
Follows the standard Model-View-Update architecture. Messages are the only way to update state, keeping the update logic testable and predictable.
