import { Entity } from "../../src/domain";
import { ConsoleLogger, InMemoryRepository } from "../../src/infrastructure";

class TestEntity extends Entity {
  validate(): boolean {
    return true;
  }
}

describe("InMemoryRepository", () => {
  it("saves, reads, lists and deletes entities", async () => {
    const repository = new InMemoryRepository<TestEntity>();
    const entity = new TestEntity("one");
    await repository.save(entity);
    expect(await repository.findById("one")).toBe(entity);
    expect(await repository.findAll()).toEqual([entity]);
    await repository.delete("one");
    expect(await repository.findById("one")).toBeNull();
    expect(await repository.findAll()).toEqual([]);
  });
});

describe("ConsoleLogger", () => {
  it("writes all log levels to their console methods", () => {
    const info = jest.spyOn(console, "log").mockImplementation();
    const error = jest.spyOn(console, "error").mockImplementation();
    const warn = jest.spyOn(console, "warn").mockImplementation();
    const logger = new ConsoleLogger();
    logger.info("info");
    logger.error("error");
    logger.warn("warn");
    expect(info).toHaveBeenCalledWith("[INFO] info");
    expect(error).toHaveBeenCalledWith("[ERROR] error");
    expect(warn).toHaveBeenCalledWith("[WARN] warn");
    jest.restoreAllMocks();
  });
});
