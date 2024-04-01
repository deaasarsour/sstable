package databasereader

import (
	"sstable/dbms/iowriter/memtablewriter"
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sstable/sstable"
)

type DatabaseReader struct {
	memtableWriter  *memtablewriter.MemtableWriterJob
	storageDir      storage.StorageDirectories
	stateManagement *statemanagement.DatabaseManagementStateManagement
}

func (databaseManagement *DatabaseReader) Read(key string) (any, error) {
	state := databaseManagement.stateManagement.GetState()

	if record := state.MemoryTable.Read(key); record != nil {
		return record, nil
	}

	for i := len(state.FulledMemoryTables) - 1; i >= 0; i-- {
		fulledMemtable := state.FulledMemoryTables[i]
		if record := fulledMemtable.Read(key); record != nil {
			return record, nil
		}
	}

	sstableFilenames := state.Metadata.MemtableToSSTable
	for i := len(sstableFilenames) - 1; i >= 0; i-- {
		storage := databaseManagement.storageDir
		sstableFilename := sstableFilenames[i]

		if sstableFile, err := storage.GetSStableFile(sstableFilename); err == nil {
			sstableClient := sstable.NewSSTable(sstableFile)
			if record, err := sstableClient.Read(key); err == nil {
				if record != nil {
					return record, nil
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return nil, nil
}

func (databaseManagement *DatabaseReader) Write(key string, value any) error {
	return databaseManagement.memtableWriter.Write(memtablewriter.WriteCommand{
		Key:   key,
		Value: value,
	})
}

func NewDatabaseReader(
	storage storage.StorageDirectories,
	stateManagement *statemanagement.DatabaseManagementStateManagement,
	memtableWriter *memtablewriter.MemtableWriterJob) *DatabaseReader {
	return &DatabaseReader{
		storageDir:      storage,
		stateManagement: stateManagement,
		memtableWriter:  memtableWriter,
	}
}
