package models

import (
	"time"
)

const (
	MALE   = "male"
	FEMALE = "female"
)

type Worker struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	DateCreated    time.Time       `json:"dateCreated"`
	IsActive       bool            `json:"isActive"`
	Position       string          `json:"position"`
	Salary         float32         `json:"salary"`
	BirthDate      time.Time       `json:"birthDate"`
	Gender         string          `json:"gender"`
	Degree         string          `json:"degree"`
	JobDescription string          `json:"jobDescription"`
	Qualifications []Qualification `json:"qualifications"`
	// Face Img
	// ID Img
}

type CreateWorkerDto struct {
	Name           string    `json:"name"`
	Position       string    `json:"position"`
	Salary         float32   `json:"salary"`
	BirthDate      time.Time `json:"birthDate"`
	Gender         string    `json:"gender"`
	Degree         string    `json:"degree"`
	JobDescription string    `json:"jobdescription"`
}

type UpdateWorkerDto struct {
	ID             int     `json:"id"`
	Position       string  `json:"position"`
	Salary         float32 `json:"salary"`
	Degree         string  `json:"degree"`
	JobDescription string  `json:"jobdescription"`
}

func (w UpdateWorkerDto) CopyValuesToWorker(worker *Worker) {
	if w.Position != "" {
		worker.Position = w.Position
	}
	if w.Salary > 0 {
		worker.Salary = w.Salary
	}
	if w.Degree != "" {
		worker.Degree = w.Degree
	}
	if w.JobDescription != "" {
		worker.JobDescription = w.JobDescription
	}
}

func (w CreateWorkerDto) ToWorker() Worker {
	return Worker{
		Name:           w.Name,
		Position:       w.Position,
		Salary:         w.Salary,
		BirthDate:      w.BirthDate,
		Gender:         w.Gender,
		Degree:         w.Degree,
		JobDescription: w.JobDescription,
	}
}
