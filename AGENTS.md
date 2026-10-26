# C++ component rules

Apply the general requirements from `guility/agents_template/AGENTS.md`. This file contains only rules for C++ components.

## Implementation

- Use C++17 or a newer standard supported by the pinned compiler; set it explicitly in CMake.
- Use CMake and CTest; keep build configuration reproducible and compiler warnings enabled.
- Prefer RAII, value semantics, explicit ownership, const-correct interfaces, and standard-library types. Avoid owning raw pointers and unchecked casts.
- Document public headers and failure behavior. Keep domain code independent of frameworks and device/runtime APIs.
- Format with pinned `clang-format`; run `clang-tidy` and `cppcheck` with reviewed, versioned configuration.
- Keep source, tests, scripts, and manually maintained configuration at or below 300 lines; follow inherited file-count and 100-line distribution limits.

## Tests

- Use CTest with non-empty unit, integration, mutation, and end-to-end suites where applicable.
- Cover all unit line and branch paths; integration tests exercise each documented adapter boundary; end-to-end tests prove acceptance behavior.
- Each unit/integration test function has at least one companion fault-injection test. Test every production function with at least two applicable mutants of different classes; require a 100% mutation score.
- Keep test data deterministic, local, and non-sensitive. Do not silently skip a required suite or hardware gate.
