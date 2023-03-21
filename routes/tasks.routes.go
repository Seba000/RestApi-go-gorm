package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seba000/RestApi-go-gorm/db"
	"github.com/seba000/RestApi-go-gorm/models"
)

//GET TASKS
func GetTasksHandler(w http.ResponseWriter, r *http.Request){
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

//POST TASK
func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(task)
}

//GET TASK
func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task not Found :c"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}


//DELETE TASK
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request){
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task not Found :c"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}


