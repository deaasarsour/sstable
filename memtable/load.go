package memtable

import (
	"encoding/json"
	"sstable/filesystem"
	"sstable/types"
	"strings"
)

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
	memtable.numOfBytes = len(content)

	memtable.isLoaded = true
	return nil
}

func (memtable *MemoryTable) IsLoaded() bool {
	return memtable.isLoaded
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
		var keyValue *types.KeyValueObject
		err := json.Unmarshal([]byte(line), &keyValue)
		if err == nil {
			memtable.records[keyValue.Key] = keyValue.Value
		}
	}
}
