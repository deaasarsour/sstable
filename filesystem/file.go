package filesystem

type FileOperation interface {
	Open() error
	Close() error
	AppendBytes(bytes []byte) error
	ReadAll() ([]byte, error)
	ReadAt(bytes []byte, offset int) (int, error)
	Size() (int, error)
}
