package task

import (
	api "api/routes"
)

var Routes = []api.Route{
	{Path: "", Func: getTasks, Method: "GET"},
	{Path: "", Func: createTask, Method: "POST"},
	{Path: "/{id}", Func: getTask, Method: "GET"},
	{Path: "/{id}", Func: updateTask, Method: "PUT"},
	{Path: "/{id}", Func: deleteTask, Method: "DELETE"},
}

var TaskRoutes = api.RouteList{
	Prefix: "/tasks",
	Routes: Routes,
}
