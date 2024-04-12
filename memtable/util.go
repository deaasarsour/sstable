package memtable

import (
	"encoding/json"
	"sstable/types"
)

func GetWriteByte(keyValue types.KeyValueObject) ([]byte, error) {
	bytes, err := json.Marshal(keyValue)
	if err != nil {
		return nil, err
	}

	return append(bytes, byte('\n')), nil
}
