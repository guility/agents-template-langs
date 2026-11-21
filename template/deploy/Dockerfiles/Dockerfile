FROM python:3.14.8-slim AS build

WORKDIR /workspace
COPY src/templates/python/pyproject.toml src/templates/python/uv.lock ./src/templates/python/
COPY src/templates/python/README.md ./src/templates/python/
COPY src/templates/python/src/ ./src/templates/python/src/
RUN python -m pip install --no-cache-dir uv==0.13.9 \
    && uv sync --project src/templates/python --locked --no-dev

FROM python:3.14.8-slim AS production
RUN apt-get update \
    && apt-get upgrade -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /workspace
COPY --from=build /workspace/src/templates/python/ ./src/templates/python/
RUN python -c "import pip, shutil; from pathlib import Path; package = Path(pip.__file__).parent; shutil.rmtree(package); [shutil.rmtree(path) for path in package.parent.glob('pip-*.dist-info')]" \
    && rm -f /usr/local/bin/pip /usr/local/bin/pip3 /usr/local/bin/pip3.14
ENV PATH="/workspace/src/templates/python/.venv/bin:${PATH}"

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build /usr/local/bin/uv /usr/local/bin/uv
COPY src/templates/python/tests/ ./src/templates/python/tests/
RUN uv sync --project src/templates/python --locked --extra dev \
    && mkdir -p src/templates/python/test-results src/templates/python/docs/coverage
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
