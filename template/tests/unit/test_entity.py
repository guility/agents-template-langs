"""Entity domain unit tests."""

from dataclasses import dataclass
from datetime import datetime

import pytest

from domain import BaseEntity, Entity


class TestEntity:
    """Тесты для Entity."""

    def test_entity_creation(self):
        """Создание сущности."""
        entity = Entity(id="entity-1")
        assert entity.id == "entity-1"
        assert entity.created_at is not None

    def test_entity_update(self):
        """Обновление сущности."""
        entity = Entity(id="entity-1")
        old_updated_at = entity.updated_at

        # Обновление должно изменить updated_at
        entity.update()

        assert entity.updated_at is not None
        if old_updated_at:
            assert entity.updated_at >= old_updated_at

    def test_entity_inherits_from_base(self):
        """Entity наследуется от BaseEntity."""
        entity = Entity(id="entity-1")
        assert isinstance(entity, BaseEntity)

    def test_entity_updates_mutable_subclass_field(self):
        """Производная сущность обновляет изменяемое доменное поле."""

        @dataclass(frozen=True)
        class NamedEntity(Entity):
            name: str = "before"

        entity = NamedEntity(id="entity-1")

        entity.update(name="after")

        assert entity.name == "after"

    def test_entity_rejects_identity_update(self):
        """Идентичность сущности нельзя изменить через update()."""
        entity = Entity(id="entity-1")

        with pytest.raises(AttributeError, match="id is immutable"):
            entity.update(id="entity-2")

    def test_entity_rejects_created_at_update(self):
        entity = Entity(id="entity-1")

        with pytest.raises(AttributeError, match="created_at is immutable"):
            entity.update(created_at=datetime(2026, 1, 1))

    def test_entity_ignores_unknown_fields(self):
        entity = Entity(id="entity-1")

        entity.update(unknown="value")

        assert not hasattr(entity, "unknown")
