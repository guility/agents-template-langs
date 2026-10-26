#include <exception>
#include <iostream>
#include <vector>

#include "../support/check.hpp"
#include "cuda/scale.hpp"

void test_multi_block_scale_matches_reference() {
  std::vector<float> input(257);
  for (std::size_t index = 0; index < input.size(); ++index) {
    input[index] = static_cast<float>(index) - 128.0F;
  }
  const auto output = example::cuda::scale_on_gpu(input, 1.5F);
  require(output.size() == input.size(), "e2e result size");
  for (std::size_t index = 0; index < input.size(); ++index) {
    require(output[index] == input[index] * 1.5F, "e2e device result");
  }
}

int main() {
  try {
    test_multi_block_scale_matches_reference();
  } catch (const std::exception &error) {
    std::cerr << error.what() << '\n';
    return 1;
  }
  return 0;
}
