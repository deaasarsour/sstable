package dbms

import (
	"sstable/dbms/core"
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sstable/filesystem"
	"sstable/test/util/mockfilesystem"
)

func NewDummyDbmsDirectory(rootDirectory filesystem.DirectoryOperation, metadata *statemanagement.DatabaseMetadata) *core.DatabaseManagementSystem {

	storage, _ := storage.NewStorageState(rootDirectory)

	if metadata != nil {
		storage.WriteMetadata(metadata)
	}

	dbms := core.NewDatabaseManagedSystemFromStorage(storage)
	return dbms
}

func NewDummyDbms(metadata *statemanagement.DatabaseMetadata) *core.DatabaseManagementSystem {
	rootDirectory := mockfilesystem.NewDummyDirectory()
	return NewDummyDbmsDirectory(rootDirectory, metadata)
}
