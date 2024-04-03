package util

func Reverse[T any](arr *[]T) {
	arrLen := len(*arr)
	for i := 0; i < arrLen/2; i++ {
		tmp := (*arr)[i]
		(*arr)[i] = (*arr)[arrLen-i-1]
		(*arr)[arrLen-i-1] = tmp
	}
}

func IsContains[T comparable](arr []T, match *T) bool {
	for _, element := range arr {
		if element == *match {
			return true
		}
	}
	return false
}

func DeepCopy[T any](src []T, dest []T, startIndex, endIndex int) []T {
	len := endIndex - startIndex + 1
	for i := 0; i < len; i++ {
		dest[i] = src[i+startIndex]
	}
	return dest
}

func CopyArray[T any](src []T) []T {
	len := len(src)
	dest := make([]T, len)
	for i := 0; i < len; i++ {
		dest[i] = src[i]
	}
	return dest
}

func Combine[T any](arrays [][]T) []T {
	totalLen := 0
	for i := range arrays {
		totalLen += len(arrays[i])
	}
	res := make([]T, totalLen)
	globalIndex := 0

	for i := range arrays {
		for j := range arrays[i] {
			res[globalIndex] = arrays[i][j]
			globalIndex++
		}
	}

	return res
}
