package repositories

import "github.com/hasona23/workis/api/models"

func UpdateQualification(newQualification models.Qualification) {
	models.Db.Updates(&models.Qualification{
		ID:       newQualification.ID,
		CertName: newQualification.CertName,
		CertImg:  newQualification.CertImg,
	})
}
func CreateQualification(qualification models.Qualification) {
	models.Db.Create(qualification)
}
func DeleteQualification(id int) {
	models.Db.Delete(&models.Qualification{ID: id})
}
