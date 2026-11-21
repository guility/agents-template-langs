FROM python:3.14.8-slim AS production

WORKDIR /workspace
COPY src/templates/python/pyproject.toml src/templates/python/uv.lock ./src/templates/python/
COPY src/templates/python/README.md ./src/templates/python/
COPY src/templates/python/src/ ./src/templates/python/src/
RUN python -m pip install --no-cache-dir uv==0.13.9 \
    && uv sync --project src/templates/python --locked --no-dev

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
COPY src/templates/python/tests/ ./src/templates/python/tests/
RUN uv sync --project src/templates/python --locked --extra dev \
    && mkdir -p src/templates/python/test-results src/templates/python/docs/coverage
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
