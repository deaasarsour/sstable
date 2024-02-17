package statemanagement_test

import (
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sstable/test/util/mockfilesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDummyStateManagement(storageState storage.MetadataOperation) *statemanagement.DatabaseManagementStateManagement {

	return statemanagement.NewDatabaseStateManager(storageState)
}

func TestStateManagement(t *testing.T) {
	//act
	rootDirectory := mockfilesystem.NewDummyDirectory()
	storageState, _ := storage.NewStorageState(rootDirectory)

	state := statemanagement.NewDatabaseState()
	stateManagement := createDummyStateManagement(storageState)

	exec := func(dbState *statemanagement.DatabaseManagementState) error {
		dbState.Metadata.MemtableFilename = "memtable"
		return nil
	}

	//arrange
	stateManagement.DatabaseState.Store(state)
	oldState := stateManagement.DatabaseState.Load()

	stateManagement.StateUpdateSafe(exec)
	newState := stateManagement.DatabaseState.Load()

	stateManagement = createDummyStateManagement(storageState)
	stateManagement.LoadMetadata()
	afterLoadState := stateManagement.DatabaseState.Load()

	//assert
	assert.Equal(t, "memtable", newState.Metadata.MemtableFilename)
	assert.Equal(t, "", oldState.Metadata.MemtableFilename)
	assert.Equal(t, "memtable", afterLoadState.Metadata.MemtableFilename)

}
