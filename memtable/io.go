package memtable

import (
	"encoding/json"
	"sstable/types"
)

const MemtableSizeCap = 1 << 20

func (memtable *MemoryTable) Read(key string) any {
	value, ok := memtable.records[key]
	if ok {
		return value
	}

	return nil
}

func (memtable *MemoryTable) Write(key string, value any) error {
	keyValue := types.KeyValueObject{Key: key, Value: value}
	bytes, err := json.Marshal(keyValue)

	if err != nil {
		return err
	}

	line := string(bytes) + "\n"

	if err = memtable.AppendBytes([]byte(line)); err != nil {
		return err
	}

	memtable.records[key] = value

	return nil
}

func (memtable *MemoryTable) WriteBatch(keyValues []types.KeyValueObject) error {
	content := make([]byte, 0)

	for _, keyValue := range keyValues {
		bytes, err := json.Marshal(keyValue)
		if err != nil {
			return err
		}

		content = append(content, bytes...)
		content = append(content, byte('\n'))
	}

	if err := memtable.AppendBytes(content); err != nil {
		return err
	}

	for _, keyValue := range keyValues {
		memtable.records[keyValue.Key] = keyValue.Value
	}
	return nil
}

func (memtable *MemoryTable) WriteBatchRaw(keyValues []types.KeyValueObject, rawBytes []byte) error {
	content := make([]byte, 0)

	for _, keyValue := range keyValues {
		bytes, err := json.Marshal(keyValue)
		if err != nil {
			return err
		}

		content = append(content, bytes...)
		content = append(content, byte('\n'))
	}

	if err := memtable.AppendBytes(content); err != nil {
		return err
	}

	for _, keyValue := range keyValues {
		memtable.records[keyValue.Key] = keyValue.Value
	}
	return nil
}

func (memtable *MemoryTable) GetRecords() map[string]any {
	return memtable.records
}

func (memtable *MemoryTable) IsFull() bool {
	return memtable.numOfBytes >= MemtableSizeCap
}

func (memtable *MemoryTable) AppendBytes(bytes []byte) error {
	if err := memtable.file.AppendBytes(bytes); err != nil {
		return err
	}

	memtable.numOfBytes += len(bytes)

	return nil
}

func (memtable *MemoryTable) WillBeFullAfterWrite(writeByteCount int) bool {
	return memtable.numOfBytes+writeByteCount >= MemtableSizeCap
}
