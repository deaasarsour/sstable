package memtablemanagement

import (
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
)

type MemtableManagement struct {
	storageDir      storage.StorageDirectories
	stateManagement *statemanagement.DatabaseManagementStateManagement
}

func NewMemtableManagement(
	storageDir storage.StorageDirectories,
	stateManagement *statemanagement.DatabaseManagementStateManagement) *MemtableManagement {
	return &MemtableManagement{
		storageDir:      storageDir,
		stateManagement: stateManagement,
	}
}
