package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
    *mux.Router
}



// func (r *MuxRouter) AddRoute(path string, method string, handler http.HandlerFunc) {
//     r.HandleFunc(path, handler).Methods(method)
// }

func (r *MuxRouter) AddRoute(path string, method string, handler http.HandlerFunc, middleware ...Middleware) {
	r.HandleFunc(path, buildHandler(handler, middleware)).Methods(method)
}

func (r *MuxRouter) AddGroup(prefix string, routes []Route) {
	router := r.PathPrefix(prefix).Subrouter()
	for _, route := range routes {
		handler := buildHandler(route.Func, route.Middleware)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}

// func (r *MuxRouter) AddGroup(prefix string, routes []Route) {
//     router := r.PathPrefix(prefix).Subrouter()
//     for _, route := range routes {

//         router.HandleFunc(route.Path, route.Func).Methods(route.Method)
//     }
// }

func (r *MuxRouter) AddMiddleware(middleware func (next http.Handler) http.Handler) {
    r.Use(middleware)
}

func (r *MuxRouter) Listen(addr int) {
    fmt.Printf("🚀 Server is running at http://localhost:%v\n", addr);
    if err := http.ListenAndServe(fmt.Sprintf(":%v", addr), r); err != nil {
		log.Fatal(err)
	}
}

func NewMuxRouter() *MuxRouter {
    return &MuxRouter{mux.NewRouter()}
}