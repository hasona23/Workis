package repositories

import (
	"github.com/hasona23/workis/api/helpers"
	"github.com/hasona23/workis/api/models"
)

func UpdateQualification(newQualification models.Qualification) (err error) {
	err = models.Db.Updates(&models.Qualification{
		ID:       newQualification.ID,
		CertName: newQualification.CertName,
		CertImg:  newQualification.CertImg,
	}).Error
	helpers.LogError(err)
	return err
}
func CreateQualification(qualification models.Qualification) (err error) {
	err = models.Db.Create(qualification).Error
	helpers.LogError(err)
	return err

}
func DeleteQualification(id int) (err error) {
	err = models.Db.Delete(&models.Qualification{ID: id}).Error
	helpers.LogError(err)
	return err
}
