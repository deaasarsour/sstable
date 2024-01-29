package mockmemtable

import (
	"sstable/filesystem"
	"sstable/memtable"
	"sstable/test/util/mockfilesystem"
)

const BASIC_TEST_DATA = "basic_1.log"
const CORRUPTED_TEST_DATA = "corrupted_memtable.log"

func NewMemtable(dataFileName string) *memtable.MemoryTable {
	dummyFile := mockfilesystem.NewDummyFileFromMemtableFolder(dataFileName)
	memtableInstance := memtable.NewMemoryTable(dummyFile)
	return memtableInstance
}

func NewReadyMemtable(dataFileName string) *memtable.MemoryTable {
	memoryTable := NewMemtable(dataFileName)
	memoryTable.LoadMemoryTable()
	return memoryTable
}

func NewReadyBasicMemtable() *memtable.MemoryTable {
	return NewReadyMemtable(BASIC_TEST_DATA)
}

func NewReadyCorruptedMemtable() (*memtable.MemoryTable, filesystem.FileOperation) {
	dummyFile := mockfilesystem.NewDummyFileFromMemtableFolder(CORRUPTED_TEST_DATA)
	memoryTable := memtable.NewMemoryTable(dummyFile)
	memoryTable.LoadMemoryTable()
	return memoryTable, dummyFile
}

func NewReadyEmptyMemtable() (*memtable.MemoryTable, filesystem.FileOperation) {
	fileOperation := mockfilesystem.NewEmptyFile()
	memoryTable := memtable.NewMemoryTable(fileOperation)
	memoryTable.LoadMemoryTable()
	return memoryTable, fileOperation
}
