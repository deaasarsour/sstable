package databasemanagement

import (
	"sstable/dbms/state"
)

func (databaseManagement *DatabaseManagement) switchAndFlushMemtable(curState *state.DatabaseManagementState) error {

	if newMemtable, newMemtableFilename, err := databaseManagement.createMemtable(); err == nil {

		exec := func(state *state.DatabaseManagementState) error {
			state.SwitchAndFlushFullMemtable(newMemtable, newMemtableFilename)
			return nil
		}

		return databaseManagement.stateManagement.StateUpdateSafe(exec)

	} else {
		return err
	}
}
