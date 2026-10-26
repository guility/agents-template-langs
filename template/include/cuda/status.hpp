#pragma once

#include <cuda_runtime_api.h>

#include <stdexcept>
#include <string>

namespace example::cuda {

inline void check_cuda_status(const cudaError_t status, const char *operation) {
  if (status != cudaSuccess) {
    throw std::runtime_error(std::string{operation} + ": " +
                             cudaGetErrorString(status));
  }
}

} // namespace example::cuda
