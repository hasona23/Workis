package services

import (
	"fmt"

	"github.com/hasona23/workis/api/helpers"
	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/repositories"
)

func CreateQualification(q models.QualificationCreateRequest) error {

	err := helpers.IsStringInBounds(q.CertName, "Certificate Name", 4, 32)
	if err != nil {
		return err
	}
	workerExist := repositories.WorkerExistsID(q.WorkerID)
	if !workerExist {
		return fmt.Errorf("worker with id %v doesnt exist", q.WorkerID)
	}
	err = repositories.CreateQualification(models.Qualification{
		CertName: q.CertName,
		WorkerID: q.WorkerID,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteQualification(id int) error {
	err := repositories.DeleteQualification(id)
	return err
}
func UpdateQualification(q models.Qualification) error {
	err := helpers.IsStringInBounds(q.CertName, "Certificate Name", 4, 32)
	if err != nil {
		return err
	}

	err = repositories.UpdateQualification(models.Qualification{ID: q.ID, CertName: q.CertName})
	return err
}
