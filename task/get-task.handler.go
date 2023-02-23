package task

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("getTask")
	
	params := mux.Vars(r)
	task, err := findTaskByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
