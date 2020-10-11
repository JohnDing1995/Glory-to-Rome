package router

import (
	"fmt"
	"glory-to-rome/auth/controller"
	"net/http"

	"github.com/gorilla/mux"
)

var SetupServer = func(appPort string) {
	var router = mux.NewRouter()
	router.HandleFunc("/users", controller.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", controller.InsertUserHandler).Methods("POST")

	err := http.ListenAndServe(":"+appPort, router)
	if err != nil {
		fmt.Print(err)
	}
}
