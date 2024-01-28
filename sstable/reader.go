package sstable

import (
	"encoding/json"
	"errors"
	"io"
	"sstable/filesystem"
)

func GetEndOfLineIndex(arr []byte) int {
	for i := range arr {
		if arr[i] == '\n' {
			return i
		}
	}

	return -1
}

func readReadLine(rowStartOffset int, file filesystem.FileOperation) ([]byte, error) {

	rowContent := make([]byte, 0)
	dest := make([]byte, 1024)

	currentOffset := rowStartOffset
	for {
		receivedBytes, err := file.ReadAt(dest, currentOffset)

		if err != nil && err != io.EOF {
			return nil, err
		}

		if receivedBytes == 0 {
			return nil, errors.New("cant find EOL")
		}
		eofIndex := GetEndOfLineIndex(dest)

		if eofIndex != -1 {
			rowContent = append(rowContent, dest[:eofIndex]...)
			break
		} else {
			rowContent = append(rowContent, dest...)
		}

		currentOffset += receivedBytes
	}

	return rowContent, nil
}

func (sstable *SSTable) readMetadata() (*sstableMetadata, int, error) {
	if metadataBytes, err := readReadLine(0, sstable.file); err == nil {
		var metadata *sstableMetadata
		json.Unmarshal(metadataBytes, &metadata)
		return metadata, len(metadataBytes) + 1, nil
	} else {
		return nil, 0, err
	}
}

func getRowsOffset(metadata *sstableMetadata, firstRecordsOffset int) []int {
	rawsCount := metadata.RowCount
	offsets := make([]int, rawsCount)
	currentOffset := firstRecordsOffset
	for i := 0; i < rawsCount; i++ {
		offsets[i] = currentOffset
		currentOffset += metadata.RowOffsets[i] + 1
	}

	return offsets
}

func (sstable *SSTable) searchFileBinarySearch(key string, metadata *sstableMetadata, firstRecordsOffset int) (any, error) {
	rowsOffset := getRowsOffset(metadata, firstRecordsOffset)
	file := sstable.file
	left := 0
	right := metadata.RowCount - 1
	for right >= left {
		mid := (left + right) / 2
		dest := make([]byte, metadata.RowOffsets[mid])
		rowBytesLen, err := file.ReadAt(dest, rowsOffset[mid])
		if err != nil {
			return nil, err
		}
		if rowBytesLen != len(dest) {
			return nil, errors.New("row isn't complete")
		}

		var row sstableRecord
		err = json.Unmarshal(dest, &row)
		if err != nil {
			return nil, err
		}

		if row.Key == key {
			return row.Value, nil
		} else if row.Key < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nil, nil
}

func (sstable *SSTable) Read(key string) (any, error) {
	if metadata, firstRecordsOffset, err := sstable.readMetadata(); err == nil {
		return sstable.searchFileBinarySearch(key, metadata, firstRecordsOffset)
	} else {
		return nil, err
	}
}
