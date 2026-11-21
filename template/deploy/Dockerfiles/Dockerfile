FROM rust:1.100.0-slim AS production

WORKDIR /workspace
RUN apt-get update \
    && apt-get upgrade -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
COPY src/templates/rust/Cargo.toml src/templates/rust/Cargo.lock ./src/templates/rust/
COPY src/templates/rust/src/ ./src/templates/rust/src/
RUN cargo build --manifest-path src/templates/rust/Cargo.toml --locked --release

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
COPY src/templates/rust/tests/ ./src/templates/rust/tests/
RUN cargo test --manifest-path src/templates/rust/Cargo.toml --locked --no-run \
    && cargo install cargo-llvm-cov --version 0.9.1 --locked \
    && mkdir -p src/templates/rust/test-results src/templates/rust/docs/coverage
RUN rustup component add llvm-tools-preview --toolchain 1.100.0
ENV RUSTUP_TOOLCHAIN=1.100.0
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
