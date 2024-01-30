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

	read1, read1Error := fileOp.ReadAll()

	read2Bytes := []byte("0000")
	read2Count, read2Error := fileOp.ReadAt(read2Bytes, 1)

	fileOp.WriteAll([]byte("hmm"))
	read3, read3Error := fileOp.ReadAll()

	//assert
	assert.Nil(t, err)

	assert.Nil(t, read1Error)
	assert.Equal(t, "Hello World :)", string(read1))

	assert.Nil(t, read2Error)
	assert.Equal(t, "ello", string(read2Bytes))
	assert.Equal(t, 4, read2Count)

	assert.Nil(t, read3Error)
	assert.Equal(t, "hmm", string(read3))

}
