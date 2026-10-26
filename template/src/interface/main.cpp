#include <charconv>
#include <iostream>
#include <string_view>
#include <system_error>

#include "application/score_service.hpp"

int main(const int argc, char *argv[]) {
  if (argc != 3 || std::string_view(argv[1]) != "--score") {
    std::cerr << "usage: score-cli --score INTEGER\n";
    return 2;
  }

  int raw_score = 0;
  const std::string_view input{argv[2]};
  const auto [end, error] =
      std::from_chars(input.data(), input.data() + input.size(), raw_score);
  if (error != std::errc{} || end != input.data() + input.size()) {
    std::cerr << "score must be an integer\n";
    return 2;
  }

  std::cout << example::application::calculate_score(raw_score) << '\n';
  return 0;
}
