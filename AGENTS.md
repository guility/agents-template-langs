# Lean component rules

Apply the general requirements from guility/agents_template/AGENTS.md. This package adds rules for Lean 4 projects that produce Rust through Lungo.

## Source and generation

- Keep proof-bearing source in lean/; pin an exact supported Lean release in lean/lean-toolchain. Do not use stable, nightly, or release-candidate pins.
- Keep lake-manifest.json and Cargo.lock committed. Pin lungo and lungo-build to the same exact version.
- Treat Rust emitted by Lungo as generated output. Do not edit it by hand or commit build output.
- Put the Rust API, host facilities, CLI, and adapters in src/. Keep GPU operations behind explicit Lean-declared host facilities.
- Give each exported Lean function at least one theorem claim checked by Lean's kernel; make the build require claims for every export.
- Keep sorry, extra axioms, unsafe, and partial out of exports. The build must deny extra axioms and unsafe dependencies.

## Proof boundary

- A Lean theorem proves a statement about its Lean definition. It does not prove that the generated Rust backend, Lungo runtime, Rust host implementation, CUDA driver, or GPU kernel implements that definition.
- Declare the generator, backend/runtime, linked libraries, Rust facilities, GPU kernels, driver, and hardware assumptions in the component's trust-boundary document.
- Do not describe an OCR inference runtime as fully verified until the code-generation/runtime bridge and every GPU kernel are covered by the accepted proof model.
- State tight resource bounds against an explicit cost model. Lungo build-worker resource limits are not runtime inference bounds.

## Tests and tools

- Run Lean proof checking and the Lungo build before Rust test suites. A failed theorem or missing required claim must fail the build.
- Keep separate unit, integration, mutation, and end-to-end suites. Mutation checks must alter Lean implementation while keeping its theorem unchanged; the kernel must reject each semantic mutant.
- Exercise every exported function through Rust tests and measure coverage for generated code and handwritten Rust with a toolchain that can report both accurately.
- Use cargo fmt --check, cargo clippy --locked --all-targets -- -D warnings, cargo audit, lake build, and the pinned mutation/coverage commands.
- Keep Lean, Rust, scripts, tests, and manually maintained configuration files within inherited 300-line maximum and 95%/100-line distribution limits.
