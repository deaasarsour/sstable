package statemanagement

import (
	"encoding/json"
	"sstable/dbms/state"
)

func (stateManagement *DatabaseManagementStateManagement) LoadMetadata() error {

	exec := func(dbState *state.DatabaseManagementState) error {
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
