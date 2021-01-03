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
	newUser := user
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
	user.Name = newUser.Name
	user.Email = newUser.Email
	db.Save(&user)
	return isAdd, err
}

func addUser(user model.User) (err error) {
	db, err := utils.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		return err
	}
	if err = db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("User not find, add new user")
			err = db.Create(&user).Error // create new record if no result
			return nil
		}
		return err
	}
	log.Println("User exists")
	return errors.New("User already exists")
}

func checkIfLoginUser(loginUser model.InputUser, foundUser *model.User) (err error, canLogin bool) {
	db, err := utils.ConnectPostgreSQL()
	if err != nil {
		log.Fatalf("Unable to connect to the database. %v", err)
		return err, false
	}
	var userInDB model.User
	if err = db.Where("Email = ?", loginUser.Email).First(&userInDB).Error; err != nil {
		log.Println("Unable to find user. %v", err)
		return err, false
	}
	foundUser.Email = userInDB.Email
	foundUser.Name = userInDB.Name
	return err, userInDB.PasswordHash == utils.HashAndSalt(loginUser.Password)
}
