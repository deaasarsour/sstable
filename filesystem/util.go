package filesystem

import (
	"github.com/samber/lo"
)

func GetOrCreateDirectory(directoryOperation DirectoryOperation, directoryName string) (DirectoryOperation, error) {
	directories, err := directoryOperation.GetDirectories()

	if err != nil {
		return nil, err
	}

	if lo.Contains(directories, directoryName) {
		if directory, err := directoryOperation.GetDirectory(directoryName); err == nil {
			return directory, nil
		} else {
			return nil, err
		}
	} else {
		if directory, err := directoryOperation.CreateDirectory(directoryName); err == nil {
			return directory, nil
		} else {
			return nil, err
		}
	}
}
