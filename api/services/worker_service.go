package services

import (
	"mime/multipart"

	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

type FileRequest struct {
	File   multipart.File
	Header multipart.FileHeader
}

func CreateWorker(worker models.WorkerCreateRequest, faceImg FileRequest, idImg FileRequest) (err error) {
	err = worker.ValidateCreateWorkerRequest()
	if err != nil {
		return err
	}
	repositories.CreateWorker(models.Worker{
		Name:           worker.Name,
		Email:          worker.Email,
		PhoneNumber:    worker.PhoneNumber,
		Address:        worker.Address,
		Degree:         worker.Degree,
		Position:       worker.Position,
		JobDescription: worker.JobDescription,
		Department:     worker.Department,
		Salary:         worker.Salary,
		BirthDate:      worker.BirthDate,
		HiredAt:        worker.HiredAt,
	})
	return err
}

func SoftDeleteWorker(id int) (err error) {
	err = repositories.SoftDeleteWorker(id)

	return err
}
func ReviveWorker(id int) (err error) {
	err = repositories.ReviveWorker(id)

	return err
}
func UpdateWorker(worker models.WorkerUpdateRequest) (err error) {

	err = worker.ValidateWorkerUpdateRequest()
	if err != nil {
		return err
	}
	err = repositories.UpdateWorker(models.Worker{
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
	})
	return err
}

func UpdateWokerImg(workerId int, img FileRequest, isFaceImg bool) error {
	return nil
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
		BirthData:      worker.BirthDate,
		HiredAt:        worker.HiredAt,
		FaceImg:        worker.FaceImg,
		IdImg:          worker.IdImg,
		Qualifications: worker.Qualifications,
	}, err
}
