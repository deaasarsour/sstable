package sstable

import (
	"encoding/json"
	"sort"
	"sstable/filesystem"
	"sstable/memtable"
)

const SEALED = "SEALED!!!"

func NewSSTable(file filesystem.FileOperation) *SSTable {
	return &SSTable{file: file}
}

func FlushSSTable(memtable memtable.MemoryTableLowLevel, file filesystem.FileOperation) (*SSTable, error) {

	if err := loadMemtableIfNeeded(memtable); err != nil {
		return nil, err
	}

	records := memtable.GetRecords()
	sortedRecords := getRecordsSorted(records)
	serializedRecords, err := serializeSSTableRecords(sortedRecords)

	if err != nil {
		return nil, err
	}

	if err := flushSSTableMetadata(serializedRecords, file); err != nil {
		return nil, err
	}

	if err := flushSerializedRecords(serializedRecords, file); err != nil {
		return nil, err
	}

	if err := sealSSTableFile(file); err != nil {
		return nil, err
	}

	return NewSSTable(file), nil
}

func loadMemtableIfNeeded(memtable memtable.MemoryTableLowLevel) error {
	if !memtable.IsLoaded() {
		if err := memtable.LoadMemoryTable(); err != nil {
			return err
		}
	}
	return nil
}

func sealSSTableFile(file filesystem.FileOperation) error {
	return file.AppendBytes([]byte(SEALED))
}

func flushSSTableMetadata(records [][]byte, file filesystem.FileOperation) error {
	recordsCount := len(records)

	metadata := sstableMetadata{
		RowCount:   recordsCount,
		RowOffsets: make([]int, recordsCount),
	}

	for i := range records {
		metadata.RowOffsets[i] = len(records[i])
	}

	if serialized, err := json.Marshal(metadata); err == nil {
		if err := file.AppendBytes(serialized); err != nil {
			return err
		}
		if err := file.AppendBytes([]byte("\n")); err != nil {
			return err
		}
		return nil
	} else {
		return nil
	}
}

func flushSerializedRecords(records [][]byte, file filesystem.FileOperation) error {
	for _, serializedRecord := range records {
		if err := file.AppendBytes(serializedRecord); err != nil {
			return err
		}
		if err := file.AppendBytes([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}

func getRecordsSorted(records map[string]any) []*sstableRecord {
	results := make([]*sstableRecord, len(records))

	index := 0
	for k, v := range records {
		results[index] = &sstableRecord{
			Key:   k,
			Value: v,
		}
		index++
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Key < results[j].Key
	})

	return results
}

func serializeSSTableRecords(records []*sstableRecord) ([][]byte, error) {
	results := make([][]byte, len(records))
	for i, record := range records {
		if serialized, err := json.Marshal(record); err == nil {
			results[i] = serialized
		} else {
			return nil, err
		}
	}

	return results, nil
}
