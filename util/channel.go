package util

func ReadBatch[T any](c chan T, maxBatchSize int) []T {
	res := make([]T, 0)

	res = append(res, <-c)

	for i := 0; i < maxBatchSize; i++ {
		select {
		case record := <-c:
			res = append(res, record)
		default:
			return res
		}
	}

	return res
}
