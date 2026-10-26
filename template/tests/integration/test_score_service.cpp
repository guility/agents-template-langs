#include <exception>
#include <iostream>

#include "../support/check.hpp"
#include "application/score_service.hpp"

void test_application_service_returns_normalized_scores() {
  require_equal(example::application::calculate_score(-5), 0, "negative input");
  require_equal(example::application::calculate_score(50), 50, "valid input");
  require_equal(example::application::calculate_score(105), 100, "high input");
}

int main() {
  try {
    test_application_service_returns_normalized_scores();
  } catch (const std::exception &error) {
    std::cerr << error.what() << '\n';
    return 1;
  }
  return 0;
}
