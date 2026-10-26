# AGENTS.md — agents-template-langs

All repository policies in [`guility/agents_template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md) also apply here and to every language tag. A policy still applies when it is not repeated in this file.

## Repository-specific rules

- Maintain one long-lived branch per language, named `python`, `nodejs`, `rust`, `go`, `csharp`, `cpp`, or `cuda`. Each branch and its version tags contain only rules and templates for that language; keep shared requirements in `agents_template`.
- Publish immutable, annotated version tags named `<language>/v<major>.<minor>.<patch>` on the matching language branch. Do not move or overwrite published tags. Each package contains a language-specific root `AGENTS.md` and its component template under `template/`.
- The old `<language>-v<version>` tags are retained as immutable compatibility refs; do not create further tags in that format.
- Keep the default branch limited to repository-wide instructions and the tag index. Do not duplicate language package contents on the default branch.
- Follow inherited limits: at most 8 files and 5 directories directly within any directory; for source, tests, scripts, and manually maintained configuration, every file is at most 300 physical lines and at least 95% are at most 100 lines. Lock files, generated files, vendored code, dependencies, and build outputs are excluded from line counts.
- Required test suites must meet inherited 100% coverage gates. GitLab CI uses the inherited stage order and runs only after merge requests are merged into the protected default branch.
- Keep package contents language-specific. Do not add general architecture, security, workflow, or CI policy to a language tag.
