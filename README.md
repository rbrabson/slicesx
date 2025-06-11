# slicesx

slicesx is a Go library that provides additional utility functions for working with slices in Go. It extends the standard library's slice functionality with useful operations.

## Installation

```bash
go get github.com/rbrabson/slicesx
```

## Usage

### Rotate

The `Rotate` function rotates a slice by the specified number of elements:

- If the number is positive, the last N elements move to the beginning.
- If the number is negative, the first N elements move to the end.

```go
package main

import (
    "fmt"

    "github.com/rbrabson/slicesx"
)

func main() {
    // Rotate integers
    nums := []int{1, 2, 3, 4, 5}

    // Rotate by 2 (positive) - last 2 elements move to front
    rotated := slicesx.Rotate(nums, 2)
    fmt.Println(rotated) // Output: [4 5 1 2 3]

    // Rotate by -1 (negative) - first element moves to end
    rotated = slicesx.Rotate(nums, -1)
    fmt.Println(rotated) // Output: [2 3 4 5 1]

    // Works with any type
    words := []string{"apple", "banana", "cherry", "date", "elderberry"}
    rotatedWords := slicesx.Rotate(words, 2)
    fmt.Println(rotatedWords) // Output: [date elderberry apple banana cherry]
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
