# agents-template-langs

Language-specific component rules and starter templates are published on a dedicated branch for each language. Each language branch contains only that language's rules and template. Version snapshots are immutable Git tags created from the matching language branch. The default branch contains only this index and repository-wide instructions.

All policies in [`guility/agents_template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md) apply here as well.

## Current packages

- Python: branch [`python`](https://github.com/guility/agents-template-langs/tree/python), current version [`python/v2.0.0`](https://github.com/guility/agents-template-langs/tree/python/v2.0.0) (Python 3.14.x, Ubuntu 26.04 LTS)
- Node.js/TypeScript: branch [`nodejs`](https://github.com/guility/agents-template-langs/tree/nodejs), current version [`nodejs/v1.0.1`](https://github.com/guility/agents-template-langs/tree/nodejs/v1.0.1)
- Rust: branch [`rust`](https://github.com/guility/agents-template-langs/tree/rust), current version [`rust/v1.0.1`](https://github.com/guility/agents-template-langs/tree/rust/v1.0.1)
- Go: branch [`go`](https://github.com/guility/agents-template-langs/tree/go), current version [`go/v1.0.1`](https://github.com/guility/agents-template-langs/tree/go/v1.0.1)
- C#: branch [`csharp`](https://github.com/guility/agents-template-langs/tree/csharp), current version [`csharp/v1.0.1`](https://github.com/guility/agents-template-langs/tree/csharp/v1.0.1)

Use `git switch --track origin/<language>` to work on a language branch, or check out its version tag for an immutable snapshot. New tags use `<language>/v<major>.<minor>.<patch>` and must point to commits on the matching branch. Previously published `<language>-v<version>` tags remain unchanged for compatibility; create no further tags in that format.
