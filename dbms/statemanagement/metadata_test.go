package statemanagement_test

import (
	"sstable/dbms/state"
	testdbms "sstable/test/util/dbms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadMetadataEmpty(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)

	//act
	dbms.StateManagement.LoadMetadata()
	state := dbms.StateManagement.GetState()
	cachedMetadata := state.Metadata

	//assert
	assert.NotNil(t, cachedMetadata)
}

func TestLoadMetadataWithData(t *testing.T) {
	//arrange
	memtableFilename := "test"
	metadata := &state.DatabaseMetadata{
		MemtableFilename: memtableFilename,
	}
	dbms := testdbms.NewDummyDbms(metadata)

	//act
	dbms.StateManagement.LoadMetadata()
	state := dbms.StateManagement.GetState()
	cachedMetadata := state.Metadata

	//assert
	assert.NotNil(t, cachedMetadata)
	assert.Equal(t, memtableFilename, cachedMetadata.MemtableFilename)
}
