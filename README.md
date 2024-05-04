# u8p

## Overview

The `u8p` package provides functionality for working with UTF-8 encoded strings in Go. It includes a key function, `Find`, which is designed to find the index of the leading UTF-8 byte in a given string.

## Function `Find`

- **Signature**: `Find(a string, l int) (int, error)`
- **Parameters**:
  - `a`: The string to be analyzed.
  - `l`: An integer representing the length of a substring to be considered.
- **Return**:
  - `int`: The index of the leading UTF-8 byte in the string.
  - `error`: An error if the input does not meet certain conditions.
- **Description**: `Find` searches for the leading UTF-8 byte in the last `l` bytes of the string `a`. If `a` is empty, it returns 0 and no error. It returns an error if `l` is less than or equal to 3, or if the length of `a` is less than `l`. It also returns an error if a valid UTF-8 leading byte is not found.

## Example Usage

Refer to `example_test.go` for example usage of the `Find` function.

## Benchmark Overview

This benchmark showcases the efficiency and simplicity of our string manipulation functions. Our methods are optimized for speed and achieve significant performance improvements compared to standard package implementations, avoiding unnecessary operations like casting.

### Test Environment

- **Hardware**: MacBook Air (M1, 2020)
- **OS**: macOS
- **Go Version**: 1.22.2

### Results

Our functions demonstrate remarkable performance across various test sizes, with minimal execution time and zero memory allocation.

```
BenchmarkFindUTF8Sizes/Size100-8                      514306195    2.253 ns/op       0 B/op   0 allocs/op
BenchmarkFindUTF8Sizes/Size1000-8                     535082955    3.833 ns/op       0 B/op   0 allocs/op
BenchmarkFindUTF8Sizes/Size10000-8                    338328892    5.406 ns/op       0 B/op   0 allocs/op
BenchmarkFindUTF8Sizes/Size100000-8                   564700567    2.153 ns/op       0 B/op   0 allocs/op
BenchmarkGetByteLengthOfRuneSlice/Size100-8             2032600    530.5 ns/op     416 B/op   1 allocs/op
BenchmarkGetByteLengthOfRuneSlice/Size500-8              444048     2728 ns/op    2176 B/op   2 allocs/op
BenchmarkGetByteLengthOfRuneSlice/Size1000-8             214286     5479 ns/op    4384 B/op   2 allocs/op
BenchmarkGetByteLengthOfRuneSlice/Size5000-8              20654    55266 ns/op   21760 B/op   2 allocs/op
BenchmarkGetByteLengthOfRuneSlice/Size10000-8              8323   137739 ns/op   43648 B/op   2 allocs/op
BenchmarkCalculateUTF8ByteLengthForRunes/Size100-8     45554337    26.72 ns/op       0 B/op   0 allocs/op
BenchmarkCalculateUTF8ByteLengthForRunes/Size500-8      8767291    139.0 ns/op       0 B/op   0 allocs/op
BenchmarkCalculateUTF8ByteLengthForRunes/Size1000-8     4332032    268.7 ns/op       0 B/op   0 allocs/op
BenchmarkCalculateUTF8ByteLengthForRunes/Size5000-8      883483     1361 ns/op       0 B/op   0 allocs/op
BenchmarkCalculateUTF8ByteLengthForRunes/Size10000-8     439929     2747 ns/op       0 B/op   0 allocs/op
BenchmarkLocateRuneAtPosition/Size100-8                 5254884    236.1 ns/op       0 B/op   0 allocs/op
BenchmarkLocateRuneAtPosition/Size500-8                  971358     1246 ns/op       0 B/op   0 allocs/op
BenchmarkLocateRuneAtPosition/Size1000-8                 502521     2651 ns/op       0 B/op   0 allocs/op
BenchmarkLocateRuneAtPosition/Size5000-8                  73554    18519 ns/op       0 B/op   0 allocs/op
BenchmarkLocateRuneAtPosition/Size10000-8                 22268    57601 ns/op       0 B/op   0 allocs/op
```

### Key Highlights

- **Speed**: Our implementation executes nearly instantaneously, even with large inputs.
- **Simplicity**: By avoiding unnecessary casts or conversions, the code remains straightforward and highly maintainable.
- **Efficiency**: With zero memory allocation, optimal performance is ensured without wasting resources.

This performance is particularly notable in scenarios where high efficiency and minimal overhead are required, making our solution ideal for high-performance applications.
