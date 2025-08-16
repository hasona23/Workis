package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/validator"
)

const QUALIFICATIONS_TABLE = "qualifications"

func CreateQualification(workerId int, cert string) error {
	err := validator.ValidateQualification(models.Qualification{WorkerId: workerId, CertName: cert})
	if err != nil {
		return err
	}

	if !WorkerIdExists(workerId) {
		return fmt.Errorf("couldnt find worker with id %v", workerId)
	}

	insertSQL := models.GetSqlScript("./sql/queries/qualifications/create.sql")
	_, err = models.DB.Exec(insertSQL, workerId, cert, true)
	if err != nil {
		return err
	}
	return nil
}

func GetAllQualifications() (qualifications []models.Qualification, err error) {
	querySQL := models.GetSqlScript("./sql/queries/qualifications/gettAll.sql")
	rows, err := models.DB.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		qualifications = append(qualifications, GetQualificationFromRow(rows))
	}
	return qualifications, nil
}

func GetWorkerQualifications(workerId int) (qualifications []models.Qualification, err error) {

	querySQL := models.GetSqlScript("./sql/queries/qualifications/getWorkerId.sql")
	rows, err := models.DB.Query(querySQL, workerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		qualifications = append(qualifications, GetQualificationFromRow(rows))
	}
	return qualifications, nil
}
func DeleteQualification(id int) error {
	delSQL := models.GetSqlScript("./sql/queries/qualifications/delete.sql")
	_, err := models.DB.Exec(delSQL, id)
	if err != nil {
		return err
	}
	return nil
}
func GetQualificationByID(id int) (qualification *models.Qualification, err error) {

	querySQL := models.GetSqlScript("./sql/queries/qualifications/getId.sql")
	rows, err := models.DB.Query(querySQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		qualificationBuf := GetQualificationFromRow(rows)
		qualification = &qualificationBuf
	}
	if qualification == nil || qualification.ID != id {
		return nil, fmt.Errorf("couldn't find qualification with id %v", id)
	}
	return qualification, nil
}
func UpdateQualification(id int, qualification models.Qualification) error {
	if !QualificationIdExist(id) {
		return fmt.Errorf("not qualification found with id %v", id)
	}
	updateSQL := models.GetSqlScript("./sql/queries/qualifications/update.sql")
	if err := validator.ValidateQualification(qualification); err != nil {
		return err
	}
	_, err := models.DB.Exec(updateSQL,
		qualification.WorkerId, qualification.CertName, qualification.IsActive, qualification.ID)
	if err != nil {
		return err
	}
	return nil
}
func QualificationIdExist(id int) bool {

	qualification, err := GetQualificationByID(id)
	return qualification != nil && err == nil
}
func GetQualificationFromRow(rows *sql.Rows) (qualification models.Qualification) {
	var (
		id       int
		workerId int
		cert     string
		isActive bool
	)

	// Scan in the CORRECT order matching your table schema
	err := rows.Scan(&id, &workerId, &cert, &isActive)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return models.Qualification{} // Return empty worker on error
	}

	qualification = models.Qualification{
		ID:       id,
		WorkerId: workerId,
		CertName: cert,
		IsActive: isActive,
	}
	return qualification
}
