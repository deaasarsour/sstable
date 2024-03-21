package state

import (
	"sstable/util"
)

type DatabaseMetadata struct {
	MemtableFilename        string   `json:"memtable_filename"`
	FulledMemtableFilenames []string `json:"fulled_memtable_filenames"`
	MemtableToSSTable       []string `json:"memtable_to_sstable"`
}

func (metadata *DatabaseMetadata) DeepCopy() *DatabaseMetadata {
	return &DatabaseMetadata{
		MemtableFilename:        metadata.MemtableFilename,
		FulledMemtableFilenames: util.CopyArray(metadata.FulledMemtableFilenames),
		MemtableToSSTable:       util.CopyArray(metadata.MemtableToSSTable),
	}
}
