package memtablemanagement

import (
	"sstable/filesystem"
	"sstable/memtable"
	"sstable/util"
)

func generateMemtableName() string {
	return util.CreateULID("memtable")
}

func createReadyMemtable(memtableFile filesystem.FileOperation) (*memtable.MemoryTable, error) {
	memtable := memtable.NewMemoryTable(memtableFile)
	if err := memtable.LoadMemoryTable(); err == nil {
		return memtable, nil
	} else {
		return nil, err
	}
}

func (memtableManagement *DatabaseMemtableManagement) updateDbmsMemtable(memtableFile filesystem.FileOperation, memtableFilename string) error {
	stateManagement := memtableManagement.stateManagement
	if memtable, err := createReadyMemtable(memtableFile); err == nil {
		return stateManagement.SetMemtable(memtable, memtableFilename)
	} else {
		return nil
	}
}

func (memtableManagement *DatabaseMemtableManagement) createMemtable() error {
	storageDir := memtableManagement.storageDir
	memtableDirectory := storageDir.GetMemtableDirectory()
	memtableFilename := generateMemtableName()
	if memtableFile, err := memtableDirectory.CreateFile(memtableFilename, nil); err == nil {
		return memtableManagement.updateDbmsMemtable(memtableFile, memtableFilename)
	} else {
		return err
	}
}

func (memtableManagement *DatabaseMemtableManagement) loadMemtable(memtableFilename string) error {
	storageDir := memtableManagement.storageDir
	memtableDirectory := storageDir.GetMemtableDirectory()
	if memtableFile, err := memtableDirectory.GetFile(memtableFilename); err == nil {
		return memtableManagement.updateDbmsMemtable(memtableFile, memtableFilename)
	} else {
		return err
	}
}

func (memtableManagement *DatabaseMemtableManagement) LoadMemtable() error {
	stateGetOperation := memtableManagement.stateGetOperation
	metadata := stateGetOperation.GetState().Metadata

	if memtableFilename := metadata.MemtableFilename; memtableFilename == "" {
		return memtableManagement.createMemtable()
	} else {
		return memtableManagement.loadMemtable(memtableFilename)
	}
}
