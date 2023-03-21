package routes

import (
	"encoding/json"
	"net/http"

	"github.com/seba000/RestApi-go-gorm/db"
	"github.com/seba000/RestApi-go-gorm/models"
)

//GET USERS
func GetUsersHandler(w http.ResponseWriter, r *http.Request){
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

//GET 1 user
func GetUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("get user"))
}

//POST
func PostUsersHandler(w http.ResponseWriter, r *http.Request){
	var user models.User 
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user) 
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

//DELETE
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("delete"))
}