package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "test:password@tcp(127.0.0.1:3306)/crudrest?charset=utf8"

type User struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
}

func InitializeMigration()  {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func CreateUser(w http.ResponseWriter, r * http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}