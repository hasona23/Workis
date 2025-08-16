package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hasona23/workis/api/models"
	"github.com/hasona23/workis/api/validator"
)

const WORKERS_TABLE = "workers"
const TIME_LAYOUT = "2006-01-02 15:04:05 +0000 UTC"

func CreateWorker(worker models.Worker) error {

	createSQL := models.GetSqlScript("./sql/queries/workers/create.sql")
	worker.IsActive = true
	worker.DateCreated = time.Now().UTC()
	if err := validator.ValidateWorker(worker); err != nil {
		return err
	}
	_, err := models.DB.Exec(createSQL,
		worker.Name, worker.Position, worker.Salary,
		worker.BirthDate, worker.DateCreated, worker.IsActive,
		worker.Gender, worker.DateCreated, worker.JobDescription)
	if err != nil {
		return err
	}
	return nil
}

func GetAllWorkers() (workers []models.Worker, err error) {

	querySQL := models.GetSqlScript("./sql/queries/workers/getAll.sql")
	rows, err := models.DB.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		workers = append(workers, getWorkerFromRow(rows))
	}
	return workers, nil
}

func UpdateWorker(workerId int, newWorker models.Worker) error {
	if !WorkerIdExists(workerId) {
		return fmt.Errorf("not worker found with id %v", workerId)
	}
	updateSQL := models.GetSqlScript("./sql/queries/workers/update.sql")
	if err := validator.ValidateWorker(newWorker); err != nil {
		return err
	}
	_, err := models.DB.Exec(updateSQL,
		newWorker.Name, newWorker.DateCreated,
		newWorker.IsActive,
		newWorker.Position,
		newWorker.Salary,
		newWorker.BirthDate,
		newWorker.Gender,
		newWorker.Degree,
		newWorker.JobDescription,
		workerId)
	if err != nil {
		return err
	}
	return nil
}

func GetWorkerByID(id int) (worker *models.Worker, err error) {

	querySQL := models.GetSqlScript("./sql/queries/workers/getId.sql")
	rows, err := models.DB.Query(querySQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		workerBuf := getWorkerFromRow(rows)
		worker = &workerBuf
	}
	if worker == nil || worker.ID != id {
		return nil, fmt.Errorf("couldn't find worker with id %v", id)
	}
	return worker, nil
}

func DeleteWorker(id int) error {
	delSQL := models.GetSqlScript("./sql/queries/workers/delete.sql")
	_, err := models.DB.Exec(delSQL, id)
	if err != nil {
		return err
	}
	return nil
}

func WorkerIdExists(id int) bool {
	worker, err := GetWorkerByID(id)
	return worker != nil && err == nil
}

func getWorkerFromRow(rows *sql.Rows) (worker models.Worker) {
	var (
		id   int
		name string

		isActive       bool
		position       string
		salary         float32
		dateCreated    time.Time
		birthDate      time.Time
		gender         string
		degree         string
		jobDescription string
	)

	// Scan in the CORRECT order matching your table schema
	err := rows.Scan(&id, &name, &dateCreated, &isActive, &position, &salary, &birthDate, &gender, &degree, &jobDescription)
	if err != nil {
		log.Printf("[ERROR] Error scanning row: %v", err)
		return models.Worker{} // Return empty worker on error
	}

	worker = models.Worker{
		ID:             id,
		Name:           name,
		IsActive:       isActive,
		Position:       position,
		Salary:         salary,
		DateCreated:    dateCreated,
		BirthDate:      birthDate,
		JobDescription: jobDescription,
		Gender:         gender,
		Degree:         degree,
	}
	worker.Qualifications, err = GetWorkerQualifications(worker.ID)
	if err != nil {
		log.Printf("[ERROR] failed to get worker qualifications %v", err)
		return models.Worker{}
	}
	return worker
}
