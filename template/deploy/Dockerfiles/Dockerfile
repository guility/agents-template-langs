FROM golang:1.27.4-bookworm AS production

WORKDIR /workspace
RUN apt-get update \
    && apt-get upgrade -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
COPY src/templates/go/go.mod src/templates/go/go.sum ./src/templates/go/
COPY src/templates/go/src/ ./src/templates/go/src/
RUN cd src/templates/go && go build ./...

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
COPY src/templates/go/tests/ ./src/templates/go/tests/
RUN cd src/templates/go && go test -run '^$' ./...
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
