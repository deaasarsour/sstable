package dbms

import (
	"sstable/filesystem"
	"sstable/memtable"
	"sstable/util"
)

func generateMemtableName() string {
	return util.CreateULID("memtable")
}

func (dbms *DatabaseManagementSystem) loadMemtableFromFile(memtableFile filesystem.FileOperation) error {
	memtable := memtable.NewMemoryTable(memtableFile)
	if err := memtable.LoadMemoryTable(); err == nil {
		dbms.memoryTable.Store(memtable)
		return nil
	} else {
		return err
	}

}

func (dbms *DatabaseManagementSystem) createMemtable() error {
	memtableFilename := generateMemtableName()
	memtableDirectory := dbms.storage.GetMemtableDirectory()

	if memtableFile, err := memtableDirectory.CreateFile(memtableFilename, nil); err == nil {
		if err := dbms.loadMemtableFromFile(memtableFile); err == nil {
			cachedMetadata := dbms.cachedMetadata.Load()
			cachedMetadata.MemtableFilename = memtableFilename
			return dbms.writeMetadataUnsafe()
		} else {
			return err
		}
	} else {
		return err
	}
}

func (dbms *DatabaseManagementSystem) loadMemtable() error {
	cachedMetadata := dbms.cachedMetadata.Load()

	memtableFilename := cachedMetadata.MemtableFilename
	memtableDirectory := dbms.storage.GetMemtableDirectory()

	if memtableFile, err := memtableDirectory.GetFile(memtableFilename); err == nil {
		return dbms.loadMemtableFromFile(memtableFile)
	} else {
		return err
	}

}

func (dbms *DatabaseManagementSystem) initializeMemtable() error {
	cachedMetadata := dbms.cachedMetadata.Load()
	if memtableFilename := cachedMetadata.MemtableFilename; memtableFilename == "" {
		return dbms.createMemtable()
	} else {
		return dbms.loadMemtable()
	}
}
