package repositories

import (
	"errors"
	"time"

	"github.com/hasona23/workis/api/models"
	"gorm.io/gorm"
)

func GetAllWorkers() (workers []models.Worker) {
	models.Db.Where("deleted_at IS NULL").Find(&workers)
	return workers
}

func GetAllWorkersWithQualifications() (workers []models.Worker) {
	models.Db.Where("deleted_at IS NULL").Preload("Qualifications").Find(&workers)
	return workers
}

func GetWorkerWithID(id int) (worker models.Worker) {
	models.Db.Where("id = ?", id).Preload("Qualifications").First(&worker)
	return worker
}

func WorkerExistsID(id int) bool {
	var worker models.Worker

	err := models.Db.Where("deleted_at IS NULL").First(&worker, id).Error
	return worker.ID != 0 && !errors.Is(err, gorm.ErrRecordNotFound)
}

func CreateWorker(worker models.Worker) {
	models.Db.Create(&worker)
}

func SoftDeleteWorker(id int) {
	deletedTime := time.Now().UTC()
	models.Db.Updates(&models.Worker{ID: id, DeletedAt: &deletedTime})
}
func DeleteWorker(id int) {
	models.Db.Select("Qualifications").Delete(&models.Worker{ID: id})
}
func UpdateWorker(newWorker models.Worker) {
	models.Db.Updates(&models.Worker{
		ID:             newWorker.ID,
		Name:           newWorker.Name,
		Email:          newWorker.Email,
		Address:        newWorker.Address,
		Position:       newWorker.Position,
		Degree:         newWorker.Degree,
		JobDescription: newWorker.JobDescription,
		Department:     newWorker.Department,
		Salary:         newWorker.Salary,
		FaceImg:        newWorker.FaceImg,
		IdImg:          newWorker.IdImg,
	})
}
func ReviveWorker(id int) {
	models.Db.Model(&models.Worker{}).Where("id = ?", id).Update("deleted_at", nil)
}
