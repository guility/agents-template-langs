namespace DddTemplate.Tests.Integration.Infrastructure;

using DddTemplate.Infrastructure;
using Xunit;

[Trait("Category", "Integration")]
public class RepositoryIntegrationTests
{
    [Fact]
    public async Task AddAndUpdate_PersistChangesAcrossRepositoryCalls()
    {
        var repository = new InMemoryRepository<IntegrationRecord>();
        var original = new IntegrationRecord("record-1", "Draft");
        await repository.AddAsync(original);
        var updated = new IntegrationRecord("record-1", "Accepted");

        await repository.UpdateAsync(updated);

        Assert.Same(updated, await repository.GetByIdAsync("record-1"));
        Assert.Single(await repository.FindWhereAsync(item => item.Status == "Accepted"));
    }

    private sealed record IntegrationRecord(string Id, string Status);
}
