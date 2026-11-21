# Node.js and TypeScript component rules

Follow the shared policy in [`agents-template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md). This package contains only Node.js and TypeScript-specific rules.

## Implementation

- Use Node.js 22.x and pnpm; commit `pnpm-lock.yaml` and use strict TypeScript.
- Prefer ES modules, async/await, and immutable updates. Avoid `any`.
- Use PascalCase for types and classes, camelCase for values, and behavior-focused test names.
- Use Prettier for formatting. Run ESLint and `tsc` in the MR gate. Use Lizard for CCN (maximum 10) and the shared complexity metrics.

## Push and MR checks

- Every push runs Prettier check, Gitleaks on outgoing commits, and Lizard CCN on changed files. The measured per-tool times are in the shared [fast-check table](https://github.com/guility/agents_template/blob/main/docs/quality/fast-check-benchmark.md).
- MR checks add ESLint, `tsc`, full SAST, dependency scanning, Dryer, build, and unit tests as defined by the shared pipeline.

## Tests and mutation

- Use Jest. Keep tests deterministic; use Testcontainers, Supertest, and Playwright only for applicable integration and e2e cases.
- Keep 100% coverage and mutation gates from the shared policy. Run Stryker and Uncle Bob Mutator after merge; the shared locked quality toolchain runs Dryer for this package.
- Publish line and branch coverage from Jest. The `test:mutation` script runs mutation tests; the Stryker job mutates production code separately.
- Keep `template/deploy/Dockerfiles/Dockerfile` with `production` and inherited `tests` targets; restore production dependencies first and add locked test dependencies only in `tests`.
