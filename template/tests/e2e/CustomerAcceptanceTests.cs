namespace DddTemplate.Tests.E2E;

using DddTemplate.Infrastructure;
using DddTemplate.Interface;
using Microsoft.AspNetCore.Mvc;
using Xunit;

[Trait("Category", "E2E")]
public class CustomerAcceptanceTests
{
    [Fact]
    public async Task StoredCustomer_IsReturnedAsSuccessfulHttpResponse()
    {
        var repository = new InMemoryRepository<AcceptanceCustomer>();
        var customer = new AcceptanceCustomer("customer-1");
        await repository.AddAsync(customer);
        var stored = await repository.GetByIdAsync(customer.Id);

        var action = ResponseDTO.Ok(stored).ToActionResult();

        var response = Assert.IsType<OkObjectResult>(action);
        var body = Assert.IsType<ResponseDTO>(response.Value);
        Assert.True(body.Success);
        Assert.Equal(customer.Id, ((AcceptanceCustomer)body.Data!).Id);
    }

    private sealed record AcceptanceCustomer(string Id);
}
