# C# component rules

Follow the shared policy in [agents-template/AGENTS.md](https://github.com/guility/agents-template/blob/main/AGENTS.md). This file contains only C#-specific rules.

## Implementation

- Use .NET 10 LTS and the exact SDK version in `template/global.json`; commit `packages.lock.json`, and restore in locked mode.
- Keep production code in `template/src/DDDTemplate.csproj`; the root `DDDTemplate.Tests.csproj` references it so Stryker mutates the production assembly under test.
- Follow C# naming conventions. Use async/await correctly, implement `IDisposable` for owned resources, and document public APIs with XML comments.
- Use xUnit or NUnit. Name test classes `{ClassName}Tests` and methods `{Method}_{Scenario}_{Expected}`.
- Use `dotnet format` and SDK analyzers for C# formatting and static analysis.

## Tests and mutation

- Use Moq or NSubstitute for test doubles.
- Use pinned Stryker.NET 5.0.0 for C# mutation checks. Uncle Bob Mutator and Dryer are not used for C#.
- Publish line and branch coverage with Coverlet.
- Build the production project first and restore test packages only in the package's `tests` image target.
