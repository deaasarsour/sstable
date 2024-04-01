package util

import (
	"sstable/dbms/storage"
	"sstable/filesystem"
	"sstable/memtable"
	"sstable/sstable"
)

func CreateSStable(memtable *memtable.MemoryTable, storageDir storage.StorageDirectories) (string, error) {
	if sstableFile, sstableFilename, err := CreateSStableFile(storageDir); err == nil {
		sstable.FlushSSTable(memtable, sstableFile)
		return sstableFilename, nil
	} else {
		return "", err
	}
}

func CreateSStableFile(storageDir storage.StorageDirectories) (filesystem.FileOperation, string, error) {
	memtableDirectory := storageDir.GetSStableDirectory()
	memtableFilename := generateSSTableName()
	return createDBFile(memtableFilename, memtableDirectory)
}
