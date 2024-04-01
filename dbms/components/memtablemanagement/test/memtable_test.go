package databasemanagement_test

import (
	"sstable/dbms/core"
	"sstable/memtable"
	"sstable/test/util/mockfilesystem"
	testdbms "sstable/test/util/testdbms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMemtable(dbms *core.DatabaseManagementSystem) *memtable.MemoryTable {
	state := dbms.StateManagement.GetState()
	return state.MemoryTable
}

func TestMemtableEmpty(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)

	//act
	dbms.Initialize()
	memtable := getMemtable(dbms)

	//assert
	assert.NotNil(t, memtable)
}

func TestMemtableLoad(t *testing.T) {
	//arrange
	rootDirectory := mockfilesystem.NewDummyDirectory()
	key, value := "k", "v"

	//act
	dbms := testdbms.NewDummyDbmsDirectory(rootDirectory, nil)

	dbms.Initialize()
	memtable := getMemtable(dbms)
	memtable.Write(key, value)

	dbms = testdbms.NewDummyDbmsDirectory(rootDirectory, nil)
	dbms.StateManagement.LoadMetadata()

	memtableFolder := dbms.Storage.GetMemtableDirectory()
	memtableFiles, _ := memtableFolder.GetFiles()
	memtableFilesCount := len(memtableFiles)

	dbms.MemtableManagement.Initialize()
	memtable = getMemtable(dbms)

	//assert
	assert.Equal(t, value, memtable.Read(key))
	assert.NotNil(t, memtable)
	assert.Equal(t, 1, memtableFilesCount)
}
