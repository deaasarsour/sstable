package storage

import (
	"sstable/filesystem"
	"sstable/util"
)

const METADATA_FOLDER_NAME = "metadata"
const MEMTABLE_FOLDER = "memtable"
const SSTABLE_FOLDER = "sstable"

type MetadataOperation interface {
	ReadMetadataRaw() ([]byte, error)
	WriteMetadata(metadata any) error
}
type StorageDirectories interface {
	GetRootDirectory() filesystem.DirectoryOperation
	GetMemtableDirectory() filesystem.DirectoryOperation
	GetSStableDirectory() filesystem.DirectoryOperation
	GetSStableFile(filename string) (filesystem.FileOperation, error)
}

type StorageState struct {
	rootDirectory     filesystem.DirectoryOperation
	metadataDirectory filesystem.DirectoryOperation
	memtable          filesystem.DirectoryOperation
	sstable           filesystem.DirectoryOperation
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
func (storageState *StorageState) createSStableFolder() error {
	return assignFolder(storageState.rootDirectory, &storageState.sstable, SSTABLE_FOLDER)
}

func NewStorageState(rootDirectory filesystem.DirectoryOperation) (*StorageState, error) {
	storageState := &StorageState{
		rootDirectory: rootDirectory,
	}
	if err := util.TryRunAll(
		storageState.createMetadataFolder,
		storageState.createMemtableFolder,
		storageState.createSStableFolder,
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

func (storageState *StorageState) GetSStableDirectory() filesystem.DirectoryOperation {
	return storageState.memtable
}

func (storageState *StorageState) GetSStableFile(filename string) (filesystem.FileOperation, error) {
	directory := storageState.memtable
	return directory.GetFile(filename)
}
