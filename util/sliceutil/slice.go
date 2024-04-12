package sliceutil

func CopySubArray[T any](src []T, dest []T, startIndex, endIndex int) []T {
	copy(dest, src[startIndex:endIndex+1])
	return dest
}

func CopyArray[T any](src []T) []T {
	len := len(src)
	dest := make([]T, len)
	copy(dest, src)
	return dest
}
