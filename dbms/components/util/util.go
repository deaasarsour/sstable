package util

import (
	"sstable/filesystem"
	"sstable/util"
)

func generateMemtableName() string {
	return util.CreateULID("memtable")
}
func generateSSTableName() string {
	return util.CreateULID("sstable")
}

func createDBFile(name string, directory filesystem.DirectoryOperation) (filesystem.FileOperation, string, error) {
	if file, err := directory.CreateFile(name, nil); err == nil {
		return file, name, nil
	} else {
		return nil, "", err
	}
}
