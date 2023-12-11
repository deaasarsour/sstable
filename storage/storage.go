package storage

import (
	"sstable/filesystem"
	"sstable/util"
)

const METADATA_FOLDER_NAME = "metadata"
const MEMTABLE_FOLDER = "memtable"

type StorageState struct {
	rootDirectory     filesystem.DirectoryOperation
	metadataDirectory filesystem.DirectoryOperation
	memtable          filesystem.DirectoryOperation
}

func assignFolder(rootDirectory filesystem.DirectoryOperation, assignDirectory *filesystem.DirectoryOperation, fileName string) error {
	if metadataDirectory, err := filesystem.GetOrCreateDirectory(rootDirectory, fileName); err == nil {
		*assignDirectory = metadataDirectory
		return nil
	} else {
		return err
	}
}

func (storageState *StorageState) createMetadataFolder() error {
	return assignFolder(storageState.rootDirectory, &storageState.metadataDirectory, METADATA_FOLDER_NAME)
}
func (storageState *StorageState) createMemtableFolder() error {
	return assignFolder(storageState.rootDirectory, &storageState.memtable, MEMTABLE_FOLDER)
}

func NewStorageState(rootDirectory filesystem.DirectoryOperation) (*StorageState, error) {
	storageState := &StorageState{
		rootDirectory: rootDirectory,
	}
	if err := util.TryRunAll(
		storageState.createMetadataFolder,
		storageState.createMemtableFolder,
	); err == nil {
		return storageState, nil
	} else {
		return nil, err
	}
}

func (storageState *StorageState) GetRootDirectory() filesystem.DirectoryOperation {
	return storageState.rootDirectory
}

func (storageState *StorageState) GetMemtableDirectory() filesystem.DirectoryOperation {
	return storageState.memtable
}
