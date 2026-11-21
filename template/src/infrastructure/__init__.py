"""Infrastructure layer - реализация портов."""

from typing import Any

from application import RepositoryPort


class InMemoryRepository(RepositoryPort[Any]):
    """In-memory реализация репозитория (для тестов)."""

    def __init__(self):
        self._storage: dict[str, Any] = {}

    def get_by_id(self, id: str) -> Any | None:
        return self._storage.get(id)

    def get_all(self) -> list[Any]:
        return list(self._storage.values())

    def save(self, entity: Any) -> Any:
        self._storage[entity.id] = entity
        return entity

    def delete(self, id: str) -> bool:
        if id in self._storage:
            del self._storage[id]
            return True
        return False
