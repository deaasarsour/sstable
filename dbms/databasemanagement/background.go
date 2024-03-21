package databasemanagement

import (
	"sstable/dbms/state"
	"time"
)

func runInLoop(exec func()) {
	for {
		exec()
	}
}

func (databaseManagement *DatabaseManagement) InitializeBackgroundJobs() {
	go runInLoop(databaseManagement.FlushFulledSStable)
}

func (databaseManagement *DatabaseManagement) FlushFulledSStable() {
	stateManagement := databaseManagement.stateManagement
	curState := stateManagement.GetState()

	if len(curState.FulledMemoryTables) > 0 {
		memtable := curState.FulledMemoryTables[0]
		if sstableFilename, err := databaseManagement.createSStable(memtable); err == nil {

			exec := func(dbState *state.DatabaseManagementState) error {
				dbState.FulledMemoryTables = dbState.FulledMemoryTables[1:] //this could lead into memory leak
				dbState.Metadata.MemtableToSSTable = append(dbState.Metadata.MemtableToSSTable, sstableFilename)
				return nil
			}

			stateManagement.StateUpdateSafe(exec)

		} else {
			panic(err)
		}
	}

	time.Sleep(time.Microsecond * 100)
}
