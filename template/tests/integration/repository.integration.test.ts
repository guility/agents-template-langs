import { Repository } from "../../src/application";
import { Entity } from "../../src/domain";
import { InMemoryRepository } from "../../src/infrastructure";

class Customer extends Entity {
  validate(): boolean {
    return this.id.length > 0;
  }
}

describe("repository port integration", () => {
  it("stores and returns an entity through the application contract", async () => {
    const repository: Repository<Customer> = new InMemoryRepository<Customer>();
    const customer = new Customer("customer-1");

    await repository.save(customer);

    expect(await repository.findById(customer.id)).toBe(customer);
    expect(await repository.findAll()).toEqual([customer]);
  });
});
