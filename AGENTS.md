# Go component rules

Follow the shared policy in [`agents-template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md). This package contains only Go-specific rules.

## Implementation

- Commit both `go.mod` and `go.sum`. The supported minimum Go version is recorded in `go.mod`; CI pins its toolchain.
- Handle errors explicitly. Keep interfaces small, use Go's exported-name conventions, and document exported APIs.
- Name tests `Test{Function}_{Scenario}`. Format with `gofmt`; use golangci-lint for linting and Lizard for CCN (maximum 10) and shared metrics.

## Push and MR checks

- Every push runs gofmt, golangci-lint, Gitleaks on outgoing commits, and Lizard CCN on changed files. The measured per-tool times are in the shared [fast-check table](https://github.com/guility/agents_template/blob/main/docs/quality/fast-check-benchmark.md).
- MR checks add dependency audit, full SAST, Dryer, build, and unit tests as defined by the shared pipeline.

## Tests and mutation

- Use Go's `testing` package and table-driven tests. Use Testcontainers and BDD/e2e tools only where the component has those boundaries.
- Keep 100% coverage and mutation gates from the shared policy. Run Gremlins v0.6.0 and Uncle Bob Mutator after merge; the shared locked quality toolchain runs Dryer for this package.
- Publish line and branch coverage from `go test -cover`.
- Keep `template/deploy/Dockerfiles/Dockerfile` with `production` and inherited `tests` targets; compile production first and resolve any test-only modules in `tests`.
