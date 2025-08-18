package model

type Qualification struct {
	ID       int
	WorkerId int
	CertName string
	*Image
}
