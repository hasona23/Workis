package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

const DB_PATH = "./app.db"

var Db *gorm.DB

func InitDB() {
	fmt.Println("Init DB")
	var err error
	Db, err = gorm.Open(sqlite.Open("./app.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintln("error init db: ", err))
	}
	Db.AutoMigrate(&Worker{}, Qualification{})

	fmt.Println("Finished Init DB")
}

func SeedData(deletePrev bool) {

	fmt.Println("Seeding Data")
	if deletePrev {
		ctx := context.Background()
		rows, err := gorm.G[Worker](Db).Where("true").Delete(ctx)
		if err != nil {
			fmt.Println("[ERROR] ", err)
		}
		fmt.Println("Rows Deleted", rows)
	}
	Db.Create(&Worker{
		Name:           "Mike Jack",
		Email:          "mike.jack@example.com",
		PhoneNumber:    "555-0101",
		Address:        "123 Main St, Springfield",
		Degree:         "BSc Logistics",
		Position:       "Warehouse Operator",
		JobDescription: "Handles warehouse operations and inventory management.",
		Department:     "Logistics",
		Salary:         42000,
		BirthData:      time.Date(1990, 5, 12, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        Image{Type: JPEG, Path: "face_mikejack.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_mikejack.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator", CertImg: Image{Type: PNG, Path: "cert_forklift_mikejack.png"}},
			{CertName: "Safety Training", CertImg: Image{Type: PNG, Path: "cert_safetytraining_mikejack.png"}},
		},
	})
	Db.Create(&Worker{
		Name:           "Anna Smith",
		Email:          "anna.smith@example.com",
		PhoneNumber:    "555-0102",
		Address:        "456 Oak Ave, Springfield",
		Degree:         "MSc Safety",
		Position:       "Safety Officer",
		JobDescription: "Ensures workplace safety and compliance.",
		Department:     "Health & Safety",
		Salary:         48000,
		BirthData:      time.Date(1988, 8, 22, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2018, 7, 15, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        Image{Type: JPEG, Path: "face_annasmith.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_annasmith.jpg"},
		Qualifications: []Qualification{
			{CertName: "First Aid", CertImg: Image{Type: PNG, Path: "cert_firstaid_annasmith.png"}},
			{CertName: "Safety Training", CertImg: Image{Type: PNG, Path: "cert_safetytraining_annasmith.png"}},
		},
	})
	Db.Create(&Worker{
		Name:           "John Doe",
		Email:          "john.doe@example.com",
		PhoneNumber:    "555-0103",
		Address:        "789 Pine Rd, Springfield",
		Degree:         "Diploma Operations",
		Position:       "Material Handler",
		JobDescription: "Handles materials and assists with logistics.",
		Department:     "Operations",
		Salary:         39000,
		BirthData:      time.Date(1992, 11, 5, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        Image{Type: JPEG, Path: "face_johndoe.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_johndoe.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator", CertImg: Image{Type: PNG, Path: "cert_forklift_johndoe.png"}},
			{CertName: "Hazmat Handling", CertImg: Image{Type: PNG, Path: "cert_hazmat_johndoe.png"}},
		},
	})

	Db.Create(&Worker{
		Name:           "Anna Smith",
		Email:          "anna.smith@example.com",
		PhoneNumber:    "555-0102",
		Address:        "456 Oak Ave, Springfield",
		Position:       "Safety Officer",
		Department:     "Health & Safety",
		JobDescription: "Ensures workplace safety and compliance.",
		FaceImg:        Image{Type: JPEG, Path: "face_annasmith.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_annasmith.jpg"},
		Qualifications: []Qualification{
			{CertName: "First Aid", CertImg: Image{Type: PNG, Path: "cert_firstaid_annasmith.png"}},
			{CertName: "Safety Training", CertImg: Image{Type: PNG, Path: "cert_safetytraining_annasmith.png"}},
		},
	})
	Db.Create(&Worker{
		Name:           "Mike Jack",
		Email:          "mike.jack@example.com",
		PhoneNumber:    "555-0101",
		Address:        "123 Main St, Springfield",
		Position:       "Warehouse Operator",
		Department:     "Logistics",
		JobDescription: "Handles warehouse operations and inventory management.",
		FaceImg:        Image{Type: JPEG, Path: "face_mikejack.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_mikejack.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator", CertImg: Image{Type: PNG, Path: "cert_forklift_mikejack.png"}},
			{CertName: "Safety Training", CertImg: Image{Type: PNG, Path: "cert_safetytraining_mikejack.png"}},
		},
	})
	Db.Create(&Worker{
		Name:           "John Doe",
		Email:          "john.doe@example.com",
		PhoneNumber:    "555-0103",
		Address:        "789 Pine Rd, Springfield",
		Position:       "Material Handler",
		Department:     "Operations",
		JobDescription: "Handles materials and assists with logistics.",
		FaceImg:        Image{Type: JPEG, Path: "face_johndoe.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_johndoe.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator", CertImg: Image{Type: PNG, Path: "cert_forklift_johndoe.png"}},
			{CertName: "Hazmat Handling", CertImg: Image{Type: PNG, Path: "cert_hazmat_johndoe.png"}},
		},
	})
	Db.Create(&Worker{
		Name:           "Lisa Ray",
		Email:          "lisa.ray@example.com",
		PhoneNumber:    "555-0104",
		Address:        "321 Maple St, Springfield",
		Position:       "Project Manager",
		Department:     "Management",
		JobDescription: "Manages projects and coordinates teams.",
		FaceImg:        Image{Type: JPEG, Path: "face_lisaray.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_lisaray.jpg"},
		Qualifications: []Qualification{
			{CertName: "Project Management", CertImg: Image{Type: PNG, Path: "cert_projectmanagement_lisaray.png"}},
			{CertName: "Safety Training", CertImg: Image{Type: PNG, Path: "cert_safetytraining_lisaray.png"}},
		},
	})
	Db.Create(&Worker{
		Name:           "Tom Lee",
		Email:          "tom.lee@example.com",
		PhoneNumber:    "555-0105",
		Address:        "654 Cedar Ave, Springfield",
		Position:       "Electrician",
		Department:     "Maintenance",
		JobDescription: "Maintains electrical systems and equipment.",
		FaceImg:        Image{Type: JPEG, Path: "face_tomlee.jpg"},
		IdImg:          Image{Type: PNG, Path: "id_tomlee.jpg"},
		Qualifications: []Qualification{
			{CertName: "Electrical Certification", CertImg: Image{Type: PNG, Path: "cert_electrical_tomlee.png"}},
			{CertName: "First Aid", CertImg: Image{Type: PNG, Path: "cert_firstaid_tomlee.png"}},
		},
	})
}
