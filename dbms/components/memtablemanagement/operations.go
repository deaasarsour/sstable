package memtablemanagement

import (
	"sstable/dbms/state"
)

func (memtableManagement *MemtableManagement) SwitchAndFlushMemtable(curState *state.DatabaseManagementState) error {

	if newMemtable, newMemtableFilename, err := memtableManagement.createMemtable(); err == nil {

		exec := func(state *state.DatabaseManagementState) error {
			state.SwitchAndFlushFullMemtable(newMemtable, newMemtableFilename)
			return nil
		}

		return memtableManagement.stateManagement.StateUpdateSafe(exec)

	} else {
		return err
	}
}
