package state

import (
	"sstable/memtable"
	"sstable/util/sliceutil"
)

type DatabaseManagementState struct {
	MemoryTable        *memtable.MemoryTable
	FulledMemoryTables []*memtable.MemoryTable
	Metadata           *DatabaseMetadata
}

func NewDatabaseState() *DatabaseManagementState {
	return &DatabaseManagementState{
		Metadata: &DatabaseMetadata{},
	}
}

func CloneDatabaseState(src *DatabaseManagementState) *DatabaseManagementState {
	return &DatabaseManagementState{
		Metadata:           src.Metadata.DeepCopy(),
		MemoryTable:        src.MemoryTable,
		FulledMemoryTables: sliceutil.CopyArray(src.FulledMemoryTables),
	}
}
