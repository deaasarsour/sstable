package memtablemanagement

import (
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
)

type DatabaseMemtableManagement struct {
	storageDir        storage.StorageDirectories
	stateManagement   *statemanagement.DatabaseManagementStateManagement
	stateGetOperation statemanagement.DatabaseManagementStateGetOperation
}

func NewMemtableManagement(
	storage storage.StorageDirectories,
	stateManagement *statemanagement.DatabaseManagementStateManagement) *DatabaseMemtableManagement {
	return &DatabaseMemtableManagement{
		storageDir:        storage,
		stateManagement:   stateManagement,
		stateGetOperation: stateManagement,
	}
}
