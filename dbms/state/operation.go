package state

import (
	"sstable/memtable"
)

func (state *DatabaseManagementState) SwitchAndFlushFullMemtable(newMemtable *memtable.MemoryTable, newMemtableFilename string) {

	metadata := state.Metadata

	metadata.FulledMemtableFilenames = append(metadata.FulledMemtableFilenames, metadata.MemtableFilename)
	metadata.MemtableFilename = newMemtableFilename

	state.FulledMemoryTables = append(state.FulledMemoryTables, state.MemoryTable)
	state.MemoryTable = newMemtable
}

func (state *DatabaseManagementState) UpdateMemtable(newMemtable *memtable.MemoryTable, newMemtableFilename string) {
	state.MemoryTable = newMemtable
	state.Metadata.MemtableFilename = newMemtableFilename
}
