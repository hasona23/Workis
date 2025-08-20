package services

import (
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

func CreateWorker(worker models.WorkerCreateRequest) (err error) {

	return err
}

func DeleteWorker(id int) (err error) {
	repositories.DeleteWorker(id)

	return err
}

func UpdateWorker(worker models.WorkerUpdateRequest) (err error) {

	return err
}

func GetAllWorkers() (workersGet []models.GetWorkerDto, err error) {

	workers, err := repositories.GetAllWorkers()

	workersGet = make([]models.GetWorkerDto, len(workers))
	for i, w := range workers {
		workersGet[i] = models.GetWorkerDto{
			ID:          w.ID,
			Name:        w.Name,
			Email:       w.Email,
			PhoneNumber: w.PhoneNumber,
			Position:    w.Position,
			Department:  w.Department,
			Salary:      w.Salary,
			FaceImg:     w.FaceImg,
		}
	}
	return workersGet, err
}

func GetWorkerByID(id int) (models.GetWorkerDetailsDto, error) {
	worker, err := repositories.GetWorkerWithID(id)
	return models.GetWorkerDetailsDto{
		ID:             worker.ID,
		Name:           worker.Name,
		Email:          worker.Email,
		PhoneNumber:    worker.PhoneNumber,
		Address:        worker.Address,
		Degree:         worker.Degree,
		Position:       worker.Position,
		JobDescription: worker.JobDescription,
		Department:     worker.Department,
		Salary:         worker.Salary,
		BirthData:      worker.BirthData,
		HiredAt:        worker.HiredAt,
		FaceImg:        worker.FaceImg,
		IdImg:          worker.IdImg,
		Qualifications: worker.Qualifications,
	}, err
}
