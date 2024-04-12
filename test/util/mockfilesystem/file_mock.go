package mockfilesystem

import (
	"io"
	"sstable/util/sliceutil"
)

type DummyFile struct {
	content []byte
}

func (file *DummyFile) Open() error {
	return nil
}

func (file *DummyFile) Close() error {
	return nil
}

func (file *DummyFile) AppendBytes(bytes []byte) error {
	file.content = append(file.content, bytes...)
	return nil
}

func (file *DummyFile) ReadAll() ([]byte, error) {
	return file.content, nil
}

func (file *DummyFile) ReadAt(dest []byte, offset int) (int, error) {
	destLen := len(dest)
	fileLen := len(file.content)

	if destLen+offset-1 < fileLen {
		sliceutil.CopySubArray[byte](file.content, dest, offset, offset+destLen-1)
		return destLen, nil
	} else if fileLen > offset {
		sliceutil.CopySubArray[byte](file.content, dest, offset, fileLen-1)
		return fileLen - offset, io.EOF
	} else {
		return 0, io.EOF
	}
}

func (file *DummyFile) WriteAll(content []byte) error {
	file.content = content
	return nil
}

func (file *DummyFile) Size() (int64, error) {
	return int64(len(file.content)), nil
}
