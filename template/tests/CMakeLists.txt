add_executable(test_cuda_unit unit/test_scale.cpp)
target_link_libraries(test_cuda_unit PRIVATE cuda_component)
add_test(NAME unit COMMAND test_cuda_unit)
set_tests_properties(unit PROPERTIES LABELS "unit")

add_executable(test_cuda_integration integration/test_scale.cpp)
target_link_libraries(test_cuda_integration PRIVATE cuda_component)
add_test(NAME integration COMMAND test_cuda_integration)
set_tests_properties(integration PROPERTIES LABELS "integration")

add_executable(test_cuda_e2e e2e/test_scale.cpp)
target_link_libraries(test_cuda_e2e PRIVATE cuda_component)
add_test(NAME e2e COMMAND test_cuda_e2e)
set_tests_properties(e2e PROPERTIES LABELS "e2e")

if(CUDA_ENABLE_MUTATION)
  add_test(
    NAME mutation
    COMMAND bash "${CMAKE_CURRENT_SOURCE_DIR}/mutation/run_mutations.sh"
            "${PROJECT_SOURCE_DIR}")
  set_tests_properties(mutation PROPERTIES LABELS "mutation")
endif()
