package filesystemos

import (
	"io"
	"os"
	"sstable/filesystem"
)

type OsFileOperation struct {
	filePath string
	osFile   *os.File
}

func NewOpenedOsFile(osFile *os.File) filesystem.FileOperation {
	return &OsFileOperation{
		osFile:   osFile,
		filePath: osFile.Name(),
	}

}

func NewOsFile(filePath string) filesystem.FileOperation {
	return &OsFileOperation{
		filePath: filePath,
	}
}

func (op *OsFileOperation) Open() error {
	if osFile, err := os.OpenFile(op.filePath, os.O_RDWR, 0600); err == nil {
		op.osFile = osFile
		return nil
	} else {
		return err
	}
}

func (op *OsFileOperation) Close() error {
	return op.osFile.Close()
}

func (op *OsFileOperation) AppendBytes(bytes []byte) error {
	if size, err := op.Size(); err == nil {
		_, err := op.osFile.WriteAt(bytes, size)
		return err
	} else {
		return err
	}
}

func (op *OsFileOperation) ReadAll() ([]byte, error) {
	return io.ReadAll(op.osFile)
}

func (op *OsFileOperation) ReadAt(dest []byte, offset int) (int, error) {
	return op.osFile.ReadAt(dest, int64(offset))
}

func (op *OsFileOperation) WriteAll(content []byte) error {
	if _, err := op.osFile.WriteAt(content, 0); err == nil {
		if err := op.osFile.Truncate(int64(len(content))); err == nil {
			_, err := op.osFile.Seek(0, 0)
			return err
		} else {
			return err
		}

	} else {
		return err
	}
}

func (op *OsFileOperation) Size() (int64, error) {
	if stat, err := op.osFile.Stat(); err == nil {
		return stat.Size(), nil
	} else {
		return 0, err
	}
}
