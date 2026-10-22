# Python component rules

Apply the general requirements from `guility/agents_template/AGENTS.md`. This file contains only rules specific to Python components.

## Implementation

- Use Python 3.9 or newer and manage dependencies with `uv`; commit `uv.lock`.
- Add type hints to public interfaces. Keep domain entities independent of frameworks, use frozen dataclasses for immutable value objects, and use context managers for resource lifetimes.
- Name classes and public types in PascalCase; use snake_case for functions, methods, modules, and variables. Name tests `test_{unit}_{scenario}_{expected_result}`.
- Follow PEP 8. Document public APIs with docstrings. Represent expected failures with exceptions rather than error-code conventions.
- Format with `black` and `isort`; lint with `ruff`, `flake8`, and `mypy`. Check complexity with `lizard` and `radon`; scan with `bandit` and `safety`.

## Tests

- Use `pytest` or `unittest`; prefer `pytest` fixtures and parametrization and use Arrange-Act-Assert.
- Use `pytest-mock` or `unittest.mock` for test doubles. Use `pytest` with Testcontainers for integration tests and Playwright or Selenium for e2e tests where applicable.
- Use `mutmut` or `cosmic-ray` for mutation testing; use `coverage.py` and `pytest-cov` for coverage reporting.
- Keep tests independent and reproducible. Meet the inherited 100% coverage and mutation gates.
