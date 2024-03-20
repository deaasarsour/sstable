package databasemanagement

import (
	"sstable/filesystem"
	"sstable/util"
)

func generateMemtableName() string {
	return util.CreateULID("memtable")
}
func generateSSTableName() string {
	return util.CreateULID("sstable")
}

func createDBFile(name string, directory filesystem.DirectoryOperation) (filesystem.FileOperation, string, error) {
	if file, err := directory.CreateFile(name, nil); err == nil {
		return file, name, nil
	} else {
		return nil, "", err
	}
}

func (databaseManagement *DatabaseManagement) createMemtableFile() (filesystem.FileOperation, string, error) {
	storageDir := databaseManagement.storageDir
	memtableDirectory := storageDir.GetMemtableDirectory()
	memtableFilename := generateMemtableName()
	return createDBFile(memtableFilename, memtableDirectory)
}

func (databaseManagement *DatabaseManagement) createSStableFile() (filesystem.FileOperation, string, error) {
	storageDir := databaseManagement.storageDir
	memtableDirectory := storageDir.GetSStableDirectory()
	memtableFilename := generateSSTableName()
	return createDBFile(memtableFilename, memtableDirectory)
}
