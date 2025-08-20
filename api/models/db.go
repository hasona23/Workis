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
		BirthDate:      time.Date(1990, 5, 12, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_mikejack.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_mikejack.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator"},
			{CertName: "Safety Training"},
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
		BirthDate:      time.Date(1988, 8, 22, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2018, 7, 15, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_annasmith.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_annasmith.jpg"},
		Qualifications: []Qualification{
			{CertName: "First Aid"},
			{CertName: "Safety Training"},
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
		BirthDate:      time.Date(1992, 11, 5, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_johndoe.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_johndoe.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator"},
			{CertName: "Hazmat Handling"},
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
		BirthDate:      time.Date(1988, 8, 22, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2018, 7, 15, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_annasmith.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_annasmith.jpg"},
		Qualifications: []Qualification{
			{CertName: "First Aid"},
			{CertName: "Safety Training"},
		},
	})
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
		BirthDate:      time.Date(1990, 5, 12, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_mikejack.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_mikejack.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator"},
			{CertName: "Safety Training"},
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
		BirthDate:      time.Date(1992, 11, 5, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_johndoe.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_johndoe.jpg"},
		Qualifications: []Qualification{
			{CertName: "Forklift Operator"},
			{CertName: "Hazmat Handling"},
		},
	})
	Db.Create(&Worker{
		Name:           "Lisa Ray",
		Email:          "lisa.ray@example.com",
		PhoneNumber:    "555-0104",
		Address:        "321 Maple St, Springfield",
		Degree:         "MBA Management",
		Position:       "Project Manager",
		JobDescription: "Manages projects and coordinates teams.",
		Department:     "Management",
		Salary:         55000,
		BirthDate:      time.Date(1985, 2, 18, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2015, 9, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_lisaray.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_lisaray.jpg"},
		Qualifications: []Qualification{
			{CertName: "Project Management"},
			{CertName: "Safety Training"},
		},
	})
	Db.Create(&Worker{
		Name:           "Tom Lee",
		Email:          "tom.lee@example.com",
		PhoneNumber:    "555-0105",
		Address:        "654 Cedar Ave, Springfield",
		Degree:         "Diploma Electrical",
		Position:       "Electrician",
		JobDescription: "Maintains electrical systems and equipment.",
		Department:     "Maintenance",
		Salary:         37000,
		BirthDate:      time.Date(1987, 6, 30, 0, 0, 0, 0, time.UTC),
		HiredAt:        time.Date(2017, 4, 20, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
		FaceImg:        &Image{Type: JPEG, Path: "face_tomlee.jpg"},
		IdImg:          &Image{Type: PNG, Path: "id_tomlee.jpg"},
		Qualifications: []Qualification{
			{CertName: "Electrical Certification"},
			{CertName: "First Aid"},
		},
	})
}
