package models

type Image struct {
	Path string
	Type string
	Size int64
}

const (
	PNG  = "png"
	JPEG = "jpeg"
)
