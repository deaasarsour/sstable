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

func (dbms *DatabaseManagementStateManagement) GetAtomicState(exec func(state *state.DatabaseManagementState)) *state.DatabaseManagementState {
	dbms.getStateLock.Lock()
	defer dbms.getStateLock.Unlock()

	state := dbms.GetState()

	if exec != nil {
		exec(state)
	}

	return state
}
