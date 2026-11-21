# Python component rules

Follow the shared policy in [`agents-template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md). This package contains only Python-specific rules.

## Implementation

- Use Python 3.14.x and `uv`; commit the generated `uv.lock`.
- Add type hints to public interfaces. Keep domain entities framework-independent, use frozen dataclasses for value objects, and manage resources with context managers.
- Use PascalCase for classes and snake_case for functions, modules, variables, and tests (`test_{unit}_{scenario}_{expected_result}`).
- Format and lint with Ruff. Run mypy during the MR gate. Use Lizard for CCN (maximum 10) and the shared complexity metrics.

## Push and MR checks

- Every push runs Ruff check, Ruff format check, Gitleaks on outgoing commits, and Lizard CCN on changed files. The measured per-tool times are in the shared [fast-check table](https://github.com/guility/agents_template/blob/main/docs/quality/fast-check-benchmark.md).
- MR checks add mypy, full SAST, dependency scanning, Dryer, build, and unit tests as defined by the shared pipeline.

## Tests and mutation

- Use pytest; keep unit, integration, e2e, and mutation tests independent and reproducible. Use Testcontainers for applicable integration tests.
- Keep 100% coverage and mutation gates from the shared policy. Run native Mutmut and Uncle Bob Mutator after merge; the shared locked quality toolchain runs Dryer for this package.
- `pytest-cov` reports line and branch coverage. Keep `mutmut` configuration in `pyproject.toml` aligned with the unit test command.
- Keep `template/deploy/Dockerfiles/Dockerfile` with a `production` target and a later `tests` target that inherits from it and adds the locked unit-test extra.
