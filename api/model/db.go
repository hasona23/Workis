package model

import (
	"fmt"

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
	Db.AutoMigrate(&Worker{})
	Db.AutoMigrate(&Qualification{})
	fmt.Println("Finished Init DB")
}

func SeedData() {
	fmt.Println("Seeding Data")
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
			{CertName: "Forklift Operator"},
			{CertName: "Safety Training"},
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
			{CertName: "First Aid"},
			{CertName: "Safety Training"},
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
			{CertName: "Forklift Operator"},
			{CertName: "Hazmat Handling"},
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
			{CertName: "Project Management"},
			{CertName: "Safety Training"},
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
			{CertName: "Electrical Certification"},
			{CertName: "First Aid"},
		},
	})

}
