"""Shared fixtures for end-to-end workflow tests."""

import pytest

from infrastructure import InMemoryRepository


@pytest.fixture
def setup_system():
    return {"repository": InMemoryRepository()}
