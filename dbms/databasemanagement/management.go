package databasemanagement

import (
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sync"
)

type DatabaseManagement struct {
	storageDir        storage.StorageDirectories
	stateManagement   *statemanagement.DatabaseManagementStateManagement
	stateGetOperation statemanagement.DatabaseManagementStateGetOperation
	mutex             sync.Mutex
}

func NewDatabaseManagement(
	storage storage.StorageDirectories,
	stateManagement *statemanagement.DatabaseManagementStateManagement) *DatabaseManagement {
	return &DatabaseManagement{
		storageDir:        storage,
		stateManagement:   stateManagement,
		stateGetOperation: stateManagement,
	}
}
