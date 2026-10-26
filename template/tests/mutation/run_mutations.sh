#!/usr/bin/env bash
set -euo pipefail

source_root="${1:?source root is required}"
temporary_root="$(mktemp -d)"
trap 'rm -rf "$temporary_root"' EXIT
killed=0
total=0

for mutant in "$source_root"/tests/mutation/{domain,application,interface}/*.patch; do
  [[ -f "$mutant" ]] || continue
  category="$(basename "$(dirname "$mutant")")"
  project="$temporary_root/project-$category-$(basename "$mutant" .patch)"
  mkdir -p "$project"
  cp -a "$source_root"/CMakeLists.txt "$source_root"/include "$source_root"/src "$source_root"/tests "$project"/
  patch -d "$project" -p1 < "$mutant"
  cmake -S "$project" -B "$project/build" -DCPP_ENABLE_MUTATION=OFF
  cmake --build "$project/build" --parallel
  case "$category" in
    domain) label=unit ;;
    application) label=integration ;;
    interface) label=e2e ;;
    *) echo "unknown mutation category: $category" >&2; exit 2 ;;
  esac
  total=$((total + 1))
  if ctest --test-dir "$project/build" --output-on-failure -L "$label"; then
    echo "surviving mutant: $mutant" >&2
    exit 1
  fi
  killed=$((killed + 1))
done

if ((total != 6 || killed != total)); then
  echo "expected six killed mutants, got $killed/$total" >&2
  exit 1
fi
