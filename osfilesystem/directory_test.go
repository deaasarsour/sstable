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

	_, e1 := directory.CreateFile("sstable.txt", []byte("Hey"))
	directoryA, e2 := directory.CreateDirectory("a")
	_, e3 := directory.CreateDirectory("b")

	_, e4 := directoryA.CreateDirectory("ia")
	_, e5 := directoryA.CreateFile("sstableI.txt", []byte("No"))

	directories, e6 := directory.GetDirectories()
	files, e7 := directory.GetFiles()

	//assert
	assert.Nil(t, err)
	assert.Nil(t, e1)
	assert.Nil(t, e2)
	assert.Nil(t, e3)
	assert.Nil(t, e4)
	assert.Nil(t, e5)
	assert.Nil(t, e6)
	assert.Nil(t, e7)

	assert.Equal(t, []string{"a", "b"}, directories)
	assert.Equal(t, []string{"sstable.txt"}, files)

}
