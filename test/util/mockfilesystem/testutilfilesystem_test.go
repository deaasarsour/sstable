package mockfilesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockBasic(t *testing.T) {
	//arrange
	file1ExpectedContent := []byte("file 1 :)")
	file2ExpectedContent := []byte("file 2 :)")

	//act
	root := NewDummyDirectory()

	root.CreateFile("file1.txt", file1ExpectedContent)
	root.CreateFile("file2.txt", file2ExpectedContent)

	allFiles, _ := root.GetFiles()
	file1, _ := root.GetFile("file1.txt")
	file2, _ := root.GetFile("file2.txt")

	file1Actual, _ := file1.ReadAll()
	file2Actual, _ := file2.ReadAll()

	//assert
	assert.Equal(t, allFiles, []string{"file1.txt", "file2.txt"})
	assert.Equal(t, file1Actual, file1ExpectedContent)
	assert.Equal(t, file2Actual, file2ExpectedContent)

}

func TestMockDeletion(t *testing.T) {
	//arrange
	file1ExpectedContent := []byte("file 1 :)")

	//act
	root := NewDummyDirectory()

	root.CreateFile("file1.txt", file1ExpectedContent)
	root.DeleteFile("file1.txt")

	allFiles, _ := root.GetFiles()
	_, err := root.GetFile("file1.txt")

	//assert
	assert.Equal(t, allFiles, []string{})
	assert.NotNil(t, err)
}

func TestAppend(t *testing.T) {
	//arrange
	part1Content := []byte("Hello")
	part2Content := []byte(" World!")
	expectedContent := []byte("Hello World!")

	//act
	file := NewDummyFile(string(part1Content))
	file.AppendBytes(part2Content)
	actualContent, _ := file.ReadAll()

	//assert
	assert.Equal(t, expectedContent, actualContent)
}

func TestReadAtByte(t *testing.T) {
	//arrange
	file := DummyFile{content: []byte{0x01, 0x77, 0xFF}}
	expectedDest := []byte{0x01}

	//act
	dest := make([]byte, 1)
	n, err := file.ReadAt(dest, 0)

	//assert
	assert.Equal(t, 1, n)
	assert.Equal(t, expectedDest, dest)
	assert.Nil(t, err)
}

func TestReadAtAll(t *testing.T) {
	//arrange
	file := DummyFile{content: []byte{0x01, 0x77, 0xFF}}
	expectedDest := []byte{0x01, 0x77, 0xFF}

	//act
	dest := make([]byte, 3)
	n, err := file.ReadAt(dest, 0)

	//assert
	assert.Equal(t, 3, n)
	assert.Equal(t, expectedDest, dest)
	assert.Nil(t, err)
}

func TestReadAtHalf(t *testing.T) {
	//arrange
	file := DummyFile{content: []byte{0x01, 0x77, 0xFF}}
	expectedDest := []byte{0x77, 0xFF}

	//act
	dest := make([]byte, 2)
	n, err := file.ReadAt(dest, 1)

	//assert
	assert.Equal(t, 2, n)
	assert.Equal(t, expectedDest, dest)
	assert.Nil(t, err)
}

func TestReadAtOverflow(t *testing.T) {
	//arrange
	file := DummyFile{content: []byte{0x01, 0x77, 0xFF}}

	//act
	dest := make([]byte, 2)
	n, err := file.ReadAt(dest, 10)

	//assert
	assert.Equal(t, 0, n)
	assert.NotNil(t, err)
}

func TestReadAtEOF(t *testing.T) {
	//arrange
	file := DummyFile{content: []byte{0x01, 0x77, 0xFF}}
	expectedDest := []byte{0x01, 0x77, 0xFF, 0x0, 0x0}

	//act
	dest := make([]byte, 5)
	n, err := file.ReadAt(dest, 0)

	//assert
	assert.Equal(t, 3, n)
	assert.Equal(t, expectedDest, dest)
	assert.NotNil(t, err)
}

func TestReadAtLastByte(t *testing.T) {
	//arrange
	file := DummyFile{content: []byte{0x01, 0x77, 0xFF}}
	expectedDest := []byte{0xFF}

	//act
	dest := make([]byte, 1)
	n, err := file.ReadAt(dest, 2)

	//assert
	assert.Equal(t, 1, n)
	assert.Equal(t, expectedDest, dest)
	assert.Nil(t, err)
}
