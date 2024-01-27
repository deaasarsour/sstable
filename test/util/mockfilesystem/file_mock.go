package mockfilesystem

type DummyFile struct {
	content string
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
