#pragma once

#include <stdexcept>
#include <string>

inline void require_equal(const int actual, const int expected,
                          const char *label) {
  if (actual != expected) {
    throw std::runtime_error(std::string{label} + ": unexpected value");
  }
}
