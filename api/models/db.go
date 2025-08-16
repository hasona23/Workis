package models

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

const DB_PATH = "./app.db"

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", DB_PATH)
	if err != nil {
		panic(fmt.Sprintln("error init db: ", err))
	}
}
func GetSqlScript(scriptPath string) (script string) {
	//scriptPath, _ = filepath.Abs(scriptPath)
	resolvedPath := filepath.Clean(scriptPath)
	fileData, err := os.ReadFile(resolvedPath)
	if err != nil {
		panic(fmt.Sprintf("error reading sql script %v %v", scriptPath, err))
	}
	script = string(fileData)
	return script
}
func readCreateTableFile(fileName string) (createTableCmd string) {
	fileData, err := os.ReadFile(fmt.Sprintf(".\\sql\\schema\\%v", fileName))
	if err != nil {
		panic(fmt.Sprintf("error reading table sql %v , %v", fileName, err))
	}
	createTableCmd = string(fileData)
	return createTableCmd
}

func createTable(cmd string) {
	_, err := DB.Exec(cmd)
	if err != nil {
		panic(fmt.Sprintln("error creating table: ", err))
	}
}
func CreateWorkerTable() {

	createCmd := readCreateTableFile("worker_table.sql")
	createTable(createCmd)

	createCmd = readCreateTableFile("qualifications_table.sql")
	createTable(createCmd)
}

func SeedWorkers() {

	addWorkersSQL :=
		`INSERT INTO Workers (Name, DateCreated, IsActive, Position, Salary, BirthDate,Gender,Degree,JobDescription) VALUES (?, ?, ?, ?, ?, ?,?,?,?)`

	getRandomDay := func() int { return rand.IntN(29) }
	getRandomMonth := func() int { return rand.IntN(13) }
	getRandomYear := func(yearsInterval int) int { return rand.IntN(yearsInterval) + 2000 }
	getRandomDate := func(yearsInterval int) time.Time {
		return time.Date(getRandomYear(yearsInterval), time.Month(getRandomMonth()), getRandomDay(), rand.IntN(25), 0, 0, 0, time.UTC)
	}
	getRandomSalary := func() float32 { return float32(rand.IntN(5) * 10000) }
	DB.Exec(addWorkersSQL, "Alice Morgan", getRandomDate(26), false, "Developer", getRandomSalary(), getRandomDate(1), FEMALE, "CS HARVARD BACHELORS", "Tech lead")
	DB.Exec(addWorkersSQL, "Mohsen Togomori", getRandomDate(26), false, "Developer", getRandomSalary(), getRandomDate(1), MALE, "CS OXFORD MASTERS", "Testing")
	DB.Exec(addWorkersSQL, "Mike Jhoncena", getRandomDate(26), true, "Sales", getRandomSalary(), getRandomDate(1), MALE, "NONE", "Marketting for product")
	DB.Exec(addWorkersSQL, "Dodi bn mickey", getRandomDate(26), true, "CEO", getRandomSalary(), getRandomDate(1), MALE, "Business OXFORD BACHELORS", "MANAGE BUSINESS")
	DB.Exec(addWorkersSQL, "Alex  ...", getRandomDate(26), true, "SECRETARY", getRandomSalary(), getRandomDate(1), FEMALE, "NONE", "Calls + paper work")

	addQualificationSQL := `INSERT INTO Qualifications (WorkerId,CertName,IsActive) VALUES (?,?,?)`
	for i := 0; i < 20; i++ {
		cert_name := "Cert1"
		if i%2 == 0 {
			cert_name = "Cert2"
		}
		DB.Exec(addQualificationSQL, rand.Int32N(5), cert_name, !(i%5 == 0))
	}

}
func DeleteWorkerTable() {
	deleteTableSQL := `DROP TABLE IF EXISTS Workers`
	_, err := DB.Exec(deleteTableSQL)
	if err != nil {
		panic(err)
	}
	deleteTableSQL = `DROP TABLE IF EXISTS Qualifications`
	_, err = DB.Exec(deleteTableSQL)
	if err != nil {
		panic(err)
	}
}
