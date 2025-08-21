package models

import (
	"fmt"
	"time"

	"github.com/hasona23/workis/api/helpers"
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
	BirthDate      time.Time
	HiredAt        time.Time
	CreatedAt      time.Time `gorm:"createdAt"`
	DeletedAt      *time.Time
	FaceImg        *Image          `gorm:"embedded;embeddedPrefix:face_"`
	IdImg          *Image          `gorm:"embedded"`
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
	BirthDate      time.Time
	HiredAt        time.Time
	FaceImg        *Image `gorm:"embedded;embeddedPrefix:face_"`
	IdImg          *Image `gorm:"embedded"`
}
type WorkerUpdateRequest struct {
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
	FaceImg        *Image
	IdImg          *Image
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
	FaceImg     *Image
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

func (w WorkerCreateRequest) ValidateCreateWorkerRequest() error {
	return helpers.Validate(
		func() error {
			return helpers.IsStringInBounds(w.Name, "Name", 3, 32)
		},
		func() error {
			return helpers.IsValidEmail(w.Email)
		},
		func() error {
			return helpers.IsValidPhoneNumber(w.PhoneNumber)
		},
		func() error {
			return helpers.IsStringInBounds(w.Address, "Address", 8, 128)
		},
		func() error {
			return helpers.IsStringInBounds(w.Degree, "Degree", 3, 32)
		},
		func() error {
			return helpers.IsStringInBounds(w.Position, "Position", 2, 32)
		},
		func() error {
			return helpers.IsStringInBounds(w.JobDescription, "JobDescription", 2, 128)
		},
		func() error {
			return helpers.IsStringInBounds(w.Department, "Department", 2, 32)
		},
		func() error {
			if w.Salary < 0 {
				return fmt.Errorf("salary must be non-negative")
			}
			return nil
		},
		func() error {
			if w.BirthDate.IsZero() {
				return fmt.Errorf("birthData must be set")
			}
			return nil
		},
		func() error {
			if w.HiredAt.IsZero() {
				return fmt.Errorf("hiredAt must be set")
			}
			return nil
		},
	)
}
func (w WorkerUpdateRequest) ValidateWorkerUpdateRequest() error {
	return helpers.Validate(
		func() error {
			return helpers.IsStringInBounds(w.Name, "Name", 3, 32)
		},
		func() error {
			return helpers.IsValidEmail(w.Email)
		},
		func() error {
			return helpers.IsValidPhoneNumber(w.PhoneNumber)
		},
		func() error {
			return helpers.IsStringInBounds(w.Address, "Address", 8, 128)
		},
		func() error {
			return helpers.IsStringInBounds(w.Degree, "Degree", 3, 32)
		},
		func() error {
			return helpers.IsStringInBounds(w.Position, "Position", 2, 32)
		},
		func() error {
			return helpers.IsStringInBounds(w.JobDescription, "JobDescription", 2, 128)
		},
		func() error {
			return helpers.IsStringInBounds(w.Department, "Department", 2, 32)
		},
		func() error {
			if w.Salary < 0 {
				return fmt.Errorf("salary must be non-negative")
			}
			return nil
		},
	)
}
