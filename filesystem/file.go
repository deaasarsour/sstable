package filesystem

type FileOperation interface {
	Open() error
	Close() error
	AppendBytes(bytes []byte) error
	ReadAll() ([]byte, error)
}
