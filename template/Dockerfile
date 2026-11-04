FROM rust:1.89.0-bookworm

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates curl git patch \
    && rm -rf /var/lib/apt/lists/*

ENV ELAN_HOME=/opt/elan
ENV PATH=/opt/elan/bin:/usr/local/cargo/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

RUN curl --proto '=https' --tlsv1.2 -sSf https://elan.lean-lang.org/elan-init.sh \
    | sh -s -- -y --default-toolchain none \
    && elan toolchain install leanprover/lean4:v4.35.0

WORKDIR /workspace
COPY . .

CMD ["cargo", "test", "--locked"]
