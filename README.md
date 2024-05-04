# u8p: UTF-8 Positioning Utility

## Overview

The `u8p` package provides functionality for working with UTF-8 encoded strings in Go. It includes a key function, `Find`, which is designed to find the index of the leading UTF-8 byte in a given string.

### Efficient Log Transmission with Go

Imagine you need to send extensive log data to another server, but sending the entire log might overwhelm the receiving server's resources, especially if multiple servers are sending logs simultaneously. To prevent such accidents, you might consider sending only a portion of the logs when they exceed a certain threshold.

The challenge then becomes deciding **where to truncate** the logs. It might seem simple to cut them off at, say, the 10,000th character, but it's not that straightforward.

In Go, if you want to truncate at the 10,000th character, casting to `rune` makes it easier:

* Casting from string to rune involves copying the entire string.
  * This defeats the purpose of truncating since copying a large string may lead to memory issues.
* Truncating at the 10,000th byte isn't suitable because characters in UTF-8 may occupy multiple bytes.
  * We don't need precision here, just a rough truncation point.

Simply using `inputString[0:20]` may seem okay, but it has its flaws:

* Characters like hiragana in UTF-8 are 3 bytes long, so arbitrary truncation may result in an invalid byte sequence.

**UTF-8 always becomes invalid if improperly truncated.**

Here's a simplified version of the information for your GitHub README:

### How to use

Here's how you can use it in your Go projects:

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

The `u8p.Find` function takes a string as its first argument and an integer as its second argument. It returns the position of the start of the UTF-8 character sequence within the string, considering the number of bytes passed as the second argument.

Internally, it searches for the position corresponding to the first byte of a UTF-8 character within the specified byte limit (second argument). If it encounters an error, it suggests that any cut at that point would result in an invalid UTF-8 sequence. In such cases, it's recommended to truncate the string approximately at that point.

Here's a practical example:

```go
if len(inputString) > maxLength {
	index, err := u8p.Find(inputString, maxLength)
	if err != nil {
		// Since any cut would result in an invalid UTF-8 sequence,
		// we truncate the string at an approximate position.
		index = maxLength
	}
	inputString = inputString[:index]
}
```

Remember that when slicing a string, the end index specifies one position before the desired end point, ensuring that the resulting string is valid UTF-8.

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

- **Speed**: Our implementation executes nearly instantaneously, even with large inputs.
- **Simplicity**: By avoiding unnecessary casts or conversions, the code remains straightforward and highly maintainable.
- **Efficiency**: With zero memory allocation, optimal performance is ensured without wasting resources.

This performance is particularly notable in scenarios where high efficiency and minimal overhead are required, making our solution ideal for high-performance applications.
