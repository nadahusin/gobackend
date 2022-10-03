package users

import (
	"github.com/gorilla/mux"
	"github.com/nadahusin/rental/src/middleware"
	"gorm.io/gorm"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckRole(ctrl.GetAllUser))).Methods("GET")
	route.HandleFunc("/", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/{id}", ctrl.UpdateUser).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.DeleteUser).Methods("DELETE")

}
