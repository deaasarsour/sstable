package statemanagement

import (
	"sstable/dbms/state"
	"sstable/dbms/storage"
	"sync"
	"sync/atomic"
)

type DatabaseManagementStateOperation interface {
	DatabaseManagementStateGetOperation
	SetState(state *state.DatabaseManagementState)
	IsState() bool
}
type DatabaseManagementStateGetOperation interface {
	GetState() *state.DatabaseManagementState
}

type DatabaseStateAtomicChange = func(dbState *state.DatabaseManagementState) error

type DatabaseManagementStateManagement struct {
	DatabaseState     atomic.Pointer[state.DatabaseManagementState]
	MetadataOperation storage.MetadataOperation
	mutex             sync.Mutex
	getStateLock      sync.Mutex
}

func (stateManagement *DatabaseManagementStateManagement) isMetadataNeedUpdate(newState *state.DatabaseManagementState) bool {
	if stateManagement.IsState() {
		currentState := stateManagement.GetState()
		if currentState.Metadata != newState.Metadata {
			return true
		}
	} else {
		return true
	}

	return false
}

func (stateManagement *DatabaseManagementStateManagement) cloneState() *state.DatabaseManagementState {
	if stateManagement.IsState() {
		return state.CloneDatabaseState(stateManagement.GetState())
	} else {
		return state.NewDatabaseState()
	}
}

func (stateManagement *DatabaseManagementStateManagement) updateState(newState *state.DatabaseManagementState, updateMetadata bool) error {
	metadataOp := stateManagement.MetadataOperation

	if updateMetadata {
		if err := metadataOp.WriteMetadata(newState.Metadata); err != nil {
			return err
		}
	}

	stateManagement.SetState(newState)
	return nil
}

func (stateManagement *DatabaseManagementStateManagement) StateUpdateSafe(exec DatabaseStateAtomicChange) error {
	stateManagement.mutex.Lock()
	defer stateManagement.mutex.Unlock()

	dbState := stateManagement.cloneState()

	if err := exec(dbState); err == nil {
		updateMetadata := stateManagement.isMetadataNeedUpdate(dbState)
		return stateManagement.updateState(dbState, updateMetadata)
	} else {
		return err
	}
}

func NewDatabaseStateManager(metadataOp storage.MetadataOperation) *DatabaseManagementStateManagement {
	return &DatabaseManagementStateManagement{
		MetadataOperation: metadataOp,
	}
}
