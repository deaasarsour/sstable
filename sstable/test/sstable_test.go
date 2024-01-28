package testsstable

import (
	"sstable/sstable"
	"sstable/test/util/mockfilesystem"
	"sstable/test/util/mockmemtable"
	"sstable/test/util/mocksstable"
	"sstable/test/util/testdatafile"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MEMTABLE_FILE_NAME = "memtable_to_sstable_1.log"
const SSTABLE_FILE_NAME = "sstable_1.sst"

func TestCreateSSTable(t *testing.T) {
	//arrange
	expectedContent := testdatafile.ReadSSTableData(SSTABLE_FILE_NAME)
	memtable := mockmemtable.NewReadyMemtable(MEMTABLE_FILE_NAME)
	sstableFile := mockfilesystem.NewEmptyFile()

	//act
	_, err := sstable.FlushSSTable(memtable, sstableFile)
	bytes, _ := sstableFile.ReadAll()
	actualContent := string(bytes)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedContent, actualContent)
}

func TestReadSSTable(t *testing.T) {
	//arrange
	sstable := mocksstable.NewSSTable(SSTABLE_FILE_NAME)
	fixtures := map[string]any{
		"score#a": 11.0,
		"score#c": 9.0,
		"score#e": nil,
		"score#f": 4.0,
		"score#h": nil,
	}

	for key, expectedValue := range fixtures {
		//act
		actual, err := sstable.Read(key)
		//assert
		assert.Equal(t, expectedValue, actual)
		assert.Nil(t, err)
	}
}
