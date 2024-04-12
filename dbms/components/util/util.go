package util

import (
	"sstable/filesystem"
	randomutil "sstable/util/random"
)

func generateMemtableName() string {
	return randomutil.CreateULID("memtable")
}
func generateSSTableName() string {
	return randomutil.CreateULID("sstable")
}

func createDBFile(name string, directory filesystem.DirectoryOperation) (filesystem.FileOperation, string, error) {
	if file, err := directory.CreateFile(name, nil); err == nil {
		return file, name, nil
	} else {
		return nil, "", err
	}
}
