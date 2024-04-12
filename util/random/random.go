package randomutil

import (
	"fmt"

	"go.bryk.io/pkg/ulid"
)

func CreateULID(prefix string) string {
	result, err := ulid.New()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v_%v", prefix, result)
}
