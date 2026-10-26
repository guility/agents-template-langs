cmake_minimum_required(VERSION 3.28)
project(cpp_component VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

include(CTest)
option(CPP_ENABLE_COVERAGE "Instrument code for line and branch coverage" OFF)
option(CPP_ENABLE_MUTATION "Register the required mutation suite" ON)

if(CPP_ENABLE_COVERAGE)
  if(NOT CMAKE_CXX_COMPILER_ID MATCHES "GNU|Clang")
    message(FATAL_ERROR "Coverage is configured for GCC or Clang")
  endif()
  add_compile_options(--coverage -O0 -g)
  add_link_options(--coverage)
endif()

add_library(cpp_component
  src/application/score_service.cpp
  src/domain/score.cpp)
target_include_directories(cpp_component PUBLIC include)
target_compile_options(cpp_component PRIVATE -Wall -Wextra -Wpedantic -Werror)

add_executable(score-cli src/interface/main.cpp)
target_link_libraries(score-cli PRIVATE cpp_component)
target_compile_options(score-cli PRIVATE -Wall -Wextra -Wpedantic -Werror)

if(BUILD_TESTING)
  add_subdirectory(tests)
endif()
