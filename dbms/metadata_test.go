package dbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadMetadataEmpty(t *testing.T) {
	//arrange
	dbms := createDummyDbms(nil)

	//act
	dbms.initializeMetadata()
	cachedMetadata := dbms.cachedMetadata.Load()

	//assert
	assert.NotNil(t, cachedMetadata)
}

func TestLoadMetadataWithData(t *testing.T) {
	//arrange
	memtableFilename := "test"
	metadata := &DatabaseMetadata{
		MemtableFilename: memtableFilename,
	}
	dbms := createDummyDbms(metadata)

	//act
	dbms.initializeMetadata()
	cachedMetadata := dbms.cachedMetadata.Load()
	//assert
	assert.NotNil(t, cachedMetadata)
	assert.Equal(t, memtableFilename, cachedMetadata.MemtableFilename)
}
