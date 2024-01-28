package mockfilesystem

import (
	"sstable/filesystem"
	"sstable/test/util/testdatafile"
)

func NewDummyFile(content string) *DummyFile {
	return &DummyFile{content: []byte(content)}
}

func NewDummyDirectory() *DummyDirectory {
	return &DummyDirectory{
		subDirectories: make(map[string]*DummyDirectory),
		files:          make(map[string]filesystem.FileOperation),
	}
}

func NewDummyFileFromMemtableFolder(dataFileName string) filesystem.FileOperation {
	content := testdatafile.ReadMemtableData(dataFileName)
	var fileOperation filesystem.FileOperation = NewDummyFile(content)
	return fileOperation
}

func NewEmptyFile() filesystem.FileOperation {
	var fileOperation filesystem.FileOperation = NewDummyFile("")
	return fileOperation
}
