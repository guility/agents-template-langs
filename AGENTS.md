# C# component rules

Follow the shared policy in [`agents-template/AGENTS.md`](https://github.com/guility/agents_template/blob/main/AGENTS.md). This package contains only C#-specific rules.

## Implementation

- Use .NET 10 LTS and the exact SDK version in `template/global.json`; commit `packages.lock.json`, and restore in locked mode.
- Keep production code in `template/src/DDDTemplate.csproj`; the root `DDDTemplate.Tests.csproj` references it so Stryker mutates the production assembly under test.
- Follow C# naming conventions. Use async/await correctly, implement `IDisposable` for owned resources, and document public APIs with XML comments.
- Use xUnit or NUnit. Name test classes `{ClassName}Tests` and methods `{Method}_{Scenario}_{Expected}`.
- Run `dotnet format` and SDK analyzers in the MR gate. These checks exceed the shared 0.5-second fast-push limit, so this language has no formatter or analyzer in the fast class. Use Lizard for CCN (maximum 10) and shared metrics.

## Tests and mutation

- Use Moq or NSubstitute for test doubles. Use Testcontainers and Playwright only for applicable integration and e2e cases.
- Keep 100% coverage and mutation gates from the shared policy. Run pinned Stryker.NET 5.0.0 after merge. Uncle Bob Mutator and Dryer are not used for C#.
- Publish line and branch coverage with Coverlet.
- Keep `template/deploy/Dockerfiles/Dockerfile` with `production` and inherited `tests` targets; build the production project first and restore test packages only in `tests`.
