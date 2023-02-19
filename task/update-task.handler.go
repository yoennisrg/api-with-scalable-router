package task

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := findTaskByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var updatedTask Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if(updatedTask.Title != ""){
		task.Title = updatedTask.Title
	}

	if(updatedTask.Description != ""){
		task.Description = updatedTask.Description
	}

	if(updatedTask.Completed){
		task.Completed = updatedTask.Completed
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
