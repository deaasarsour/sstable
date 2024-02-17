package statemanagement

import (
	"encoding/json"
	"sstable/memtable"
)

func (stateManagement *DatabaseManagementStateManagement) SetMemtable(memtable *memtable.MemoryTable, memtableFilename string) error {

	exec := func(dbState *DatabaseManagementState) error {
		dbState.Metadata.MemtableFilename = memtableFilename
		dbState.MemoryTable = memtable
		return nil
	}

	return stateManagement.StateUpdateSafe(exec)
}

func (stateManagement *DatabaseManagementStateManagement) LoadMetadata() error {

	exec := func(dbState *DatabaseManagementState) error {
		if content, err := stateManagement.MetadataOperation.ReadMetadataRaw(); err == nil {
			if len(content) != 0 {
				if err := json.Unmarshal(content, dbState.Metadata); err != nil {
					return err
				}
			}
			return nil
		} else {
			return err
		}
	}

	return stateManagement.StateUpdateSafe(exec)
}
