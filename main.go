package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seba000/RestApi-go-gorm/db"
	"github.com/seba000/RestApi-go-gorm/models"
	"github.com/seba000/RestApi-go-gorm/routes"
)



func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//se asigna el new router de mux a la variable r
	r := mux.NewRouter()


	//se definen las rutas
	r.HandleFunc("/", routes.HomeHandler)



	//se inicia e asigna elp puerto al server
	http.ListenAndServe(":3000", r)
}

//go run .