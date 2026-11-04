# Lean-to-Rust component starter

This example defines a rectangle and its area in Lean, proves a theorem about the exported function, and builds a Rust library and CLI with Lungo. Keep Cargo and Lake lockfiles committed. Lean source is authoritative; generated Rust lives under target/.

## Container workflow

~~~powershell
docker build -t lean-rust-component .
docker run --rm lean-rust-component
docker run --rm -v "$((Get-Location).Path):/workspace" -w /workspace lean-rust-component cargo test --locked
~~~

For an interactive shell, override the image command with bash. The image pins Rust 1.89.0 and installs Lean 4.35.0.

## Required suites

~~~bash
cargo test --locked --test unit
cargo test --locked --test integration
cargo test --locked --test e2e
bash tests/mutation/run_mutants.sh
lake -d lean build
cargo fmt --check
cargo clippy --locked --all-targets -- -D warnings
~~~

The two mutation patches change Lean implementation while retaining its theorem; each must fail in Lean proof checking. Add two applicable mutants for every new production function. Measure coverage for both handwritten and generated Rust; do not accept a report that silently excludes generated code.

## Trust and proof limits

The build rejects missing claims, extra axioms, and unsafe exported code. A theorem applies to the Lean definition. Lungo's Rust backend/runtime, generated code, host Rust implementations, and GPU kernels are not covered by that theorem; Lungo's differential tests are not a proof of compiler correctness. Record project-specific assumptions and trusted components in [docs/trust-boundary.md](docs/trust-boundary.md). This starter is an interop candidate, not evidence that a full OCR inference runtime or its GPU kernels are formally verified.
