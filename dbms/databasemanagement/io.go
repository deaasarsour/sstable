package databasemanagement

import "sstable/sstable"

func (databaseManagement *DatabaseManagement) Read(key string) (any, error) {
	state := databaseManagement.stateGetOperation.GetState()

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

func (databaseManagement *DatabaseManagement) Write(key string, value any) error {
	databaseManagement.mutex.Lock()
	defer databaseManagement.mutex.Unlock()

	state := databaseManagement.stateGetOperation.GetState()

	if err := state.MemoryTable.Write(key, value); err == nil {
		if state.MemoryTable.IsFull() {
			return databaseManagement.switchAndFlushMemtable(state)
		} else {
			return nil
		}
	} else {
		return err
	}
}
