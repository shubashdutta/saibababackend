package router

import (
	"github.com/gorilla/mux"
	"github.com/shubash/saibaba/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/", controller.Home).Methods("GET")
	router.HandleFunc("/api/singup", controller.Singup).Methods("POST")
	router.HandleFunc("/api/login", controller.Login).Methods("POST")
	// router.HandleFunc("/api/users", Controller.GetAllUser).Methods("GET")
	// router.HandleFunc("/api/update/{id}", Controller.UPDATEPASSWORD).Methods("PUT")

	// router.HandleFunc("/api/deleteusers", Controller.DeleteAll).Methods("DELETE")

	return router
}
