package databasemanagement

import (
	"sstable/memtable"
	"sstable/sstable"
)

func (databaseManagement *DatabaseManagement) createSStable(memtable *memtable.MemoryTable) (string, error) {
	if sstableFile, sstableFilename, err := databaseManagement.createSStableFile(); err == nil {
		sstable.FlushSSTable(memtable, sstableFile)
		return sstableFilename, nil
	} else {
		return "", err
	}
}
