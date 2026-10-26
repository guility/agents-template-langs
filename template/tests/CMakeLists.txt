add_executable(test_score_unit unit/test_score.cpp)
target_link_libraries(test_score_unit PRIVATE cpp_component)
add_test(NAME unit COMMAND test_score_unit)
set_tests_properties(unit PROPERTIES LABELS "unit")

add_executable(test_score_integration integration/test_score_service.cpp)
target_link_libraries(test_score_integration PRIVATE cpp_component)
add_test(NAME integration COMMAND test_score_integration)
set_tests_properties(integration PROPERTIES LABELS "integration")

add_test(
  NAME e2e
  COMMAND "${CMAKE_COMMAND}" "-DPROGRAM=$<TARGET_FILE:score-cli>"
          -P "${CMAKE_CURRENT_SOURCE_DIR}/e2e/check_cli.cmake")
set_tests_properties(e2e PROPERTIES LABELS "e2e")

if(CPP_ENABLE_MUTATION)
  add_test(
    NAME mutation
    COMMAND bash "${CMAKE_CURRENT_SOURCE_DIR}/mutation/run_mutations.sh"
            "${PROJECT_SOURCE_DIR}")
  set_tests_properties(mutation PROPERTIES LABELS "mutation")
endif()
