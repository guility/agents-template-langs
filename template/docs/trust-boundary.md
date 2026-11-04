# Trust boundary

This file is part of the component contract. Update it whenever an export, host facility, backend, or runtime dependency changes.

## Proof-bearing source

- Lean version: 4.35.0, pinned by lean/lean-toolchain.
- Theorems: list each exported definition and the claim required by build.rs.
- Kernel assumptions: list any axioms, unsafe declarations, or external constants. The starter build rejects extra axioms and unsafe code.
- Cost model: define input sizes, output sizes, allocations, copies, and primitive-operation costs before stating time or memory bounds.

## Trusted computing base

- Lungo version and source revision: 1.120.164; see lean/lake-manifest.json.
- Lungo compiler, bridge IR, Rust backend, generated runtime, and Rust standard library.
- Handwritten Rust host functions, FFI, linked libraries, and build scripts.
- For GPU workloads: kernel source/compiler, device runtime, driver, firmware, and hardware.
- Toolchain and build environment used to produce the binary.

## Validation evidence

- List proof-checking and differential-test commands with exact versions.
- List unit, integration, mutation, and end-to-end coverage, including whether generated code is measurable.
- List benchmark hardware and methodology for observed time and memory. Measurements are not formal bounds.
- State which claims are proofs, which are tests, and which depend on annotations or external components.
