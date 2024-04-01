package memtable

import (
	"sstable/filesystem"
	"sync/atomic"
)

type MemoryTableIO interface {
	Read(key string) any
	Write(key string, value any) error
}

type MemoryTableLowLevel interface {
	GetRecords() map[string]any
	IsLoaded() bool
	LoadMemoryTable() error
}

type MemoryTable struct {
	file      filesystem.FileOperation
	records   map[string]any
	isLoaded  bool
	flushLock atomic.Bool
}

func NewMemoryTable(file filesystem.FileOperation) *MemoryTable {
	records := make(map[string]any)
	return &MemoryTable{file: file, records: records}
}
