package repositories

import (
	"github.com/hasona23/workis/api/helpers"
	"github.com/hasona23/workis/api/models"
)

func UpdateQualification(newQualification models.Qualification) (err error) {
	err = models.Db.Updates(&models.Qualification{
		ID:       newQualification.ID,
		CertName: newQualification.CertName,
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

func QualificationExistID(id int) bool {
	var q models.Qualification
	models.Db.First(&q, id)
	return q.CertName != "" && q.ID != 0
}
