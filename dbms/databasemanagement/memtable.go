package databasemanagement

import (
	"sstable/dbms/state"
	"sstable/filesystem"
	"sstable/memtable"
)

func createReadyMemtable(memtableFile filesystem.FileOperation) (*memtable.MemoryTable, error) {
	memtable := memtable.NewMemoryTable(memtableFile)
	if err := memtable.LoadMemoryTable(); err == nil {
		return memtable, nil
	} else {
		return nil, err
	}
}

func (databaseManagement *DatabaseManagement) createMemtable() (*memtable.MemoryTable, string, error) {
	if memtableFile, memtableFilename, err := databaseManagement.createMemtableFile(); err == nil {
		if memtable, err := createReadyMemtable(memtableFile); err == nil {
			return memtable, memtableFilename, nil
		} else {
			return nil, "", err
		}
	} else {
		return nil, "", err
	}
}

func (databaseManagement *DatabaseManagement) updateDbmsMemtable(memtableFile filesystem.FileOperation, memtableFilename string) error {
	stateManagement := databaseManagement.stateManagement
	if memtable, err := createReadyMemtable(memtableFile); err == nil {

		exec := func(dbState *state.DatabaseManagementState) error {
			dbState.UpdateMemtable(memtable, memtableFilename)
			return nil
		}

		return stateManagement.StateUpdateSafe(exec)
	} else {
		return nil
	}
}

func (databaseManagement *DatabaseManagement) createAndLoadMemtable() error {

	if memtableFile, memtableFilename, err := databaseManagement.createMemtableFile(); err == nil {
		return databaseManagement.updateDbmsMemtable(memtableFile, memtableFilename)
	} else {
		return err
	}
}

func (databaseManagement *DatabaseManagement) loadMemtable(memtableFilename string) error {
	storageDir := databaseManagement.storageDir
	memtableDirectory := storageDir.GetMemtableDirectory()
	if memtableFile, err := memtableDirectory.GetFile(memtableFilename); err == nil {
		return databaseManagement.updateDbmsMemtable(memtableFile, memtableFilename)
	} else {
		return err
	}
}

func (databaseManagement *DatabaseManagement) LoadMemtable() error {
	stateGetOperation := databaseManagement.stateGetOperation
	metadata := stateGetOperation.GetState().Metadata

	if memtableFilename := metadata.MemtableFilename; memtableFilename == "" {
		return databaseManagement.createAndLoadMemtable()
	} else {
		return databaseManagement.loadMemtable(memtableFilename)
	}
}
