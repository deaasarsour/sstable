package filesystem

import testdata "sstable/test/util"

type DummyFile struct {
	content string
}

func NewDummyFile(content string) *DummyFile {
	return &DummyFile{content: content}
}

func NewDummyFileFromAnotherFile(filePaths string) *DummyFile {
	content := testdata.ReadTestData(filePaths)
	return &DummyFile{content: content}
}

func (file *DummyFile) Open() error {
	return nil
}

func (file *DummyFile) Close() error {
	return nil
}

func (file *DummyFile) AppendBytes(bytes []byte) error {
	file.content = file.content + string(bytes)
	return nil
}

func (file *DummyFile) ReadAll() ([]byte, error) {
	return []byte(file.content), nil
}
