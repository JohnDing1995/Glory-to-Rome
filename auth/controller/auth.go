package controller

import (
	"encoding/json"
	"fmt"
	"glory-to-rome/auth/model"
	"glory-to-rome/auth/utils"
	"log"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser model.InputUser
	var res map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if !loginUser.ValidateInputUser() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
		res = utils.GenerateResponse(false, err.Error(), "")
		utils.Respond(w, res)
	}
	var foundUser model.User
	err, canLogin := checkIfLoginUser(loginUser, &foundUser)
	if err != nil || !canLogin {
		// cannot log user in
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// if can login
	token, expirationTime, err := utils.CreateJWT(foundUser)
	if err != nil {
		// If there is an error in creating the JWT
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	utils.Respond(w, utils.GenerateResponse(true, "User logged in",
		map[string]interface{}{"name": foundUser.Name, "email": foundUser.Email},
	))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var inputUser model.InputUser
	var res map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	fmt.Println(inputUser)
	if !inputUser.ValidateInputUser() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
		res = utils.GenerateResponse(false, err.Error(), "")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, res)
		return
	}
	// Create user here
	user := model.User{
		Name:  inputUser.Name,
		Email: inputUser.Email,
	}
	user.PasswordHash = utils.HashAndSalt(inputUser.Password)
	if err := addUser(user); err != nil {
		log.Println(err.Error())
		utils.Respond(w, utils.GenerateResponse(false, err.Error(), ""))
		return
	}
	// Generate the JWT token
	token, expirationTime, err := utils.CreateJWT(user)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set the client cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	// There is no conflict, user added
	utils.Respond(w, utils.GenerateResponse(true, "User added",
		map[string]interface{}{"name": inputUser.Name, "email": inputUser.Email},
	))
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}

func RefeshHandler(w http.ResponseWriter, r *http.Request) {

}
