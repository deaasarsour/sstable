package mockfilesystem

import (
	"errors"
	"sstable/filesystem"
)

type DummyDirectory struct {
	subDirectories map[string]*DummyDirectory
	files          map[string]filesystem.FileOperation
}

func (directory *DummyDirectory) GetFiles() ([]string, error) {

	filesName := make([]string, len(directory.files))

	i := 0
	for fileName := range directory.files {
		filesName[i] = fileName
		i++
	}

	return filesName, nil
}

func (directory *DummyDirectory) GetFile(fileName string) (filesystem.FileOperation, error) {
	if value, ok := directory.files[fileName]; ok {
		return value, nil
	}
	return nil, errors.New("file not exist")
}

func (directory *DummyDirectory) DeleteFile(fileName string) error {
	if _, ok := directory.files[fileName]; ok {
		delete(directory.files, fileName)
		return nil
	}

	return errors.New("file not exist")
}

func (directory *DummyDirectory) CreateFile(fileName string, fileBytes []byte) (filesystem.FileOperation, error) {
	if _, ok := directory.files[fileName]; !ok {
		newFile := NewDummyFile(string(fileBytes))
		directory.files[fileName] = newFile
		return newFile, nil
	}
	return nil, errors.New("file is already exist")
}

func (directory *DummyDirectory) GetDirectories() ([]string, error) {
	result := []string{}
	for key := range directory.subDirectories {
		result = append(result, key)
	}
	return result, nil
}

func (directory *DummyDirectory) GetDirectory(directoryName string) (filesystem.DirectoryOperation, error) {
	if directory, ok := directory.subDirectories[directoryName]; ok {
		return directory, nil
	} else {
		return nil, errors.New("directory not exist")
	}
}

func (directory *DummyDirectory) CreateDirectory(directoryName string) (filesystem.DirectoryOperation, error) {
	if _, ok := directory.subDirectories[directoryName]; ok {
		return nil, errors.New("directory is already exist")
	} else {
		newDirectory := NewDummyDirectory()
		directory.subDirectories[directoryName] = newDirectory
		return newDirectory, nil
	}
}
