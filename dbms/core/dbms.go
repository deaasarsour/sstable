package core

import (
	"sstable/dbms/databasemanagement"
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sstable/filesystem"
)

type DatabaseManagementSystem struct {
	Storage            *storage.StorageState
	StateManagement    *statemanagement.DatabaseManagementStateManagement
	DatabaseManagement *databasemanagement.DatabaseManagement
}

func NewDatabaseManagedSystemFromStorage(storage *storage.StorageState) *DatabaseManagementSystem {
	stateManagement := statemanagement.NewDatabaseStateManager(storage)
	databaseManagement := databasemanagement.NewDatabaseManagement(storage, stateManagement)

	dbms := &DatabaseManagementSystem{
		Storage:            storage,
		StateManagement:    stateManagement,
		DatabaseManagement: databaseManagement,
	}

	return dbms
}

func NewDatabaseManagedSystem(rootDirectory filesystem.DirectoryOperation) (*DatabaseManagementSystem, error) {
	if storage, err := storage.NewStorageState(rootDirectory); err == nil {
		return NewDatabaseManagedSystemFromStorage(storage), nil
	} else {
		return nil, err
	}
}

func NewReadyNewDatabaseManagedSystem(rootDirectory filesystem.DirectoryOperation) (*DatabaseManagementSystem, error) {
	if dbms, err := NewDatabaseManagedSystem(rootDirectory); err == nil {
		if err := dbms.Initialize(); err != nil {
			return nil, err
		}
		return dbms, nil
	} else {
		return nil, err
	}
}

func (dbms *DatabaseManagementSystem) Initialize() error {
	if err := dbms.StateManagement.LoadMetadata(); err != nil {
		return err
	}
	if err := dbms.DatabaseManagement.LoadMemtable(); err != nil {
		return err
	}

	dbms.DatabaseManagement.InitializeBackgroundJobs()

	return nil
}
