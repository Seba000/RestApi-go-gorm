package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seba000/RestApi-go-gorm/db"
	"github.com/seba000/RestApi-go-gorm/models"
	"github.com/seba000/RestApi-go-gorm/routes"
)



func main() {

	//base de datos
	db.DBConnection()

	//modelos de task y user
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//se asigna el new router de mux a la variable r
	r := mux.NewRouter()


	//se definen las rutas
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")




	//se inicia e asigna elp puerto al server
	http.ListenAndServe(":3000", r)
}

//go run .