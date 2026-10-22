# C# component rules

Apply the general requirements from `guility/agents_template/AGENTS.md`. This file contains only rules specific to C# components.

## Implementation

- Commit `packages.lock.json` and enable locked restore for package dependencies.
- Follow C# naming conventions. Use async/await correctly, implement `IDisposable` for owned resources, use LINQ where it improves clarity, and document public APIs with XML documentation.
- Use pattern matching where it makes control flow clearer. Format with `dotnet format` and analyze with Roslyn analyzers; check complexity with the selected analyzer.

## Tests

- Use xUnit or NUnit. Name tests `{Method}_{Scenario}_{Expected}` and use `[Fact]` or `[Theory]` as appropriate.
- Use `Moq` or `NSubstitute` for test doubles. Use `testcontainers-dotnet` for integration tests and SpecFlow or Playwright for applicable e2e tests.
- Use Stryker.NET for mutation testing and Coverlet or dotCover for coverage. Keep tests independent and reproducible, and meet the inherited 100% coverage and mutation gates.
