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
    AddMiddleware(middleware func (next http.Handler) http.Handler)

	Listen(addr string)
}


type Route struct {
	Path   string
	Func   http.HandlerFunc
	Method string
}

type RouteList struct {
	Prefix  string
	Routes []Route
}

func GetRouter() *MuxRouter {
    return NewMuxRouter()
}