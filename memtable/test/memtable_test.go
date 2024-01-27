package testmemtable

import (
	"encoding/json"

	"sstable/test/util/mockmemtable"
	"sstable/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const BASIC_TEST_DATA = "basic_1.log"

func getKeyValueJson(key string, value any) string {

	keyValue := util.KeyValueObject{Key: key, Value: value}
	bytes, err := json.Marshal(keyValue)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func TestMemtableReadInt(t *testing.T) {
	//arrange
	expectValue := 14

	//act
	sut := mockmemtable.NewReadyBasicMemtable()
	result := sut.Read("score#deea")

	//assert
	assert.EqualValues(t, expectValue, result)
}

type dict = map[string]any

func TestMemtableReadString(t *testing.T) {
	//arrange
	expectValue := "deeax99"

	//act
	sut := mockmemtable.NewReadyBasicMemtable()
	result := sut.Read("nickname#deea")

	//assert
	assert.EqualValues(t, expectValue, result)
}

func TestMemtableReadObject(t *testing.T) {
	//arrange
	expectValue := dict{
		"name":    "deea",
		"user_id": 12.0,
		"metadata": dict{
			"last_login": 7.0,
		},
	}

	//act
	sut := mockmemtable.NewReadyBasicMemtable()
	result := sut.Read("profile#deea")

	//assert
	assert.Equal(t, expectValue, result)
}

func TestMemtableWrite(t *testing.T) {
	//arrange
	memoryTable, dummyFile := mockmemtable.NewReadyEmptyMemtable()
	key := "user_score#deeax99"
	value := 10.0
	expectedJson := getKeyValueJson(key, value)

	//act
	memoryTable.Write(key, value)
	memtableValue := memoryTable.Read(key)
	bytes, _ := dummyFile.ReadAll()
	content := string(bytes)

	//assert
	assert.Equal(t, value, memtableValue)
	assert.True(t, strings.Contains(content, expectedJson))
}
