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
