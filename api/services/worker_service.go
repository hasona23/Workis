package services

import (
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

func CreateWorker(worker models.WorkerCreateRequest) {

}

func DeleteWorker(id int) {
	repositories.DeleteWorker(id)
}

func UpdateWorker(worker models.WorkerUpdateRequest) {

}

func GetAllWorkers(withQualifications bool) (workersGet []models.GetWorkerDto) {

	workers := repositories.GetAllWorkers()

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
	return workersGet
}

func GetWorkerByID(id int) models.GetWorkerDetailsDto {
	worker := repositories.GetWorkerWithID(id)
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
	}
}
