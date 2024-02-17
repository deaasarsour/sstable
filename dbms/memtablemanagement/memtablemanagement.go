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
