package filesystemos

import (
	"os"
	"path"
	"sstable/filesystem"
)

type OsDirectory struct {
	path string
}

func (directory *OsDirectory) getFileWithType(isDir bool) ([]string, error) {
	if files, err := os.ReadDir(directory.path); err == nil {
		fileNames := []string{}
		for i := range files {
			if files[i].Type().IsDir() == isDir {
				fileNames = append(fileNames, files[i].Name())
			}
		}
		return fileNames, nil
	} else {
		return nil, err
	}

}

func NewOsDirectory(directoryPath string) filesystem.DirectoryOperation {
	return &OsDirectory{
		path: directoryPath,
	}
}

func (directory *OsDirectory) GetFiles() ([]string, error) {
	return directory.getFileWithType(false)
}

func (directory *OsDirectory) GetFile(fileName string) (filesystem.FileOperation, error) {
	fullPath := path.Join(directory.path, fileName)
	return NewOsFile(fullPath), nil
}

func (directory *OsDirectory) DeleteFile(fileName string) error {
	fullPath := path.Join(directory.path, fileName)
	return os.Remove(fullPath)
}

func (directory *OsDirectory) CreateFile(fileName string, fileBytes []byte) (filesystem.FileOperation, error) {
	fullPath := path.Join(directory.path, fileName)
	if osFile, err := os.Create(fullPath); err == nil {
		fileOp := NewOpenedOsFile(osFile)
		if err := fileOp.WriteAll(fileBytes); err == nil {
			return fileOp, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (directory *OsDirectory) GetDirectories() ([]string, error) {
	return directory.getFileWithType(true)
}

func (directory *OsDirectory) GetDirectory(directoryName string) (filesystem.DirectoryOperation, error) {
	fullPath := path.Join(directory.path, directoryName)
	return NewOsDirectory(fullPath), nil
}

func (directory *OsDirectory) CreateDirectory(directoryName string) (filesystem.DirectoryOperation, error) {
	fullPath := path.Join(directory.path, directoryName)
	if err := os.MkdirAll(fullPath, 0777); err == nil {
		return NewOsDirectory(fullPath), nil
	} else {
		return nil, err
	}
}
