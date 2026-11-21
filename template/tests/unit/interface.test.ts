import { UseCase } from "../../src/application";
import { Controller, RequestDTO, ResponseDTO } from "../../src/interface";

class TestRequest extends RequestDTO {
  constructor(private readonly valid: boolean) {
    super();
    Object.freeze(this);
  }

  validate(): boolean {
    return this.valid;
  }
}

class TestController extends Controller {
  handle(request: TestRequest, useCase: UseCase<TestRequest, string>) {
    return this.handleRequest(request, useCase);
  }
}

describe("ResponseDTO", () => {
  it("freezes a successful response", () => {
    const response = new ResponseDTO(true, "value");
    expect(response.success).toBe(true);
    expect(response.data).toBe("value");
    expect(Object.isFrozen(response)).toBe(true);
  });
});

describe("Controller", () => {
  it("returns a successful result", async () => {
    const useCase: UseCase<TestRequest, string> = {
      execute: async () => "result",
    };
    await expect(
      new TestController().handle(new TestRequest(true), useCase),
    ).resolves.toMatchObject({ success: true, data: "result" });
  });

  it("rejects invalid requests without executing the use case", async () => {
    const execute = jest.fn();
    const response = await new TestController().handle(new TestRequest(false), {
      execute,
    });
    expect(response.success).toBe(false);
    expect(response.error).toBe("Invalid request");
    expect(execute).not.toHaveBeenCalled();
  });

  it("converts Error and non-Error rejections", async () => {
    const errorResponse = await new TestController().handle(
      new TestRequest(true),
      {
        execute: async () => {
          throw new Error("failed");
        },
      },
    );
    const unknownResponse = await new TestController().handle(
      new TestRequest(true),
      {
        execute: async () => {
          throw "failed";
        },
      },
    );
    expect(errorResponse.error).toBe("failed");
    expect(errorResponse.success).toBe(false);
    expect(unknownResponse.error).toBe("Unknown error");
    expect(unknownResponse.success).toBe(false);
  });
});
