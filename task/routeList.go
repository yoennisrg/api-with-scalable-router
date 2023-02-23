package task

import (
	api "api/routes"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

var Routes = []api.Route{
	{Path: "", Func: getTasks, Method: "GET"},
	{Path: "", Func: createTask, Method: "POST", Middleware: []api.Middleware{validateRequest}},
	{Path: "/{id}", Func: getTask, Method: "GET", Middleware: []api.Middleware{
		firstMiddleware,
	}},
	{Path: "/{id}", Func: updateTask, Method: "PUT"},
	{Path: "/{id}", Func: deleteTask, Method: "DELETE"},
}

var TaskRoutes = api.RouteList{
	Prefix: "/tasks",
	Routes: Routes,
}

func firstMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("firstMiddleware")
		next(w, r)
	}
}

func getValidationErrors(err error) map[string]string {

	parseString := func(message string) string {
		if index := strings.Index(message, "Error:"); index != -1 {
			return strings.TrimSpace(message[index+len("Error:"):])
		}
		return message
	}

	errors := make(map[string]string)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				errors[e.Field()] = parseString(e.Error())
			}
		}
	}
	return errors
}

type KeyTaskContext string;
var keyTaskContext KeyTaskContext ="task"

func validateRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

		defer r.Body.Close()
        r.Body = ioutil.NopCloser(bytes.NewReader(body))


		var task Task
		if err := json.Unmarshal(body, &task); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(err.Error())
            return
        }

		validate := validator.New()
		if err := validate.Struct(task); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			errors := getValidationErrors(err)
			json.NewEncoder(w).Encode(errors)
			return
		}
	
		ctx := context.WithValue(r.Context(), keyTaskContext, task)
        next(w, r.WithContext(ctx))
	}
}
