package mocksstable

import (
	"sstable/sstable"
	"sstable/test/util/mockfilesystem"
	"sstable/test/util/testdatafile"
)

func NewSSTable(fileName string) *sstable.SSTable {
	content := testdatafile.ReadSSTableData(fileName)
	file := mockfilesystem.NewDummyFile(content)
	return sstable.NewSSTable(file)
}
