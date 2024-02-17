package statemanagement

import (
	"sstable/dbms/storage"
	"sync"
	"sync/atomic"
)

type DatabaseStateAtomicChange = func(dbState *DatabaseManagementState) error

type DatabaseManagementStateManagement struct {
	DatabaseState     atomic.Pointer[DatabaseManagementState]
	MetadataOperation storage.MetadataOperation
	mutex             sync.Mutex
}

func (stateManagement *DatabaseManagementStateManagement) isMetadataNeedUpdate(newState *DatabaseManagementState) bool {
	if stateManagement.IsState() {
		currentState := stateManagement.GetState()
		if *currentState.Metadata != *newState.Metadata {
			return true
		}
	} else {
		return true
	}

	return false
}

func (stateManagement *DatabaseManagementStateManagement) cloneState() *DatabaseManagementState {
	if stateManagement.IsState() {
		newDbState := *stateManagement.GetState()
		newMetadata := *newDbState.Metadata
		newDbState.Metadata = &newMetadata
		return &newDbState
	} else {
		return NewDatabaseState()
	}
}

func (stateManagement *DatabaseManagementStateManagement) updateState(newState *DatabaseManagementState, updateMetadata bool) error {
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
