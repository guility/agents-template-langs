# Python component rules

Follow the shared policy in [agents-template/AGENTS.md](https://github.com/guility/agents-template/blob/main/AGENTS.md). This file contains only Python-specific rules.

## Implementation

- Use Python 3.14.x and `uv`; commit the generated `uv.lock`.
- Add type hints to public interfaces. Keep domain entities framework-independent, use frozen dataclasses for value objects, and manage resources with context managers.
- Use PascalCase for classes and snake_case for functions, modules, variables, and tests (`test_{unit}_{scenario}_{expected_result}`).
- Format and lint with Ruff. Use mypy for Python type checking.

## Tests and mutation

- Use pytest and pytest-cov. Keep fixtures deterministic and independent of deployed infrastructure.
- Use Mutmut and Uncle Bob Mutator for Python mutation checks.
- `pytest-cov` reports line and branch coverage. Keep `mutmut` configuration in `pyproject.toml` aligned with the unit test command.
- Add the locked Python test extra only in the package's `tests` image target.
