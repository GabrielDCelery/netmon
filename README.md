# diagnoose

Terminal UI for Linux networking diagnostics. Displays live network connection data from `ss` in an interactive table with color-coded states and auto-refresh.

## Quick Start

```sh
git clone https://github.com/GabrielDCelery/diagnoose.git
cd diagnoose
mise run setup
mise run build
sudo ./bin/diagnoose
```

> `sudo` is needed to see process information from `ss -tunap`.

## Architecture

- **Language/Framework**: Go 1.24 + [Bubbletea](https://github.com/charmbracelet/bubbletea) (TUI) + [Lipgloss](https://github.com/charmbracelet/lipgloss) (styling)
- **Database**: None
- **Hosting**: Local CLI tool
- **External APIs**: None — reads from `ss` system command

### Package Structure

```
cmd/diagnoose/     Entry point
internal/tui/      Bubbletea model, update, view
internal/netstat/  ss command execution and output parsing
internal/styles/   Lipgloss style definitions
```

## Configuration

No configuration required. The tool reads directly from the `ss` command.

## Development

```sh
mise run test      # Run tests with race detector and coverage
mise run lint      # Run golangci-lint
mise run fmt       # Format code with gofmt
mise run vet       # Run go vet
mise run build     # Build binary to bin/diagnoose
mise run coverage  # Generate HTML coverage report
```

## Deployment

Local CLI binary — no deployment infrastructure. Build with `mise run build` and distribute the binary.
