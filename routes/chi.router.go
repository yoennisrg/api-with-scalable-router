package api

// type ChiRouter struct {
//     *chi.Mux
// }

// func (r *ChiRouter) AddRoute(path string, method string, handler http.Handler) {
//     switch method {
//     case "GET":
//         r.Get(path, handler)
//     case "POST":
//         r.Post(path, handler)
//     case "PUT":
//         r.Put(path, handler)
//     // ... handle other methods
//     }
// }

// func (r *ChiRouter) AddGroup(prefix string, routes []Route) {
//     for _, route := range routes {
//         router.MethodFunc(route.Method, route.Path, route.Func)
//     }
//     r.Mount(prefix, r)
// }

// func (r *ChiRouter) AddMiddleware(middleware http.Handler) {
//     r.Use(middleware)
// }

// func NewChiRouter() *ChiRouter {
//     return &ChiRouter{mux.NewRouter()}
// }