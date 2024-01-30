package filesystemos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsDirectory(t *testing.T) {
	//arrange
	tmp, err := os.MkdirTemp("", "")

	//act
	directory := NewOsDirectory(tmp)

	_, createFileError := directory.CreateFile("sstable.txt", []byte("Hey"))
	directoryA, directoryAError := directory.CreateDirectory("a")
	_, directoryBError := directory.CreateDirectory("b")

	_, InnerDirectoryError := directoryA.CreateDirectory("ia")
	_, InnerFileError := directoryA.CreateFile("sstableI.txt", []byte("No"))

	directories, getDirectoriesError := directory.GetDirectories()
	files, getFilesError := directory.GetFiles()

	//assert
	assert.Nil(t, err)
	assert.Nil(t, createFileError)
	assert.Nil(t, directoryAError)
	assert.Nil(t, directoryBError)
	assert.Nil(t, InnerDirectoryError)
	assert.Nil(t, InnerFileError)
	assert.Nil(t, getDirectoriesError)
	assert.Nil(t, getFilesError)

	assert.Equal(t, []string{"a", "b"}, directories)
	assert.Equal(t, []string{"sstable.txt"}, files)

}
