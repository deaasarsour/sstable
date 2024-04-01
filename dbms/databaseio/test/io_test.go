package databaseio_test

import (
	"sstable/test/util/mockmemtable"
	testdbms "sstable/test/util/testdbms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteRead(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)
	testdbms.InitializeDbmsPartially(dbms)
	databaseIO := dbms.DatabaseIO

	//act
	databaseIO.Write("name", "deea")

	readResult, _ := databaseIO.Read("name")

	//assert
	assert.Equal(t, "deea", readResult)
}

func TestWriteReadFullMemtable(t *testing.T) {
	//arrange
	dbms := testdbms.NewDummyDbms(nil)
	testdbms.InitializeDbmsPartially(dbms)
	testdbms.UpdateMemtable(dbms, mockmemtable.NewAlmostFullMemtable())

	databaseIO := dbms.DatabaseIO

	//act
	databaseIO.Write("name", "deea")
	readResult, _ := databaseIO.Read("name")
	state := dbms.StateManagement.GetState()

	//assert
	assert.Equal(t, "deea", readResult)
	assert.Equal(t, 1, len(state.FulledMemoryTables))
}
