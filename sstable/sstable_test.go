package sstable

import (
	"sstable/test/util/mockfilesystem"
	"sstable/test/util/mockmemtable"
	"sstable/test/util/testdatafile"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MEMTABLE_FILE_NAME string = "memtable_to_sstable_1.log"

func TestCreateSSTable(t *testing.T) {
	//arrange
	expectedContent := testdatafile.ReadTestData("sstable/sstable_1.sst")
	mockmemtable.NewReadyMemtable(MEMTABLE_FILE_NAME)
	memtable := mockmemtable.NewReadyMemtable(MEMTABLE_FILE_NAME)
	sstableFile := mockfilesystem.NewEmptyFile()

	//act
	_, err := FlushSSTable(memtable, sstableFile)
	bytes, _ := sstableFile.ReadAll()
	actualContent := string(bytes)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedContent, actualContent)
}
