function(assert_cli expected_status expected_output)
  execute_process(
    COMMAND "${PROGRAM}" ${ARGN}
    RESULT_VARIABLE actual_status
    OUTPUT_VARIABLE actual_output
    ERROR_QUIET)
  if(NOT actual_status EQUAL expected_status)
    message(FATAL_ERROR "CLI exit ${actual_status}; expected ${expected_status}")
  endif()
  if(NOT "${actual_output}" STREQUAL "${expected_output}")
    message(FATAL_ERROR "CLI output '${actual_output}'; expected '${expected_output}'")
  endif()
endfunction()

assert_cli(0 "100\n" --score 101)
assert_cli(2 "" --score invalid)
assert_cli(2 "" --score 0 extra)
assert_cli(2 "" --wrong 1)
assert_cli(2 "" --score 12tail)
