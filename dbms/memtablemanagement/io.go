package memtablemanagement

import "sstable/memtable"

func (memtableManagement *DatabaseMemtableManagement) GetMemtable() *memtable.MemoryTable {
	stateGetOperation := memtableManagement.stateGetOperation
	state := stateGetOperation.GetState()
	return state.MemoryTable
}
