package fullmemtableflusher

import (
	componentsutil "sstable/dbms/components/util"
	"sstable/dbms/state"
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sstable/util"

	"time"
)

type FullMemtableFlusher struct {
	stateManagement *statemanagement.DatabaseManagementStateManagement
	storageDir      storage.StorageDirectories
}

func (memtableFlusher *FullMemtableFlusher) Initialize() {
	go util.RunInLoop(memtableFlusher.FlushFulledSStable)
}

func (memtableFlusher *FullMemtableFlusher) FlushFulledSStable() {
	stateManagement := memtableFlusher.stateManagement
	curState := stateManagement.GetState()

	if len(curState.FulledMemoryTables) > 0 {
		memtable := curState.FulledMemoryTables[0]

		if sstableFilename, err := componentsutil.CreateSStable(memtable, memtableFlusher.storageDir); err == nil {

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

func NewFullMemtableFlusher(storageDir storage.StorageDirectories,
	stateManagement *statemanagement.DatabaseManagementStateManagement) *FullMemtableFlusher {
	return &FullMemtableFlusher{
		stateManagement: stateManagement,
		storageDir:      storageDir,
	}
}
