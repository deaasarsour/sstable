package statemanagement

import "sstable/dbms/state"

func (dbsm *DatabaseManagementStateManagement) GetState() *state.DatabaseManagementState {
	return dbsm.DatabaseState.Load()
}
func (dbsm *DatabaseManagementStateManagement) SetState(state *state.DatabaseManagementState) {
	dbsm.DatabaseState.Store(state)
}
func (dbsm *DatabaseManagementStateManagement) IsState() bool {
	return dbsm.GetState() != nil
}
