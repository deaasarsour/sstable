package state_test

import (
	"sstable/dbms/state"
	"sstable/memtable"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {

	//arrange
	dbState := state.NewDatabaseState()

	dbState.MemoryTable = &memtable.MemoryTable{}
	dbState.FulledMemoryTables = []*memtable.MemoryTable{{}, {}}
	dbState.Metadata.FulledMemtableFilenames = []string{"a", "b"}
	dbState.Metadata.MemtableFilename = "c"
	dbState.Metadata.MemtableToSSTable = []string{"x", "y"}

	//act
	cloneState := state.CloneDatabaseState(dbState)

	//assert
	assert.Equal(t, dbState, cloneState)
	assert.NotSame(t, dbState, cloneState)

	assert.NotSame(t, dbState.FulledMemoryTables, cloneState.FulledMemoryTables)
	assert.Same(t, dbState.MemoryTable, cloneState.MemoryTable)
	assert.NotSame(t, dbState.Metadata, cloneState.Metadata)

	assert.NotSame(t, dbState.Metadata.FulledMemtableFilenames, cloneState.Metadata.FulledMemtableFilenames)
	assert.NotSame(t, dbState.Metadata.MemtableToSSTable, cloneState.Metadata.MemtableToSSTable)
}

func TestSwitchAndFlushMemtable(t *testing.T) {
	//arrange
	curMemtable := &memtable.MemoryTable{}
	newMemtable := &memtable.MemoryTable{}
	fulledMemtable := &memtable.MemoryTable{}

	state := &state.DatabaseManagementState{
		MemoryTable:        curMemtable,
		FulledMemoryTables: []*memtable.MemoryTable{fulledMemtable},
		Metadata: &state.DatabaseMetadata{
			MemtableFilename:        "cur",
			FulledMemtableFilenames: []string{"full"},
		},
	}
	//act
	state.SwitchAndFlushFullMemtable(newMemtable, "new")

	//assert
	assert.Equal(t, 2, len(state.FulledMemoryTables))
	assert.Same(t, fulledMemtable, state.FulledMemoryTables[0])
	assert.Same(t, curMemtable, state.FulledMemoryTables[1])

	assert.Same(t, newMemtable, state.MemoryTable)

	assert.Equal(t, state.Metadata.MemtableFilename, "new")
	assert.Equal(t, state.Metadata.FulledMemtableFilenames, []string{"full", "cur"})
}

func TestUpdateMemtable(t *testing.T) {
	//arrange
	state := state.NewDatabaseState()
	newMemtable := &memtable.MemoryTable{}

	//act
	state.UpdateMemtable(newMemtable, "new")

	//assert
	assert.Same(t, newMemtable, state.MemoryTable)
	assert.Equal(t, state.Metadata.MemtableFilename, "new")
}
