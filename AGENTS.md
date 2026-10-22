# Rust component rules

Apply the general requirements from `guility/agents_template/AGENTS.md`. This file contains only rules specific to Rust components.

## Implementation

- Commit `Cargo.lock` for application components.
- Respect ownership and borrowing. Use `Result` and `Option` for explicit outcomes; avoid `unwrap()` in production code.
- Name types and traits in PascalCase; use snake_case for functions, methods, modules, and variables. Name tests `test_{function}_{scenario}`.
- Follow Rust API Guidelines and document public APIs with rustdoc. Keep module tests alongside the code where appropriate.
- Format with `rustfmt` and lint with `clippy`. Scan dependencies with `cargo-audit` and check complexity with the selected analyzer.

## Tests

- Use Rust's built-in `#[test]`. Return `Result<T, E>` from tests that exercise fallible paths.
- Use `mockall` for test doubles when needed. Use Testcontainers modules for integration tests, `fantoccini` for applicable browser e2e tests, and `cucumber-rust` for BDD.
- Use `cargo-mutants` for mutation testing and `cargo-tarpaulin` for coverage reporting. Keep tests independent and reproducible, and meet the inherited 100% coverage and mutation gates.
