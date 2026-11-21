FROM node:22.17.1-bookworm-slim AS production

WORKDIR /workspace
RUN corepack enable
COPY src/templates/nodejs/package.json src/templates/nodejs/pnpm-lock.yaml ./src/templates/nodejs/
COPY src/templates/nodejs/tsconfig.json ./src/templates/nodejs/
COPY src/templates/nodejs/src/ ./src/templates/nodejs/src/
RUN cd src/templates/nodejs \
    && pnpm install --frozen-lockfile \
    && pnpm build \
    && pnpm prune --prod

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
RUN cd src/templates/nodejs && pnpm install --frozen-lockfile
COPY src/templates/nodejs/tests/ ./src/templates/nodejs/tests/
COPY src/templates/nodejs/jest.config.js src/templates/nodejs/stryker.config.json ./src/templates/nodejs/
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
