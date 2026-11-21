namespace DddTemplate.Tests.Unit.Interface;

using DddTemplate.Interface;
using Microsoft.AspNetCore.Mvc;
using Xunit;

public class ResponseTests
{
    [Fact]
    public void Factories_SetResponseStatusAndPayload()
    {
        var success = ResponseDTO.Ok("value");
        var failure = ResponseDTO.Fail("reason");

        Assert.True(success.Success);
        Assert.Equal("value", success.Data);
        Assert.False(failure.Success);
        Assert.Equal("reason", failure.Error);
    }

    [Fact]
    public void RequestDTO_StoresRequestFields()
    {
        var request = new TestRequest { Id = "one", Params = new() { ["key"] = "value" } };

        Assert.Equal("one", request.Id);
        Assert.Equal("value", request.Params["key"]);
    }

    [Fact]
    public void ToActionResult_MapsSuccessAndFailureResponses()
    {
        IActionResult noContent = ResponseDTO.Ok().ToActionResult();
        IActionResult ok = ResponseDTO.Ok("value").ToActionResult();
        IActionResult badRequest = ResponseDTO.Fail("reason").ToActionResult();

        Assert.IsType<NoContentResult>(noContent);
        Assert.IsType<OkObjectResult>(ok);
        Assert.IsType<BadRequestObjectResult>(badRequest);
    }

    private sealed class TestRequest : RequestDTO { }
}
