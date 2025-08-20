package models

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

type FileRequest struct {
	File   multipart.File
	Header *multipart.FileHeader
}

func (f *FileRequest) Close() {
	defer f.File.Close()
}
func (f FileRequest) SaveFile(dir string, name string) error {

	data, err := io.ReadAll(f.File)
	if err != nil {
		return err
	}

	filePath := path.Join(dir, name)
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
