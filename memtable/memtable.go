package memtable

import (
	"encoding/json"
	"sstable/filesystem"
	"sstable/util"
	"strings"
)

type MemoryTableIO interface {
	Read(key string) any
	Write(key string, value any) error
}

type MemoryTableLowLevel interface {
	GetRecords() map[string]any
	IsLoaded() bool
	LoadMemoryTable() error
}

type MemoryTable struct {
	file     filesystem.FileOperation
	records  map[string]any
	isLoaded bool
}

func NewMemoryTable(file filesystem.FileOperation) *MemoryTable {
	records := make(map[string]any)
	return &MemoryTable{file: file, records: records}
}

func (memtable *MemoryTable) Read(key string) any {
	value, ok := memtable.records[key]
	if ok {
		return value
	}

	return nil
}

func (memtable *MemoryTable) Write(key string, value any) error {
	keyValue := util.KeyValueObject{Key: key, Value: value}
	bytes, err := json.Marshal(keyValue)

	if err != nil {
		return err
	}

	line := string(bytes) + "\n"

	if err = memtable.file.AppendBytes([]byte(line)); err != nil {
		return err
	}

	memtable.records[key] = value

	return nil
}

func (memtable *MemoryTable) LoadMemoryTable() error {
	fileOp := memtable.file
	bytes, error := fileOp.ReadAll()
	if error != nil {
		return error
	}

	bytes, err := memtable.checkAndFixCorruption(bytes, memtable.file)
	if err != nil {
		return err
	}

	content := string(bytes)
	memtable.enrichRecordsFromContent(content)

	memtable.isLoaded = true
	return nil
}

func (memtable *MemoryTable) IsLoaded() bool {
	return memtable.isLoaded
}

func (memtable *MemoryTable) GetRecords() map[string]any {
	return memtable.records
}

func (memtable *MemoryTable) checkAndFixCorruption(bytes []byte, file filesystem.FileOperation) ([]byte, error) {
	if len(bytes) > 0 && bytes[len(bytes)-1] != '\n' {
		newBytes := []byte{}
		for i := len(bytes) - 1; i >= 0; i-- {
			if bytes[i] == '\n' {
				newBytes = bytes[:i+1]
				break
			}
		}
		if err := file.WriteAll(newBytes); err == nil {
			return newBytes, nil
		} else {
			return nil, err
		}
	}
	return bytes, nil
}

func (memtable *MemoryTable) enrichRecordsFromContent(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		var keyValue *util.KeyValueObject
		err := json.Unmarshal([]byte(line), &keyValue)
		if err == nil {
			memtable.records[keyValue.Key] = keyValue.Value
		}
	}
}
