package dbms

import (
	"sstable/memtable"
	"sstable/storage"
	"sstable/util"
	"sync/atomic"
)

type DatabaseManagementSystem struct {
	memoryTable    atomic.Pointer[memtable.MemoryTable]
	storage        *storage.StorageState
	cachedMetadata atomic.Pointer[DatabaseMetadata]
}

func (dbms *DatabaseManagementSystem) initialize() error {
	return util.TryRunAll(
		dbms.initializeMetadata,
		dbms.initializeMemtable,
	)
}

func CreateDatabase(storage *storage.StorageState) (*DatabaseManagementSystem, error) {
	dbms := &DatabaseManagementSystem{
		storage: storage,
	}

	if err := dbms.initialize(); err != nil {
		return nil, err
	}

	return dbms, nil
}
