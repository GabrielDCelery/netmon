# Claude Code Workflow Rules

## Branching

- **Never commit directly to `main`**. Always create a feature branch.
- Branch naming: `<type>/<short-description>` (e.g., `feat/ss-table-view`, `fix/parser-nil-check`)
- Open a PR via `gh pr create` when work is ready for review.

## Commits

- Use [Conventional Commits](https://www.conventionalcommits.org/): `feat:`, `fix:`, `refactor:`, `test:`, `chore:`, `docs:`
- Keep commits atomic — one logical change per commit.

## Quality Gates

Before pushing or creating a PR, always run:

```sh
mise run test
mise run lint
```

Both must pass with zero errors.

## Off-Limits Files

Do **not** modify these files without explicit user approval:

- `CLAUDE.md` (this file)
- `.github/workflows/ci.yml`
- `mise.toml` (task/tool config)
- Any file containing secrets or credentials

## Code Style

- Follow `gofmt` and `goimports` formatting.
- Use `internal/` packages — this is a CLI app, not a library.
- Prefer interfaces for external command execution (testability).
- Table-driven tests for parsers and pure functions.

## Dependencies

- Run `go mod tidy` after adding or removing imports.
- Keep dependencies minimal — justify new additions.

## Project Structure

```
cmd/netmon/     — CLI entry point
internal/tui/      — Bubbletea model, update, view
internal/netstat/  — ss command execution and parsing
internal/styles/   — Lipgloss style definitions
docs/              — Architecture and testing docs
```
