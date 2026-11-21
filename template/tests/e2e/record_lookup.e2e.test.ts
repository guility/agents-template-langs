import { Repository, UseCase } from "../../src/application";
import { Entity } from "../../src/domain";
import { InMemoryRepository } from "../../src/infrastructure";
import { Controller, RequestDTO, ResponseDTO } from "../../src/interface";

class Customer extends Entity {
  validate(): boolean {
    return this.id.length > 0;
  }
}

class LookupRequest extends RequestDTO {
  constructor(readonly customerId: string) {
    super();
  }

  validate(): boolean {
    return this.customerId.length > 0;
  }
}

class LookupCustomer implements UseCase<LookupRequest, Customer | null> {
  constructor(private readonly repository: Repository<Customer>) {}

  execute(request: LookupRequest): Promise<Customer | null> {
    return this.repository.findById(request.customerId);
  }
}

class CustomerController extends Controller {
  process(
    request: LookupRequest,
    useCase: UseCase<LookupRequest, Customer | null>,
  ): Promise<ResponseDTO<Customer | null>> {
    return this.handleRequest(request, useCase);
  }
}

describe("customer lookup acceptance", () => {
  it("returns a stored customer through the controller response", async () => {
    const repository = new InMemoryRepository<Customer>();
    const customer = new Customer("customer-1");
    await repository.save(customer);
    const response: ResponseDTO<Customer | null> =
      await new CustomerController().process(
        new LookupRequest(customer.id),
        new LookupCustomer(repository),
      );

    expect(response.success).toBe(true);
    expect(response.data).toBe(customer);
  });

  it("rejects a request without a customer identity", async () => {
    const response = await new CustomerController().process(
      new LookupRequest(""),
      new LookupCustomer(new InMemoryRepository<Customer>()),
    );

    expect(response.success).toBe(false);
    expect(response.error).toBe("Invalid request");
  });
});
