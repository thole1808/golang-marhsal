package main

import (
	"golang-web-service-api/controllers/authcontroller"
	"golang-web-service-api/controllers/kontrakvacontroller"
	"golang-web-service-api/middlewares"
	"golang-web-service-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// koneksi models ke database
	models.ConnectDatabase()

	// setting routes dengan gorilla mux
	r := mux.NewRouter()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	//sub route & middlewares
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/kontrakva", kontrakvacontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8181", r))

}
