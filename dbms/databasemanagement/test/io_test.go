package databasemanagement_test

import (
	"fmt"
	"sstable/dbms/core"
	"sstable/dbms/state"
	"sstable/memtable"

	testdbms "sstable/test/util/dbms"
	"sstable/test/util/mockfilesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createAlmostFullMemtable() *memtable.MemoryTable {
	dummyFile := mockfilesystem.NewEmptyFile()
	memoryTable := memtable.NewMemoryTable(dummyFile)

	for i := 0; i < memtable.MemtableFull-1; i++ {
		memoryTable.Write(fmt.Sprintf("key_%v", i), i)
	}

	return memoryTable
}

func updateMemtable(dbms *core.DatabaseManagementSystem, memtable *memtable.MemoryTable) {
	dbms.StateManagement.StateUpdateSafe(func(dbState *state.DatabaseManagementState) error {
		dbState.UpdateMemtable(memtable, "-")
		return nil
	})
}

func initializeDbms(dbms *core.DatabaseManagementSystem) {
	dbms.StateManagement.LoadMetadata()
	dbms.DatabaseManagement.LoadMemtable()
}

func TestWriteRead(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)
	initializeDbms(dbms)
	databaseManagement := dbms.DatabaseManagement

	//act
	databaseManagement.Write("name", "deea")

	readResult, _ := databaseManagement.Read("name")

	//assert
	assert.Equal(t, "deea", readResult)
}

func TestWriteReadFullMemtable(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)
	initializeDbms(dbms)
	updateMemtable(dbms, createAlmostFullMemtable())

	databaseManagement := dbms.DatabaseManagement

	//act
	databaseManagement.Write("name", "deea")
	readResult, _ := databaseManagement.Read("name")
	state := dbms.StateManagement.GetState()

	//assert
	assert.Equal(t, "deea", readResult)
	assert.Equal(t, 1, len(state.FulledMemoryTables))
}
