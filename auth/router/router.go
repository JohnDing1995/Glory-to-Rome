package router

import (
	"fmt"
	"glory-to-rome/auth/controller"
	"net/http"

	"github.com/gorilla/mux"
)

var SetupServer = func(appPort string) {
	var router = mux.NewRouter()
	router.HandleFunc("/", controller.MainHandler).Methods("GET")
	router.HandleFunc("/users", controller.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controller.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users", controller.InsertOrUpdateUserHandler).Methods("POST", "PUT")
	router.HandleFunc("/auth/singin", controller.SignInHandler).Methods("POST")
	router.HandleFunc("/auth/signout", controller.SignOutHandler).Methods("POST")
	router.HandleFunc("/auth/register", controller.RegisterHandler).Methods("POST")
	router.HandleFunc("/auth/refresh", controller.RefeshHandler).Methods("POST")
	err := http.ListenAndServe("0.0.0.0:"+appPort, router)
	if err != nil {
		fmt.Print(err)
	}
}
