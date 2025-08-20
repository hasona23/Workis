package models

type Qualification struct {
	ID       int
	WorkerId int
	CertName string
	CertImg  *Image `gorm:"embedded"`
}
