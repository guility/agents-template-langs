# Rust component rules

Follow the shared policy in [`agents-template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md). This package contains only Rust-specific rules.

## Implementation

- Commit `Cargo.lock` for applications and use the pinned Rust 1.100.0 toolchain in CI.
- Respect ownership and borrowing. Use `Result` and `Option` for explicit outcomes; avoid `unwrap()` in production code.
- Use PascalCase for types and traits, snake_case for functions and modules, and `test_{function}_{scenario}` for tests.
- Format with rustfmt and lint with Clippy. Use Lizard for CCN (maximum 10) and the shared complexity metrics.

## Push and MR checks

- Every push runs `cargo fmt --check`, Clippy, Gitleaks on outgoing commits, and Lizard CCN on changed files. The measured per-tool times are in the shared [fast-check table](https://github.com/guility/agents_template/blob/main/docs/quality/fast-check-benchmark.md).
- MR checks add dependency auditing, full SAST, Dryer, build, and unit tests as defined by the shared pipeline.

## Tests and mutation

- Use Rust's built-in test harness. Keep tests deterministic; use Testcontainers and browser tooling only for applicable integration and e2e cases.
- Keep 100% line-coverage and mutation gates from the shared policy. Run `cargo-mutants` 27.1.0 (Rust 1.88+) and Uncle Bob Mutator after merge; the shared locked quality toolchain runs Dryer for this package.
- Generate LCOV with `cargo-llvm-cov` 0.9.1 and fail below 100% line coverage. Report branch coverage as `n/a` while the upstream Rust branch reporter emits no branch counters; the shared [DECISION-001](https://github.com/guility/agents-template/blob/main/docs/design/DECISION-001-rust-branch-coverage.md) records the limitation.
- Keep `template/deploy/Dockerfiles/Dockerfile` with `production` and inherited `tests` targets; compile production first and add test dependencies and `cargo-llvm-cov` only in `tests`.
