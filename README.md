# Lean-to-Rust component package

This package defines proof-bearing logic in Lean 4 and exposes it to Rust through Lungo. The starter pins Lean 4.35.0 and Lungo 1.120.164, whose worker adapter supports that Lean release.

The generated Rust is not a verified compiler result: Lungo documents differential checks against Lean's native backend, while listing its adapter, generators, and runtime in the trusted base. Rust host functions and GPU kernels are outside Lean's proof. See the package README and retain the project's full proof-boundary decision before using this as a production OCR runtime.

Fetch the immutable language package from lean/v1.0.0 after review and publication. The package is intended for a Lean proof component consumed by a Rust application; generated build output stays under Cargo's target directory.
