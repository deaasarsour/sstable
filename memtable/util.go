package memtable

import (
	"encoding/json"
	"sstable/util"
)

func GetWriteByte(keyValue util.KeyValueObject) ([]byte, error) {
	bytes, err := json.Marshal(keyValue)
	if err != nil {
		return nil, err
	}

	return append(bytes, byte('\n')), nil
}
