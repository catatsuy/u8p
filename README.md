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
