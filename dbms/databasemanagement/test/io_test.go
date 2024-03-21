package databasemanagement_test

import (
	testdbms "sstable/test/util/dbms"
	"sstable/test/util/mockmemtable"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteRead(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)
	testdbms.InitializeDbmsPartially(dbms)
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
	testdbms.InitializeDbmsPartially(dbms)
	testdbms.UpdateMemtable(dbms, mockmemtable.NewAlmostFullMemtable())

	databaseManagement := dbms.DatabaseManagement

	//act
	databaseManagement.Write("name", "deea")
	readResult, _ := databaseManagement.Read("name")
	state := dbms.StateManagement.GetState()

	//assert
	assert.Equal(t, "deea", readResult)
	assert.Equal(t, 1, len(state.FulledMemoryTables))
}
