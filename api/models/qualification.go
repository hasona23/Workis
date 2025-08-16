package models

type Qualification struct {
	ID       int    `json:"id"`
	WorkerId int    `json:"workerId"`
	CertName string `json:"certName"`
	IsActive bool   `json:"isActive"`
}
