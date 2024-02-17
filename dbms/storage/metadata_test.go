package storage

import (
	"encoding/json"
	"sstable/test/util/mockfilesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

type metadata struct {
	MemtableName string   `json:"memtable"`
	SSTables     []string `json:"sstables"`
}

func TestMetadata(t *testing.T) {
	//arrange
	rootDirectory := mockfilesystem.NewDummyDirectory()
	storageState, _ := NewStorageState(rootDirectory)
	metadataInstance := metadata{
		MemtableName: "test_123",
		SSTables:     []string{"a", "b", "c"},
	}

	//act
	storageState.WriteMetadata(metadataInstance)
	filesCapture, _ := storageState.metadataDirectory.GetFiles()
	rawMetadata, _ := storageState.ReadMetadataRaw()
	var actualMetadata metadata
	json.Unmarshal(rawMetadata, &actualMetadata)

	//assert
	assert.Equal(t, len(filesCapture), 1)
	assert.EqualValues(t, metadataInstance, actualMetadata)
}

func TestMetadataInitializer(t *testing.T) {
	//arrange
	rootDirectory := mockfilesystem.NewDummyDirectory()
	storageState, _ := NewStorageState(rootDirectory)
	metadataInstance := metadata{
		MemtableName: "test_123",
		SSTables:     []string{"a", "b", "c"},
	}

	//act
	storageState.WriteMetadata(metadataInstance)
	storageState, _ = NewStorageState(rootDirectory)

	rawMetadata, _ := storageState.ReadMetadataRaw()
	var actualMetadata metadata
	json.Unmarshal(rawMetadata, &actualMetadata)

	//assert
	assert.EqualValues(t, metadataInstance, actualMetadata)
}

func TestMetadataEmptry(t *testing.T) {
	//arrange
	rootDirectory := mockfilesystem.NewDummyDirectory()
	storageState, _ := NewStorageState(rootDirectory)

	//act
	metadata, err := storageState.ReadMetadataRaw()

	assert.Nil(t, err)
	assert.Nil(t, metadata)
}
