package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
// For some reason adding parseTime parameter to URL fixed returning null records?? 
const DNS = "root:VDw846#dHPj@tcp(127.0.0.1:3306)/go-api-db?parseTime=true"

type User struct {
	gorm.Model // This line makes this struct as the model for GORM
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string 		`json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Cannot connect to DB")
	}

	// If there is a Schema existing, use it - if not, create it 
	DB.AutoMigrate(&User{})
}

// *** Type "hand" for autocompleted handler function

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
}