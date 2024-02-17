package statemanagement

import "sstable/memtable"

type DatabaseManagementStateOperation interface {
	DatabaseManagementStateGetOperation
	SetState(state *DatabaseManagementState)
	IsState() bool
}
type DatabaseManagementStateGetOperation interface {
	GetState() *DatabaseManagementState
}

type DatabaseManagementState struct {
	MemoryTable *memtable.MemoryTable
	Metadata    *DatabaseMetadata
}

func (dbsm *DatabaseManagementStateManagement) GetState() *DatabaseManagementState {
	return dbsm.DatabaseState.Load()
}
func (dbsm *DatabaseManagementStateManagement) SetState(state *DatabaseManagementState) {
	dbsm.DatabaseState.Store(state)
}
func (dbsm *DatabaseManagementStateManagement) IsState() bool {
	return dbsm.GetState() != nil
}

func NewDatabaseState() *DatabaseManagementState {
	return &DatabaseManagementState{
		Metadata: &DatabaseMetadata{},
	}
}
