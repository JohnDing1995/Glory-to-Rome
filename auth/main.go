package main

import (
	"fmt"
	"glory-to-rome/auth/model"
	"glory-to-rome/auth/router"
	"glory-to-rome/auth/utils"
	"os"
)

func main() {
	var appPort = os.Getenv("APP_PORT")
	db, _ := utils.ConnectPostgreSQL()
	// remove in production?
	db.AutoMigrate(&model.User{})
	if appPort == "" {
		appPort = "8080"
	}
	fmt.Println("Start auth server at port: ", appPort)
	router.SetupServer(appPort)
}
