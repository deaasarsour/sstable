package filesystem_test

import (
	mockfilesystem "sstable/filesystem/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockBasic(t *testing.T) {
	//arrange
	file1ExpectedContent := []byte("file 1 :)")
	file2ExpectedContent := []byte("file 2 :)")

	//act
	root := mockfilesystem.NewDummyDirectory()

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
	root := mockfilesystem.NewDummyDirectory()

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
	file := mockfilesystem.NewDummyFile(string(part1Content))
	file.AppendBytes(part2Content)
	actualContent, _ := file.ReadAll()

	//assert
	assert.Equal(t, expectedContent, actualContent)
}
