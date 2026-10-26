# CUDA component starter

This starter demonstrates an explicit CUDA architecture setting, checked runtime calls, scoped device-memory ownership, and host/device shared math. Replace the sample with one bounded component and preserve the approved contracts and release gates.

## Build

Select a Toolkit-supported host compiler and the target SM list explicitly:

```bash
cmake -S . -B build -DCMAKE_BUILD_TYPE=Debug \
  -DCMAKE_CUDA_ARCHITECTURES=<target-sm-list>
cmake --build build --parallel
```

## Required suites

```bash
ctest --test-dir build --output-on-failure -L unit
ctest --test-dir build --output-on-failure -L integration
ctest --test-dir build --output-on-failure -L e2e
ctest --test-dir build --output-on-failure -L mutation
compute-sanitizer --tool memcheck --error-exitcode 1 build/tests/test_cuda_e2e
```

The integration, e2e, kernel-mutation, and sanitizer suites require the target GPU. They fail when it is unavailable. Host line/branch reports do not cover CUDA device functions: NVIDIA documents Bullseye instrumentation for host functions only. The inherited 100% coverage gate remains open until an approved device-code coverage method is available; never treat GPU smoke, mutation, or sanitizer checks as a replacement for that gate.

The official toolchain references are [CMake CUDA architecture selection](https://cmake.org/cmake/help/latest/variable/CMAKE_CUDA_ARCHITECTURES.html), [NVIDIA NVCC host-compiler support](https://docs.nvidia.com/cuda/cuda-compiler-driver-nvcc/), and [NVIDIA Compute Sanitizer](https://docs.nvidia.com/compute-sanitizer/).
