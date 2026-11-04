#!/usr/bin/env bash
set -euo pipefail

if [[ ! -d /workspace/lean/.lake/packages/lungo ]]; then
  (cd /workspace/lean && lake update)
fi

exec "$@"
