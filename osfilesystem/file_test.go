package filesystemos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsFile(t *testing.T) {
	//arrange
	tempFile, err := os.CreateTemp("", "*")
	filePath := tempFile.Name()
	tempFile.Close()

	//act
	fileOp := NewOsFile(filePath)
	fileOp.Open()

	fileOp.AppendBytes([]byte("Hello"))
	fileOp.AppendBytes([]byte(" World :)"))

	rd1, e1 := fileOp.ReadAll()

	rd2 := []byte("0000")
	rd2c, e2 := fileOp.ReadAt(rd2, 1)

	fileOp.WriteAll([]byte("hmm"))
	rd3, e3 := fileOp.ReadAll()

	//assert
	assert.Nil(t, err)

	assert.Nil(t, e1)
	assert.Equal(t, "Hello World :)", string(rd1))

	assert.Nil(t, e2)
	assert.Equal(t, "ello", string(rd2))
	assert.Equal(t, 4, rd2c)

	assert.Nil(t, e3)
	assert.Equal(t, "hmm", string(rd3))

}
