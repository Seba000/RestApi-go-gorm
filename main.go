package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/seba000/RestApi-go-gorm/db"
	"github.com/seba000/RestApi-go-gorm/routes"
)



func main() {
	db.DBConnection()
	//se asigna el new router de mux a la variable r
	r := mux.NewRouter()


	//se definen las rutas
	r.HandleFunc("/", routes.HomeHandler)



	//se inicia e asigna elp puerto al server
	http.ListenAndServe(":3000", r)
}

//go run .