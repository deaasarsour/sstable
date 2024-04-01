package core

import (
	"sstable/dbms/components/memtablemanagement"
	databasereader "sstable/dbms/ioreader"
	"sstable/dbms/iowriter/fullmemtableflusher"
	"sstable/dbms/iowriter/memtablewriter"
	"sstable/dbms/statemanagement"
	"sstable/dbms/storage"
	"sstable/filesystem"
)

type DatabaseManagementSystem struct {
	Storage             *storage.StorageState
	StateManagement     *statemanagement.DatabaseManagementStateManagement
	MemtableWriterJob   *memtablewriter.MemtableWriterJob
	MemtableManagement  *memtablemanagement.MemtableManagement
	FullMemtableFlusher *fullmemtableflusher.FullMemtableFlusher
	DatabaseReader      *databasereader.DatabaseReader
}

func NewDatabaseManagedSystemFromStorage(storage *storage.StorageState) *DatabaseManagementSystem {
	stateManagement := statemanagement.NewDatabaseStateManager(storage)
	memtableManagement := memtablemanagement.NewMemtableManagement(storage, stateManagement)
	memtableWriterJob := memtablewriter.NewMemtableWriteJob(stateManagement, memtableManagement)
	fullMemtableFlusher := fullmemtableflusher.NewFullMemtableFlusher(storage, stateManagement)
	databaseReader := databasereader.NewDatabaseReader(storage, stateManagement, memtableWriterJob)
	dbms := &DatabaseManagementSystem{
		Storage:             storage,
		StateManagement:     stateManagement,
		DatabaseReader:      databaseReader,
		MemtableWriterJob:   memtableWriterJob,
		MemtableManagement:  memtableManagement,
		FullMemtableFlusher: fullMemtableFlusher,
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
	if err := dbms.MemtableManagement.Initialize(); err != nil {
		return err
	}

	dbms.MemtableWriterJob.Initialize()
	dbms.FullMemtableFlusher.Initialize()

	return nil
}
