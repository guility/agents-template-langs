#pragma once

#include <stdexcept>
#include <string>

inline void require(const bool condition, const char *message) {
  if (!condition) {
    throw std::runtime_error(std::string{message});
  }
}
