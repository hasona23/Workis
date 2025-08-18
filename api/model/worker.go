package model

type Worker struct {
	ID             int
	Name           string
	Email          string
	PhoneNumber    string
	Address        string
	Position       string
	JobDescription string
	Department     string
	FaceImg        Image `gorm:"embedded;embeddedPrefix:face_"`
	IdImg          Image `gorm:"embedded"`
	Qualifications []Qualification
}
