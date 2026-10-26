# C++ component starter

This starter uses a small domain function, an application service, and a command-line adapter. Replace the example with one bounded component while preserving the approved contracts and required test suites.

## Build

```bash
cmake -S . -B build -DCMAKE_BUILD_TYPE=Debug -DCPP_ENABLE_COVERAGE=ON
cmake --build build --parallel
```

## Required suites

```bash
ctest --test-dir build --output-on-failure -L unit
ctest --test-dir build --output-on-failure -L integration
ctest --test-dir build --output-on-failure -L e2e
ctest --test-dir build --output-on-failure -L mutation
```

Generate the combined line and branch report after unit, integration, and e2e suites with `uvx --from gcovr==8.7 gcovr --root . --object-directory build --exclude 'tests/.*' --fail-under-line 100 --fail-under-branch 100`. CI must fail below 100% or when a report is absent; CI images pin the same tool version. Pin all tool versions in the component CI image; do not fetch unpinned build dependencies during a job.
