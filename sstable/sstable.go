package sstable

import "sstable/filesystem"

type sstableMetadata struct {
	RowCount   int   `json:"row_count"`
	RowOffsets []int `json:"row_offsets"`
}

type sstableRecord struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type SSTable struct {
	file filesystem.FileOperation
}
