package state

import "sstable/util/sliceutil"

type DatabaseMetadata struct {
	MemtableFilename        string   `json:"memtable_filename"`
	FulledMemtableFilenames []string `json:"fulled_memtable_filenames"`
	MemtableToSSTable       []string `json:"memtable_to_sstable"`
}

func (metadata *DatabaseMetadata) DeepCopy() *DatabaseMetadata {
	return &DatabaseMetadata{
		MemtableFilename:        metadata.MemtableFilename,
		FulledMemtableFilenames: sliceutil.CopyArray(metadata.FulledMemtableFilenames),
		MemtableToSSTable:       sliceutil.CopyArray(metadata.MemtableToSSTable),
	}
}
