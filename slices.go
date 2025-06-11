package slicesx

// Rotate rotates a slice by the specified number of elements.
// A positive num rotates the slice so that the last num elements move to the beginning.
// A negative num rotates the slice so that the last |num| elements move to the beginning.
func Rotate[T any](slice []T, num int) []T {
	length := len(slice)

	if length == 0 {
		return slice
	}

	var index int
	if num < 0 {
		index = length - ((-num) % length)
	} else {
		index = num % length
	}
	// No rotation needed
	if index == 0 {
		return slice
	}

	return append(slice[length-index:], slice[:length-index]...)
}
