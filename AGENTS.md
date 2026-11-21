# Node.js and TypeScript component rules

Follow the shared policy in [agents-template/AGENTS.md](https://github.com/guility/agents-template/blob/main/AGENTS.md). This file contains only Node.js and TypeScript-specific rules.

## Implementation

- Use Node.js 22.x and pnpm; commit `pnpm-lock.yaml` and use strict TypeScript.
- Prefer ES modules, async/await, and immutable updates. Avoid `any`.
- Use PascalCase for types and classes, camelCase for values, and behavior-focused test names.
- Use Prettier for formatting, ESLint for linting, and `tsc` for type checking.

## Tests and mutation

- Use Jest and keep test fixtures deterministic.
- Use Stryker and Uncle Bob Mutator for JavaScript and TypeScript mutation checks.
- Publish line and branch coverage from Jest. The `test:mutation` script runs mutation tests; Stryker mutates production code separately.
- Restore production dependencies first and add locked test dependencies only in the package's `tests` image target.
