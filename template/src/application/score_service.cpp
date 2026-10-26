#include "application/score_service.hpp"

#include "domain/score.hpp"

namespace example::application {

int calculate_score(const int raw_score) noexcept {
  return domain::normalize_score(raw_score);
}

} // namespace example::application
