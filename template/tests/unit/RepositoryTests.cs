namespace DddTemplate.Tests.Unit.Infrastructure;

using DddTemplate.Infrastructure;
using Xunit;

public class InMemoryRepositoryTests
{
    [Fact]
    public async Task AddAndRead_ReturnStoredEntity()
    {
        var repository = new InMemoryRepository<Record>();
        var record = new Record("one");

        await repository.AddAsync(record);

        Assert.Same(record, await repository.GetByIdAsync("one"));
        Assert.Single(await repository.GetAllAsync());
    }

    [Fact]
    public async Task Add_RejectsDuplicateIdentity()
    {
        var repository = new InMemoryRepository<Record>();
        await repository.AddAsync(new Record("one"));

        var duplicateError = await Assert.ThrowsAsync<InvalidOperationException>(
            () => repository.AddAsync(new Record("one")));
        Assert.Equal("Entity with id one already exists", duplicateError.Message);
    }

    [Fact]
    public async Task Add_GeneratesIdentityWhenMissing()
    {
        var repository = new InMemoryRepository<UnidentifiedRecord>();
        var added = await repository.AddAsync(new UnidentifiedRecord());

        Assert.NotNull(added);
        Assert.Single(await repository.GetAllAsync());
    }

    [Fact]
    public async Task Add_GeneratesIdentityWhenIdIsNull()
    {
        var repository = new InMemoryRepository<NullIdentityRecord>();

        await repository.AddAsync(new NullIdentityRecord());

        Assert.Single(await repository.GetAllAsync());
    }

    [Fact]
    public async Task Update_RejectsMissingAndUnknownIdentity()
    {
        var repository = new InMemoryRepository<Record>();
        var nullIdentityRepository = new InMemoryRepository<NullIdentityRecord>();
        var unidentifiedRepository = new InMemoryRepository<UnidentifiedRecord>();

        var emptyIdError = await Assert.ThrowsAsync<InvalidOperationException>(
            () => repository.UpdateAsync(new Record("")));
        var nullIdError = await Assert.ThrowsAsync<InvalidOperationException>(
            () => nullIdentityRepository.UpdateAsync(new NullIdentityRecord()));
        var missingIdError = await Assert.ThrowsAsync<InvalidOperationException>(
            () => unidentifiedRepository.UpdateAsync(new UnidentifiedRecord()));
        Assert.Equal("Entity must have an Id", emptyIdError.Message);
        Assert.Equal(emptyIdError.Message, nullIdError.Message);
        Assert.Equal(emptyIdError.Message, missingIdError.Message);
        var unknownIdError = await Assert.ThrowsAsync<KeyNotFoundException>(
            () => repository.UpdateAsync(new Record("missing")));
        Assert.Equal("Entity with id missing not found", unknownIdError.Message);
    }

    [Fact]
    public async Task UpdateAndDelete_ChangeStoredEntities()
    {
        var repository = new InMemoryRepository<Record>();
        var initial = new Record("one");
        await repository.AddAsync(initial);
        var updated = new Record("one");

        await repository.UpdateAsync(updated);
        Assert.Same(updated, await repository.GetByIdAsync("one"));
        Assert.Single(await repository.FindWhereAsync(item => item.Id == "one"));
        Assert.Empty(await repository.FindWhereAsync(item => item.Id == "two"));

        await repository.DeleteAsync("one");
        Assert.Null(await repository.GetByIdAsync("one"));
        var missingError = await Assert.ThrowsAsync<KeyNotFoundException>(
            () => repository.DeleteAsync("one"));
        Assert.Equal("Entity with id one not found", missingError.Message);
    }

    public sealed record Record(string Id);
    public sealed class UnidentifiedRecord { }
    public sealed record NullIdentityRecord(string? Id = null);
}
