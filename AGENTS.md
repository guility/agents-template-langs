# Go component rules

Apply the general requirements from `guility/agents_template/AGENTS.md`. This file contains only rules specific to Go components.

## Implementation

- Commit both `go.mod` and `go.sum`.
- Handle errors explicitly. Keep interfaces small, use exported names with initial capitals, follow Effective Go, and document exported APIs in Go doc style.
- Use Go's MixedCaps naming (initial capital for exported identifiers, lower case for unexported identifiers). Name tests `Test{Function}_{Scenario}`.
- Format with `gofmt` and `goimports`; lint with `golangci-lint`. Check complexity with `lizard` or `gocyclo` and scan with `gosec`.

## Tests

- Use Go's built-in `testing` package and table-driven tests for related cases.
- Use `gomock` or `testify/mock` for test doubles. Use `testcontainers-go` for applicable integration tests and `godog` or `ginkgo` for e2e/BDD scenarios.
- Use a supported Go mutation tool and `go test -cover` for coverage. Keep tests independent and reproducible, and meet the inherited 100% coverage and mutation gates.
