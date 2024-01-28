package testdatafile

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const TEST_DATA_ROOT = "/test/data"
const MEMTABLE_TEST_DATA_FOLDER = "memtable"
const SSTABLE_TEST_DATA_FOLDER = "sstable"

func ReadTestData(filePath string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(b))))

	fullPath := path.Join(basepath, TEST_DATA_ROOT, filePath)

	bytes, err := os.ReadFile(fullPath)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func ReadSSTableData(fileName string) string {
	fullPath := path.Join(SSTABLE_TEST_DATA_FOLDER, fileName)
	return ReadTestData(fullPath)
}

func ReadMemtableData(fileName string) string {
	fullPath := path.Join(MEMTABLE_TEST_DATA_FOLDER, fileName)
	return ReadTestData(fullPath)
}
