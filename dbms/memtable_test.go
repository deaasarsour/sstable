package dbms

import (
	filesystem "sstable/filesystem/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemtableEmpty(t *testing.T) {
	//arrange
	dbms := createDummyDbms(nil)
	dbms.initializeMetadata()

	//act
	dbms.initializeMemtable()
	memtable := dbms.memoryTable.Load()

	//assert
	assert.NotNil(t, memtable)
}

func TestMemtableLoad(t *testing.T) {
	//arrange
	rootDirectory := filesystem.NewDummyDirectory()
	key, value := "k", "v"

	//act
	dbms := createDummyDbmsWithDirectory(rootDirectory, nil)
	dbms.initializeMetadata()
	dbms.initializeMemtable()
	memtable := dbms.memoryTable.Load()
	memtable.Write(key, value)

	dbms = createDummyDbmsWithDirectory(rootDirectory, nil)
	dbms.initializeMetadata()

	memtableFolder := dbms.storage.GetMemtableDirectory()
	memtableFiles, _ := memtableFolder.GetFiles()
	memtableFilesCount := len(memtableFiles)

	dbms.initializeMemtable()
	memtable = dbms.memoryTable.Load()

	//assert
	assert.Equal(t, value, memtable.Read(key))
	assert.NotNil(t, memtable)
	assert.Equal(t, 1, memtableFilesCount)
}
