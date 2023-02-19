package task

import (
	"net/http"

	"github.com/gorilla/mux"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
