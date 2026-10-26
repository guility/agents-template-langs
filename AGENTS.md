# CUDA component rules

Apply the general requirements from `guility/agents_template/AGENTS.md` and the C++ rules in the `cpp` language package. This file contains only CUDA-specific rules.

## Toolchain and implementation

- Pin the CUDA Toolkit, CMake, host compiler, and GPU architecture in the component build image and release record. Select only a host compiler supported by that Toolkit; never bypass `nvcc` compatibility checks.
- Set `CMAKE_CUDA_ARCHITECTURES` explicitly for every build. Do not rely on compiler defaults or `native` in release builds.
- Keep host and device responsibilities explicit. Check CUDA API errors, use scoped ownership for device memory, synchronize deliberately, and document stream and device assumptions.
- Treat architecture-specific kernels as a separate compatibility surface; include at least one real-GPU smoke on every supported release architecture.
- Run Compute Sanitizer memory checks on GPU integration/e2e cases; fail on any sanitizer finding.

## Tests and coverage

- Provide non-empty unit, integration, mutation, and GPU end-to-end suites. GPU tests must fail when the required device is unavailable; do not silently skip them.
- Meet inherited 100% line/branch coverage and mutation gates. Every production function has at least two applicable mutants of different classes, and each unit/integration test function has a companion fault-injection test.
- NVIDIA documents Bullseye coverage for host functions only; it does not cover device functions. Do not report full CUDA coverage from host-only instrumentation. Before release, provide an approved device-code coverage method or documented owner decision satisfying the inherited gate; GPU smoke, mutation, and sanitizers remain required supplemental evidence.
