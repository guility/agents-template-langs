# agents-template-langs

Language-specific component rules and starter templates are published on a dedicated branch for each language. Each language branch contains only that language's rules and template. Version snapshots are immutable Git tags created from the matching language branch. The default branch contains only this index and repository-wide instructions.

All shared policies in [`guility/agents-template/AGENTS.md`](https://github.com/guility/agents-template/blob/main/AGENTS.md) apply here as well. Language branch names are exceptions only in this repository.

The [component × metric quality matrix](https://github.com/guility/agents-template/blob/3bd13749c2cb7ad44424365dcc99b46930b06bfe/docs/quality/quality-matrix.md) and [per-tool push timings](https://github.com/guility/agents-template/blob/3bd13749c2cb7ad44424365dcc99b46930b06bfe/docs/quality/fast-check-benchmark.md) are published in `agents-template`.

## Current packages

- Python: branch [`python`](https://github.com/guility/agents-template-langs/tree/python), current version [`python/v2.0.4`](https://github.com/guility/agents-template-langs/tree/python/v2.0.4) (Python 3.14.8)
- Node.js/TypeScript: branch [`nodejs`](https://github.com/guility/agents-template-langs/tree/nodejs), current version [`nodejs/v1.0.5`](https://github.com/guility/agents-template-langs/tree/nodejs/v1.0.5)
- Rust: branch [`rust`](https://github.com/guility/agents-template-langs/tree/rust), current version [`rust/v1.0.4`](https://github.com/guility/agents-template-langs/tree/rust/v1.0.4)
- Go: branch [`go`](https://github.com/guility/agents-template-langs/tree/go), current version [`go/v1.0.4`](https://github.com/guility/agents-template-langs/tree/go/v1.0.4)
- C#: branch [`csharp`](https://github.com/guility/agents-template-langs/tree/csharp), current version [`csharp/v1.0.6`](https://github.com/guility/agents-template-langs/tree/csharp/v1.0.6)

Use `git switch --track origin/<language>` to work on a language branch, or check out its version tag for an immutable snapshot. New tags use `<language>/v<major>.<minor>.<patch>` and must point to commits on the matching branch. Previously published `<language>-v<version>` tags remain unchanged for compatibility; create no further tags in that format.
