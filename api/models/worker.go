package models

import (
	"fmt"
	"time"
)

type Worker struct {
	ID             int
	Name           string
	Email          string
	PhoneNumber    string
	Address        string
	Degree         string
	Position       string
	JobDescription string
	Department     string
	Salary         int
	BirthData      time.Time
	HiredAt        time.Time
	CreatedAt      time.Time `gorm:"createdAt"`
	DeletedAt      *time.Time
	FaceImg        Image           `gorm:"embedded;embeddedPrefix:face_"`
	IdImg          Image           `gorm:"embedded"`
	Qualifications []Qualification `gorm:"constraint:OnDelete:CASCADE"`
}

type WorkerCreateRequest struct {
	Name           string
	Email          string
	PhoneNumber    string
	Address        string
	Degree         string
	Position       string
	JobDescription string
	Department     string
	Salary         int
	BirthData      time.Time
	HiredAt        time.Time
	FaceImg        Image
	IdImg          Image
}
type WorkerUpdateRequest struct {
	Name           string
	Email          string
	PhoneNumber    string
	Address        string
	Degree         string
	Position       string
	JobDescription string
	Department     string
	Salary         int
	BirthData      time.Time
	FaceImg        Image
	IdImg          Image
}

type GetWorkerDetailsDto struct {
	ID             int
	Name           string
	Email          string
	PhoneNumber    string
	Address        string
	Degree         string
	Position       string
	JobDescription string
	Department     string
	Salary         int
	BirthData      time.Time
	HiredAt        time.Time
	FaceImg        Image
	IdImg          Image
	Qualifications []Qualification
}
type GetWorkerDto struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
	Position    string
	Department  string
	Salary      int
	FaceImg     Image
}

func (w Worker) DisplayWorkerCmd() {
	fmt.Println("Worker: ", w.ID, " - ", w.Name)
	fmt.Println("Contact: ", w.Email, " - ", w.PhoneNumber, " - ", w.Address)
	fmt.Println("Job: ", w.Position, " - ", w.Department, " - ", w.JobDescription)
	fmt.Println("Imgs: \n", w.FaceImg, "\n", w.IdImg)
	fmt.Println("Qualifications: ")
	for i, q := range w.Qualifications {
		fmt.Println(i, ") ", q)
	}
}
