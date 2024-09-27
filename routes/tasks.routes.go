package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emicantarero/rest-API-gorillaMux/db"
	"github.com/emicantarero/rest-API-gorillaMux/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	json.NewDecoder(r.Body).Decode(&tasks)
	createdTask := db.DB.Create(&tasks)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tarea no encontrada"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tarea no encontrada"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Tarea eliminada correctamente"))
}
