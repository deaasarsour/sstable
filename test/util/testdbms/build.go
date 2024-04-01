package testdbms

import (
	"sstable/dbms/core"
	"sstable/dbms/state"
	"sstable/dbms/storage"
	"sstable/filesystem"
	"sstable/memtable"
	"sstable/test/util/mockfilesystem"
)

func NewDummyDbmsDirectory(rootDirectory filesystem.DirectoryOperation, metadata *state.DatabaseMetadata) *core.DatabaseManagementSystem {

	storage, _ := storage.NewStorageState(rootDirectory)

	if metadata != nil {
		storage.WriteMetadata(metadata)
	}

	dbms := core.NewDatabaseManagedSystemFromStorage(storage)
	return dbms
}

func NewDummyDbms(metadata *state.DatabaseMetadata) *core.DatabaseManagementSystem {
	rootDirectory := mockfilesystem.NewDummyDirectory()
	return NewDummyDbmsDirectory(rootDirectory, metadata)
}

func UpdateMemtable(dbms *core.DatabaseManagementSystem, memtable *memtable.MemoryTable) {
	dbms.StateManagement.StateUpdateSafe(func(dbState *state.DatabaseManagementState) error {
		dbState.UpdateMemtable(memtable, "-")
		return nil
	})
}

func AddFullMemtable(dbms *core.DatabaseManagementSystem, memtable *memtable.MemoryTable) {
	dbms.StateManagement.StateUpdateSafe(func(dbState *state.DatabaseManagementState) error {
		dbState.FulledMemoryTables = append(dbState.FulledMemoryTables, memtable)
		return nil
	})
}

func InitializeDbmsPartially(dbms *core.DatabaseManagementSystem) {
	dbms.StateManagement.LoadMetadata()
	dbms.MemtableManagement.Initialize()
	dbms.MemtableWriterJob.Initialize()

}
