package repositories

import (
	"errors"
	"time"

	"github.com/hasona23/workis/api/helpers"
	"github.com/hasona23/workis/api/models"
	"gorm.io/gorm"
)

func GetAllWorkers() (workers []models.Worker, err error) {
	err = models.Db.Where("deleted_at IS NULL").Find(&workers).Error
	helpers.LogError(err)
	return workers, err
}

func GetAllWorkersWithQualifications() (workers []models.Worker, err error) {
	err = models.Db.Where("deleted_at IS NULL").Preload("Qualifications").Find(&workers).Error
	helpers.LogError(err)
	return workers, err
}

func GetWorkerWithID(id int) (worker models.Worker, err error) {
	err = models.Db.Where("id = ?", id).Preload("Qualifications").First(&worker).Error
	helpers.LogError(err)
	return worker, err
}

func WorkerExistsID(id int) bool {
	var worker models.Worker

	err := models.Db.Where("deleted_at IS NULL").First(&worker, id).Error
	return worker.ID != 0 && !errors.Is(err, gorm.ErrRecordNotFound)
}

func CreateWorker(worker models.Worker) (err error) {
	err = models.Db.Create(&worker).Error
	helpers.LogError(err)

	return err
}

func SoftDeleteWorker(id int) (err error) {
	deletedTime := time.Now().UTC()
	err = models.Db.Updates(&models.Worker{ID: id, DeletedAt: &deletedTime}).Error
	helpers.LogError(err)
	return err
}
func DeleteWorker(id int) (err error) {
	err = models.Db.Select("Qualifications").Delete(&models.Worker{ID: id}).Error
	helpers.LogError(err)
	return err

}
func UpdateWorker(newWorker models.Worker) (err error) {
	err = models.Db.Where("id = ?", newWorker.ID).Updates(newWorker).Error
	helpers.LogError(err)
	return err
}
func ReviveWorker(id int) (err error) {
	err = models.Db.Model(&models.Worker{}).Where("id = ?", id).Update("deleted_at", nil).Error
	helpers.LogError(err)
	return err
}
