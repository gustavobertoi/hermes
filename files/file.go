package files

import (
	"io"
	"os"
	"path"

	"github.com/gustavobertoi/hermes/pkg"
)

type File struct {
	ID          string
	name        string
	folderPath  string
	fullPath    string
	content     []byte
	sizeInBytes int64
}

// This is a constructor for the File struct, creates a empty file with a unique ID
func NewFile(fullPath string) *File {
	id := pkg.NewUUID()
	folderPath := path.Dir(fullPath)
	fileName := path.Base(fullPath)
	return &File{
		ID:          id,
		name:        fileName,
		folderPath:  folderPath,
		fullPath:    fullPath,
		content:     make([]byte, 0),
		sizeInBytes: 0,
	}
}

func (f *File) Load() error {
	file, err := os.Open(f.fullPath)
	if err != nil {
		return err
	}
	f.name = file.Name()
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	f.content = content
	f.sizeInBytes = int64(len(content))
	return nil
}

func (f *File) Save() error {
	file, err := os.Create(f.fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(f.content)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) SetContent(content []byte) {
	f.content = content
	f.sizeInBytes = int64(len(content))
}

func (f *File) Name() string {
	return f.name
}

func (f *File) FolderPath() string {
	return f.folderPath
}

func (f *File) Path() string {
	return f.fullPath
}

func (f *File) Content() []byte {
	return f.content
}

func (f *File) SizeInBytes() int64 {
	return f.sizeInBytes
}
