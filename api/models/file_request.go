package models

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"slices"
)

const MAX_IMG_SIZE = 2 * 1024 * 1024 // 2 MB

var ValidFileTypes = []string{
	mime.TypeByExtension(".png"),
	mime.TypeByExtension(".jpeg"),
	mime.TypeByExtension(".jpg"),
}

type FileRequest struct {
	Header *multipart.FileHeader
}

// Close closes the *original* multipart.File stream.
// Only call this in the handler (NOT in SaveFile).

// SaveFile safely validates & saves the uploaded file.
func (f FileRequest) SaveFile(dir string, name string) error {
	if f.Header.Size > MAX_IMG_SIZE {
		return fmt.Errorf("file must be less than %v MB", MAX_IMG_SIZE/(1024*1024))
	}

	// Open a new stream from the header
	file, err := f.Header.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// --- Validate file type ---
	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	if err := f.ValidateImgType(buf[:n]); err != nil {
		return err
	}

	// Reset to beginning
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	// --- Save file ---
	filePath := path.Join(dir, name)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return err
	}

	return nil
}

// ValidateImgType ensures the file is PNG/JPEG/JPG
func (f FileRequest) ValidateImgType(data []byte) error {
	fType := http.DetectContentType(data)
	if !slices.Contains(ValidFileTypes, fType) {
		return fmt.Errorf("only allowed file types are .png, .jpeg, .jpg")
	}
	return nil
}
