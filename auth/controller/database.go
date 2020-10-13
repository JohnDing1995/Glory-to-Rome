package controller

import (
	"errors"
	"glory-to-rome/auth/model"
	"glory-to-rome/auth/utils"
	"log"
	"net/http"

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

func getUserByID(id interface{}) (*model.User, error) {
	db, err := utils.ConnectPostgreSQL()
	user := model.User{}
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		return &user, err
	}
	err = db.Where("ID = ?", id).First(&user).Error
	return &user, err
}

func deleteUserByID(id interface{}) (err error) {
	db, err := utils.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		return err
	}
	if (db.Delete(&model.User{}, id).RowsAffected == 0) {
		return gorm.ErrRecordNotFound
	}
	return err
}

func addOrUpdateUser(user model.User) (isAdd bool, err error) {
	oldUser := user
	db, err := utils.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		return false, err
	}
	if err = db.Where("ID = ?", user.ID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not find, add new user")
			err = db.Create(&user).Error // create new record if no result
			isAdd = true
		}
	}
	// Record exist, update
	user.Name = oldUser.Name
	user.Email = oldUser.Email
	db.Save(&user)
	return isAdd, err
}
