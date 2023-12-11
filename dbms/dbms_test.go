package dbms

import (
	filesystem "sstable/filesystem"
	mockfilesystem "sstable/filesystem/mock"
	"sstable/storage"
)

func createDummyDbmsWithDirectory(rootDirectory filesystem.DirectoryOperation, metadata *DatabaseMetadata) *DatabaseManagementSystem {
	storage, _ := storage.NewStorageState(rootDirectory)

	if metadata != nil {
		storage.WriteMetadata(metadata)
	}

	dbms := &DatabaseManagementSystem{
		storage: storage,
	}
	return dbms
}

func createDummyDbms(metadata *DatabaseMetadata) *DatabaseManagementSystem {
	rootDirectory := mockfilesystem.NewDummyDirectory()
	return createDummyDbmsWithDirectory(rootDirectory, metadata)
}
