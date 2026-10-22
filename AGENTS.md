# Node.js and TypeScript component rules

Apply the general requirements from `guility/agents_template/AGENTS.md`. This file contains only rules specific to Node.js and TypeScript components.

## Implementation

- Use Node.js 18 or newer and pnpm 8 or newer; commit `pnpm-lock.yaml`.
- Enable strict TypeScript checking and avoid `any`. Prefer ES module imports, async/await, immutable data updates, and destructuring when it improves clarity.
- Format with Prettier; lint and type-check with ESLint and `tsc`. Check complexity with the tools selected for the project and scan dependencies with `npm audit` or Snyk.

## Tests

- Use Jest or Vitest. Name tests with nested `describe` and behavior-focused `it` descriptions.
- Use `jest.mock()` or `vi.mock()` for test doubles. Use Supertest and Testcontainers for applicable integration tests; Playwright or Cypress for e2e; Cucumber or Jest BDD for behavior tests.
- Use Stryker for mutation testing and the runner's built-in coverage reporting. Keep tests independent and reproducible, and meet the inherited 100% coverage and mutation gates.
