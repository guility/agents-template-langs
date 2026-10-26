#include <exception>
#include <iostream>

#include "../support/check.hpp"
#include "domain/score.hpp"

void test_score_normalization_boundaries() {
  require_equal(example::domain::normalize_score(-1), 0, "below lower bound");
  require_equal(example::domain::normalize_score(0), 0, "lower bound");
  require_equal(example::domain::normalize_score(42), 42, "inside range");
  require_equal(example::domain::normalize_score(100), 100, "upper bound");
  require_equal(example::domain::normalize_score(101), 100,
                "above upper bound");
}

int main() {
  try {
    test_score_normalization_boundaries();
  } catch (const std::exception &error) {
    std::cerr << error.what() << '\n';
    return 1;
  }
  return 0;
}
