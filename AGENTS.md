# Rust component rules

Follow the shared policy in [agents-template/AGENTS.md](https://github.com/guility/agents-template/blob/main/AGENTS.md). This file contains only Rust-specific rules.

## Implementation

- Commit `Cargo.lock` for applications and use the pinned Rust 1.100.0 toolchain.
- Respect ownership and borrowing. Use `Result` and `Option` for explicit outcomes; avoid `unwrap()` in production code.
- Use PascalCase for types and traits, snake_case for functions and modules, and `test_{function}_{scenario}` for tests.
- Format with rustfmt and lint with Clippy.

## Tests and mutation

- Use Rust's built-in test harness and keep test fixtures deterministic.
- Use `cargo-mutants` 27.1.0 (Rust 1.88+) and Uncle Bob Mutator for Rust mutation checks.
- Generate LCOV with `cargo-llvm-cov` 0.9.1. Report branch coverage as `n/a`; [DECISION-001](https://github.com/guility/agents-template/blob/main/docs/design/DECISION-001-rust-branch-coverage.md) records the missing branch counters.
- Add test dependencies and `cargo-llvm-cov` only in the package's `tests` image target.
