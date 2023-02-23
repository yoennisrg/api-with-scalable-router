package main

import (
	api "api/routes"
	"api/task"
	"log"
	"net/http"
)

func main() {
	r := api.GetRouter()
	r.AddGroup("/tasks", task.Routes)
	r.AddMiddleware(loggingMiddleware)
	r.Listen(8080)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
