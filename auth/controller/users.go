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

func MainHandler(w http.ResponseWriter, r *http.Request) {
	res := utils.GenerateResponse(true, "This is the main page, välkommen!", "")
	utils.Respond(w, res)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var res map[string]interface{}
	vars := mux.Vars(r)
	user, err := getUserByID(vars["id"])
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

func InsertOrUpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user model.User
	var res map[string]interface{}

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
		res = utils.GenerateResponse(false, err.Error(), "")
		utils.Respond(w, res)
	}
	isAdd, err := addOrUpdateUser(user)
	// format a response object
	if err != nil {
		log.Fatalf("Unable to insert to the database. %v", err)
		res = utils.GenerateResponse(false, err.Error(), "")
	} else {
		if isAdd {
			// Added a new object
			log.Println("Added new user")
			res = utils.GenerateResponse(true, "New user added", "")
		}
		res = utils.GenerateResponse(true, "user updated", user)
	}
	utils.Respond(w, res)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := deleteUserByID(vars["id"])
	if err != nil {
		res := utils.GenerateResponse(false, fmt.Sprintf(err.Error()), "")
		utils.Respond(w, res)
	} else {
		res := utils.GenerateResponse(true, fmt.Sprintf("User with id %s deleted", vars["id"]), "")
		utils.Respond(w, res)
	}

}
