#pragma once

#include <vector>

#if defined(__CUDACC__)
#define EXAMPLE_HOST_DEVICE __host__ __device__
#else
#define EXAMPLE_HOST_DEVICE
#endif

namespace example::cuda {

[[nodiscard]] EXAMPLE_HOST_DEVICE inline float
scale_value(const float value, const float factor) noexcept {
  return value * factor;
}

[[nodiscard]] std::vector<float> scale_on_gpu(const std::vector<float> &values,
                                              float factor);

} // namespace example::cuda

#undef EXAMPLE_HOST_DEVICE
