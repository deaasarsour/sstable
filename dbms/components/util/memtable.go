package util

import (
	"sstable/dbms/storage"
	"sstable/filesystem"
)

func CreateMemtableFile(storageDir storage.StorageDirectories) (filesystem.FileOperation, string, error) {
	memtableDirectory := storageDir.GetMemtableDirectory()
	memtableFilename := generateMemtableName()
	return createDBFile(memtableFilename, memtableDirectory)
}
