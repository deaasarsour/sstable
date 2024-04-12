package storage

import (
	"encoding/json"
	"fmt"
	"sort"
	"sstable/filesystem"
	"time"

	"github.com/samber/lo"
)

func getMetadataFromSortedFiles(sortedFiles []string, dir filesystem.DirectoryOperation) ([]byte, string) {
	var (
		metadata         []byte = nil
		metadataFileName string = ""
	)

	for _, fileName := range sortedFiles {
		fileOperation, err := dir.GetFile(fileName)
		if err == nil {
			bytes, err := fileOperation.ReadAll()
			if err == nil {
				metadata = bytes
				metadataFileName = fileName
				break
			}
		}
	}

	return metadata, metadataFileName
}

func deleteUnneededMetadata(files []string, dir filesystem.DirectoryOperation, validMetadata string) {
	for _, file := range files {
		if file != validMetadata {
			dir.DeleteFile(file)
		}
	}
}

func getNewMetadataFileName() string {
	return fmt.Sprintf("metadata_%v", time.Now().String())
}

func (storage *StorageState) ReadMetadataRaw() ([]byte, error) {
	dir := storage.metadataDirectory
	files, err := dir.GetFiles()

	if err != nil {
		return nil, err
	}

	sort.Strings(files)
	files = lo.Reverse(files)

	metadata, metadataFileName := getMetadataFromSortedFiles(files, dir)

	go deleteUnneededMetadata(files, dir, metadataFileName)

	return metadata, nil
}

func (storage *StorageState) WriteMetadata(metadata any) error {
	bytes, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	fileName := getNewMetadataFileName()

	dir := storage.metadataDirectory
	_, err = dir.CreateFile(fileName, bytes)

	return err
}
