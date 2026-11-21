import { BaseEntity, Entity, ValueObject } from "../../src/domain";

class ConcreteEntity extends Entity {
  validate(): boolean {
    return this.id.length > 0;
  }
}

class ConcreteValueObject extends ValueObject {}

describe("BaseEntity", () => {
  it("sets identity, creation time and update time", () => {
    const entity = new ConcreteEntity("test-123");
    expect(entity.id).toBe("test-123");
    expect(entity.createdAt).toBeInstanceOf(Date);
    expect(entity.updatedAt).toBeUndefined();
    entity.updateTimestamp();
    expect(entity.updatedAt).toBeInstanceOf(Date);
  });
});

describe("ValueObject", () => {
  it("is immutable", () => {
    expect(Object.isFrozen(new ConcreteValueObject())).toBe(true);
  });
});

describe("Entity", () => {
  it("validates nonempty and empty identities", () => {
    const entity = new ConcreteEntity("entity-1");
    expect(entity).toBeInstanceOf(BaseEntity);
    expect(entity.validate()).toBe(true);
    expect(new ConcreteEntity("").validate()).toBe(false);
  });
});
