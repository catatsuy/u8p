# u8p: A Go Utility for Precise UTF-8 String Truncation

## Overview

The `u8p` package offers tools to handle UTF-8 encoded strings in Go efficiently. Its primary function, `u8p.Find`, helps identify the index of the first byte of a UTF-8 character in a string, ensuring precise truncation points.

### Efficient Log Transmission with Go

Transmitting extensive log data to another server can sometimes overwhelm its resources, especially when multiple servers are involved. To manage this, you may need to send only portions of the logs when they exceed a specific size.

However, deciding where to truncate the logs is crucial. While truncating after a certain number of characters seems straightforward, it poses several challenges:

- Casting a string to a rune array involves copying the entire string, which can lead to memory inefficiencies.
- Direct byte truncation could disrupt the integrity of UTF-8 characters, as they can occupy multiple bytes.

Using a simple slice like `inputString[0:20]` could result in invalid UTF-8 sequences if not handled correctly. UTF-8 strings must be carefully truncated to maintain valid byte sequences.

The `u8p` utility ensures that UTF-8 strings are truncated correctly, preserving both data integrity and efficiency.

Here's a refined version of the "How to use" section for the README:

### How to Use

To integrate `u8p` into your Go projects, here is a simple example:

```go
b := "Hello, 🌍. Hi!"
lb := 13
index, err := u8p.Find(b, lb)
if err != nil {
	fmt.Printf("Error: %v\n", err)
} else {
	// result: Emoji example - Hello, 🌍.
	fmt.Printf("Emoji example - %s\n", b[:index])
}
```

The function `u8p.Find` requires two parameters: a string and an integer. It returns the index where the UTF-8 character sequence begins within the string based on the byte limit provided as the second argument.

Internally, `u8p.Find` searches for the start of a UTF-8 character up to the specified byte limit. If an error occurs, it indicates that truncating at this point would produce an invalid UTF-8 sequence. In such cases, it's advisable to truncate the string near that index instead.

Here's how you might handle overly long strings:

```go
if len(inputString) > maxLength {
	index, err := u8p.Find(inputString, maxLength)
	if err != nil {
		// Adjust to the closest valid UTF-8 position.
		index = maxLength
	}
	inputString = inputString[:index]
}
```

Note: When slicing a string, the end index marks one position before the actual endpoint to ensure the integrity of the resulting UTF-8 sequence.

## Function `Find`

- **Signature**: `Find(a string, l int) (int, error)`
- **Parameters**:
  - `a`: The string to be analyzed.
  - `l`: The number of bytes from the end of `a` to be considered for analysis.
- **Returns**:
  - `int`: The index of the first UTF-8 byte found within the specified range.
  - `error`: An error indicating issues with the input parameters or the absence of a valid UTF-8 byte.
- **Description**: The `Find` function searches for the first UTF-8 byte within the last `l` bytes of string `a`. If `a` is empty, it returns 0 without an error. It triggers an error if `l` is 3 or less, if `a` is shorter than `l`, or if it cannot locate a valid UTF-8 starting byte within the specified range.

## Example Usage

For practical implementations of the `Find` function, please refer to `example_test.go`.

## Validating UTF-8 String Extraction with Fuzzing

In our project, we use Fuzzing to ensure that we can extract valid UTF-8 substrings from various strings.

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

- **Speed**: The implementation operates almost instantly, handling even large inputs with remarkable speed.
- **Simplicity**: We maintain a straightforward and clean codebase by avoiding unnecessary conversions and casts, enhancing maintainability.
- **Efficiency**: Our method ensures peak performance with zero memory allocations, optimizing resource usage.

These attributes are especially advantageous in environments demanding high efficiency and minimal operational overhead, making this solution perfect for high-performance applications.
