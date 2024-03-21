package databasemanagement

import (
	testdbms "sstable/test/util/dbms"
	"sstable/test/util/mockmemtable"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFulledMemtable(t *testing.T) {
	//arrange

	memtable := mockmemtable.NewAlmostFullMemtable()
	memtable.Write("name", "deea")
	dbms := testdbms.NewDummyDbms(nil)
	testdbms.InitializeDbmsPartially(dbms)
	testdbms.AddFullMemtable(dbms, memtable)

	//act
	databaseManagement := dbms.DatabaseManagement
	databaseManagement.FlushFulledSStable()
	state := dbms.StateManagement.GetState()
	readResult, _ := databaseManagement.Read("name")

	//assert
	assert.Equal(t, 0, len(state.FulledMemoryTables))
	assert.Equal(t, 1, len(state.Metadata.MemtableToSSTable))
	assert.Equal(t, "deea", readResult)
}
