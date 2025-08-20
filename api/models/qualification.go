package models

type Qualification struct {
	ID       int
	WorkerID int
	CertName string
}

type QualificationCreateRequest struct {
	WorkerID int
	CertName string
}
