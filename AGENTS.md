# Go component rules

Follow the shared policy in [agents-template/AGENTS.md](https://github.com/guility/agents-template/blob/main/AGENTS.md). This file contains only Go-specific rules.

## Implementation

- Commit both `go.mod` and `go.sum`. The supported minimum Go version is recorded in `go.mod`.
- Handle errors explicitly. Keep interfaces small, use Go's exported-name conventions, and document exported APIs.
- Name tests `Test{Function}_{Scenario}`. Format with `gofmt` and lint with golangci-lint.

## Tests and mutation

- Use Go's `testing` package and table-driven tests.
- Use Gremlins v0.6.0 and Uncle Bob Mutator for Go mutation checks.
- Go's built-in cover profile measures statements, not physical line or branch coverage. Report line and branch as `n/a`; [DECISION-002](https://github.com/guility/agents-template/blob/main/docs/design/DECISION-002-go-coverage.md) records the limitation.
- Resolve test-only Go modules only in the package's `tests` image target.
