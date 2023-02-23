package api

import (
	"net/http"
)

type Router interface {
	// Adds a route to the router
	AddRoute(path string, method string, handler http.Handler)

	// Adds a group of routes to the router
	AddGroup(prefix string, routes []Route)

	// Adds a middleware to the router
	AddMiddleware(middleware func(next http.Handler) http.Handler)

	Listen(addr string)
}

type Route struct {
	Path       string
	Func       http.HandlerFunc
	Method     string
	Middleware []Middleware
}

type RouteList struct {
	Prefix string
	Routes []Route
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func buildHandler(h http.HandlerFunc, middleware []Middleware) http.HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func GetRouter() *MuxRouter {
	return NewMuxRouter()
}
