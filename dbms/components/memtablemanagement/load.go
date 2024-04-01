package memtablemanagement

import (
	"sstable/dbms/components/util"
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

func (memtableManagement *MemtableManagement) createMemtable() (*memtable.MemoryTable, string, error) {
	if memtableFile, memtableFilename, err := util.CreateMemtableFile(memtableManagement.storageDir); err == nil {
		if memtable, err := createReadyMemtable(memtableFile); err == nil {
			return memtable, memtableFilename, nil
		} else {
			return nil, "", err
		}
	} else {
		return nil, "", err
	}
}

func (memtableManagement *MemtableManagement) updateDbmsMemtable(memtableFile filesystem.FileOperation, memtableFilename string) error {
	stateManagement := memtableManagement.stateManagement
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

func (memtableManagement *MemtableManagement) createAndLoadMemtable() error {
	storageDir := memtableManagement.storageDir
	if memtableFile, memtableFilename, err := util.CreateMemtableFile(storageDir); err == nil {
		return memtableManagement.updateDbmsMemtable(memtableFile, memtableFilename)
	} else {
		return err
	}
}

func (memtableManagement *MemtableManagement) loadMemtable(memtableFilename string) error {
	storageDir := memtableManagement.storageDir
	memtableDirectory := storageDir.GetMemtableDirectory()
	if memtableFile, err := memtableDirectory.GetFile(memtableFilename); err == nil {
		return memtableManagement.updateDbmsMemtable(memtableFile, memtableFilename)
	} else {
		return err
	}
}

func (memtableManagement *MemtableManagement) Initialize() error {
	stateGetOperation := memtableManagement.stateManagement
	metadata := stateGetOperation.GetState().Metadata

	if memtableFilename := metadata.MemtableFilename; memtableFilename == "" {
		return memtableManagement.createAndLoadMemtable()
	} else {
		return memtableManagement.loadMemtable(memtableFilename)
	}
}
