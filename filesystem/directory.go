package filesystem

type DirectoryOperation interface {
	GetFiles() ([]string, error)
	GetFile(fileName string) (FileOperation, error)
	DeleteFile(fileName string) error
	CreateFile(fileName string, fileBytes []byte) (FileOperation, error)
	GetDirectories() ([]string, error)
	GetDirectory(directoryName string) (DirectoryOperation, error)
	CreateDirectory(directoryName string) (DirectoryOperation, error)
	// DeleteDirectory(directoryName string)
}
