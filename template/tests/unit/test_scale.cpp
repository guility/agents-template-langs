#include <cuda_runtime_api.h>

#include <cmath>
#include <exception>
#include <iostream>
#include <limits>

#include "../support/check.hpp"
#include "cuda/scale.hpp"
#include "cuda/status.hpp"

void test_host_contracts_and_cuda_status() {
  require(example::cuda::scale_value(2.0F, 3.0F) == 6.0F, "host/device scale");
  example::cuda::check_cuda_status(cudaSuccess, "expected success");

  bool status_rejected = false;
  try {
    example::cuda::check_cuda_status(cudaErrorInvalidValue, "expected error");
  } catch (const std::runtime_error &) {
    status_rejected = true;
  }
  require(status_rejected, "CUDA errors must throw");

  bool factor_rejected = false;
  try {
    (void)example::cuda::scale_on_gpu({},
                                      std::numeric_limits<float>::infinity());
  } catch (const std::invalid_argument &) {
    factor_rejected = true;
  }
  require(factor_rejected, "non-finite factor must fail before device access");
  require(example::cuda::scale_on_gpu({}, 2.0F).empty(),
          "empty input must return without device access");
}

int main() {
  try {
    test_host_contracts_and_cuda_status();
  } catch (const std::exception &error) {
    std::cerr << error.what() << '\n';
    return 1;
  }
  return 0;
}
