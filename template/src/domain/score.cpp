#include "domain/score.hpp"

namespace example::domain {

int normalize_score(const int raw_score) noexcept {
  if (raw_score < 0) {
    return 0;
  }
  if (raw_score > 100) {
    return 100;
  }
  return raw_score;
}

} // namespace example::domain
