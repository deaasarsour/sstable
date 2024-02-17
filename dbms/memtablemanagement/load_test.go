package memtablemanagement_test

import (
	testdbms "sstable/test/util/dbms"
	"sstable/test/util/mockfilesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemtableEmpty(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)

	//act
	dbms.Initialize()
	memtable := dbms.MemtableManagement.GetMemtable()

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
	memtable := dbms.MemtableManagement.GetMemtable()
	memtable.Write(key, value)

	dbms = testdbms.NewDummyDbmsDirectory(rootDirectory, nil)
	dbms.StateManagement.LoadMetadata()

	memtableFolder := dbms.Storage.GetMemtableDirectory()
	memtableFiles, _ := memtableFolder.GetFiles()
	memtableFilesCount := len(memtableFiles)

	dbms.MemtableManagement.LoadMemtable()
	memtable = dbms.MemtableManagement.GetMemtable()

	//assert
	assert.Equal(t, value, memtable.Read(key))
	assert.NotNil(t, memtable)
	assert.Equal(t, 1, memtableFilesCount)
}
