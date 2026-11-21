FROM node:22.23.5-bookworm-slim AS dependencies

WORKDIR /workspace
RUN corepack enable
COPY src/templates/nodejs/package.json src/templates/nodejs/pnpm-lock.yaml ./src/templates/nodejs/
COPY src/templates/nodejs/tsconfig.json ./src/templates/nodejs/
COPY src/templates/nodejs/src/ ./src/templates/nodejs/src/
RUN cd src/templates/nodejs \
    && pnpm install --frozen-lockfile

FROM dependencies AS application-build
RUN cd src/templates/nodejs && pnpm build

FROM dependencies AS production-dependencies
RUN cd src/templates/nodejs && pnpm prune --prod

FROM node:22.23.5-bookworm-slim AS production
RUN apt-get update \
    && apt-get upgrade -y --no-install-recommends \
    && rm -rf /usr/local/lib/node_modules/npm /usr/local/lib/node_modules/corepack \
        /usr/local/bin/npm /usr/local/bin/npx /usr/local/bin/corepack \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /workspace
COPY --from=application-build /workspace/src/templates/nodejs/package.json ./src/templates/nodejs/
COPY --from=application-build /workspace/src/templates/nodejs/dist/ ./src/templates/nodejs/dist/
COPY --from=production-dependencies /workspace/src/templates/nodejs/node_modules/ ./src/templates/nodejs/node_modules/

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
COPY --from=dependencies /workspace/src/templates/nodejs/node_modules/ ./src/templates/nodejs/node_modules/
COPY --from=application-build /workspace/src/templates/nodejs/src/ ./src/templates/nodejs/src/
COPY src/templates/nodejs/tests/ ./src/templates/nodejs/tests/
COPY src/templates/nodejs/jest.config.js src/templates/nodejs/stryker.config.json ./src/templates/nodejs/
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
