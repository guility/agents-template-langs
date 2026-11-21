"""Application layer - use cases и порты (интерфейсы)."""

from abc import ABC, abstractmethod


class RepositoryPort[T](ABC):
    """Порт для репозиториев (контракт инфраструктуры)."""

    @abstractmethod
    def get_by_id(self, id: str) -> T | None:
        """Получение сущности по ID."""
        raise NotImplementedError

    @abstractmethod
    def get_all(self) -> list[T]:
        """Получение всех сущностей."""
        raise NotImplementedError

    @abstractmethod
    def save(self, entity: T) -> T:
        """Сохранение сущности."""
        raise NotImplementedError

    @abstractmethod
    def delete(self, id: str) -> bool:
        """Удаление сущности."""
        raise NotImplementedError


class UseCase[T](ABC):
    """Базовый класс для use case."""

    @abstractmethod
    def execute(self, **kwargs) -> T:
        """Выполнение use case."""
        raise NotImplementedError
