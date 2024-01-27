package mockfilesystem

import (
	"path"
	"sstable/filesystem"
	"sstable/test/util/testdatafile"
)

const MEMTABLE_TEST_DATA_FOLDER = "memtable"

func NewDummyFile(content string) *DummyFile {
	return &DummyFile{content: content}
}

func NewDummyFileFromAnotherFile(filePaths string) *DummyFile {
	content := testdatafile.ReadTestData(filePaths)
	return &DummyFile{content: content}
}

func NewDummyDirectory() *DummyDirectory {
	return &DummyDirectory{
		subDirectories: make(map[string]*DummyDirectory),
		files:          make(map[string]filesystem.FileOperation),
	}
}

func NewDummyFileFromMemtableFolder(dataFileName string) filesystem.FileOperation {
	fullPath := path.Join(MEMTABLE_TEST_DATA_FOLDER, dataFileName)
	var fileOperation filesystem.FileOperation = NewDummyFileFromAnotherFile(fullPath)
	return fileOperation
}

func NewEmptyFile() filesystem.FileOperation {
	var fileOperation filesystem.FileOperation = NewDummyFile("")
	return fileOperation
}
