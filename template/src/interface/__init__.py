"""Interface layer - контроллеры и DTO."""

from dataclasses import dataclass
from typing import Any


@dataclass
class ResponseDTO:
    """Базовый DTO для ответов."""

    success: bool
    data: Any = None
    error: str | None = None

    @classmethod
    def ok(cls, data: Any = None):
        return cls(success=True, data=data)

    @classmethod
    def fail(cls, error: str):
        return cls(success=False, error=error)

    def to_dict(self) -> dict[str, Any]:
        return {
            "success": self.success,
            "data": self.data,
            "error": self.error,
        }
