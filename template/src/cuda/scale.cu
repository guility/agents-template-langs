#include "cuda/scale.hpp"

#include <cuda_runtime.h>

#include <algorithm>
#include <cmath>
#include <cstddef>
#include <memory>
#include <stdexcept>
#include <vector>

#include "cuda/status.hpp"

namespace example::cuda {
namespace {

constexpr unsigned int kThreadsPerBlock = 128;

__global__ void scale_kernel(float *values, const float factor) {
  const std::size_t index =
      static_cast<std::size_t>(blockIdx.x) * blockDim.x + threadIdx.x;
  values[index] = scale_value(values[index], factor);
}

} // namespace

std::vector<float> scale_on_gpu(const std::vector<float> &values,
                                const float factor) {
  if (!std::isfinite(factor)) {
    throw std::invalid_argument("scale factor must be finite");
  }
  if (values.empty()) {
    return {};
  }

  const std::size_t threads = kThreadsPerBlock;
  const std::size_t padded_count =
      ((values.size() + threads - 1) / threads) * threads;
  std::vector<float> padded(padded_count, 0.0F);
  std::copy(values.begin(), values.end(), padded.begin());

  float *raw_device_values = nullptr;
  check_cuda_status(cudaMalloc(reinterpret_cast<void **>(&raw_device_values),
                               padded.size() * sizeof(float)),
                    "cudaMalloc");
  std::unique_ptr<float, decltype(&cudaFree)> device_values(raw_device_values,
                                                            &cudaFree);
  check_cuda_status(cudaMemcpy(device_values.get(), padded.data(),
                               padded.size() * sizeof(float),
                               cudaMemcpyHostToDevice),
                    "copy inputs to device");

  const auto block_count = static_cast<unsigned int>(padded_count / threads);
  scale_kernel<<<block_count, kThreadsPerBlock>>>(device_values.get(), factor);
  check_cuda_status(cudaGetLastError(), "launch scale kernel");
  check_cuda_status(cudaDeviceSynchronize(), "synchronize scale kernel");
  check_cuda_status(cudaMemcpy(padded.data(), device_values.get(),
                               padded.size() * sizeof(float),
                               cudaMemcpyDeviceToHost),
                    "copy results from device");
  check_cuda_status(cudaFree(device_values.get()), "free device memory");
  device_values.release();

  return {padded.begin(), padded.begin() + values.size()};
}

} // namespace example::cuda
