#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "$BASH_SOURCE")/../.." && pwd)"
temporary="$(mktemp -d)"
trap 'rm -rf "$temporary"' EXIT

for mutant in "$root"/tests/mutation/*.patch; do
  name="$(basename "$mutant" .patch)"
  sandbox="$temporary/$name"
  mkdir -p "$sandbox"
  tar --exclude=target -cf - -C "$root" . | tar -xf - -C "$sandbox"
  patch --batch -p1 -d "$sandbox" < "$mutant"
  log="$temporary/$name.log"
  if (cd "$sandbox" && CARGO_TARGET_DIR="$root/target/mutations" cargo build --locked) >"$log" 2>&1; then
    echo "mutant survived: $name" >&2
    exit 1
  fi
  if ! grep -Eq 'LNG0201|unsolved goals' "$log"; then
    cat "$log" >&2
    echo "mutant failed for an unexpected reason: $name" >&2
    exit 1
  fi
done

echo "all Lean semantic mutants rejected by proof checking"
