package controller

import (
	"encoding/json"
	"fmt"
	"glory-to-rome/auth/model"
	"glory-to-rome/auth/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func connectDBWithResponse(w http.ResponseWriter) *gorm.DB {
	db, err := utils.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		res := utils.GenerateResponse(false, err.Error(), "")
		utils.Respond(w, res)
	}
	return db
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var res map[string]interface{}

	vars := mux.Vars(r)
	fmt.Println(vars)
	db := connectDBWithResponse(w)
	err := db.Where("id = ?", vars["id"]).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res = utils.GenerateResponse(false, "User not find", "")
			log.Printf("Unable to execute the query or find the user. %v", err)
		} else {
			res = utils.GenerateResponse(false, "Query error", "")
		}
		utils.Respond(w, res)
	} else {
		res = utils.GenerateResponse(true, "Get user infomation", user)
		utils.Respond(w, res)
	}
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	db := connectDBWithResponse(w)
	db.Find(&users)
	res := utils.GenerateResponse(true, "Get user infomation", users)
	utils.Respond(w, res)
}

func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user model.User
	var res map[string]interface{}

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	db, err := utils.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		res = utils.GenerateResponse(false, err.Error(), "")
		utils.Respond(w, res)
	}
	result := db.Create(&user)
	// format a response object
	if result.Error != nil {
		log.Fatalf("Unable to insert to the database. %v", err)
		res = utils.GenerateResponse(false, err.Error(), "")
	} else {
		res = utils.GenerateResponse(true, "New user added", user)
	}
	utils.Respond(w, res)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
