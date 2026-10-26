#include <exception>
#include <iostream>
#include <vector>

#include "../support/check.hpp"
#include "cuda/scale.hpp"

void test_scale_values_on_target_gpu() {
  const std::vector<float> input{2.0F, -1.5F, 0.25F};
  const auto output = example::cuda::scale_on_gpu(input, -2.0F);
  require(output.size() == 3, "integration result size");
  require(output[0] == -4.0F && output[1] == 3.0F && output[2] == -0.5F,
          "integration device result");
}

int main() {
  try {
    test_scale_values_on_target_gpu();
  } catch (const std::exception &error) {
    std::cerr << error.what() << '\n';
    return 1;
  }
  return 0;
}
