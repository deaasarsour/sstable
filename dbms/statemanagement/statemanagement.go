package statemanagement

import (
	"sstable/dbms/storage"
	"sync"
	"sync/atomic"
)

var mutex sync.Mutex

const UPDATE_NONE = 0
const UPDATE_ALL = 1
const UPDATE_STATE_WITHOUT_METADATA = 2

type DatabaseStateAtomicChange = func(dbState *DatabaseManagementState) error

type DatabaseManagementStateManagement struct {
	DatabaseState     atomic.Pointer[DatabaseManagementState]
	MetadataOperation storage.MetadataOperation
}

func (stateManagement *DatabaseManagementStateManagement) getUpdateType(newState *DatabaseManagementState) int {
	if stateManagement.IsState() {
		currentState := stateManagement.GetState()
		if *currentState.Metadata != *newState.Metadata {
			return UPDATE_ALL
		}
		if currentState.MemoryTable != newState.MemoryTable {
			return UPDATE_STATE_WITHOUT_METADATA
		}
	} else {
		return UPDATE_ALL
	}

	return UPDATE_NONE
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

func (stateManagement *DatabaseManagementStateManagement) updateState(newState *DatabaseManagementState, updateType int) error {
	if updateType == UPDATE_NONE {
		return nil
	}

	metadataOp := stateManagement.MetadataOperation

	if updateType == UPDATE_ALL {
		if err := metadataOp.WriteMetadata(newState.Metadata); err != nil {
			return err
		}
	}

	stateManagement.SetState(newState)
	return nil
}

func (stateManagement *DatabaseManagementStateManagement) StateUpdateSafe(exec DatabaseStateAtomicChange) error {
	mutex.Lock()
	defer mutex.Unlock()

	dbState := stateManagement.cloneState()

	if err := exec(dbState); err == nil {
		updateType := stateManagement.getUpdateType(dbState)
		return stateManagement.updateState(dbState, updateType)
	} else {
		return err
	}
}

func NewDatabaseStateManager(metadataOp storage.MetadataOperation) *DatabaseManagementStateManagement {
	return &DatabaseManagementStateManagement{
		MetadataOperation: metadataOp,
	}
}
