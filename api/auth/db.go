package auth

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var UserDB *gorm.DB

func CreateUserDB() {
	var err error
	UserDB, err = gorm.Open(sqlite.Open("./user.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintln("failed to make User DB: ", err))
	}
	UserDB.AutoMigrate(&User{})
}
