package router

import (
	"tugas-6/internal/http/handler"
	"tugas-6/pkg/route"
	"net/http"
)

func PublicRoutes(userHandler handler.UserHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: userHandler.Register,
		},
	}
}

func PrivateRoutes(userHandler handler.UserHandler, todosHandler handler.TodoHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAll,
			Roles:   []string{"admin"},
		},
		{
			Method:  http.MethodPost,
			Path:    "/admin/todos",
			Handler: todosHandler.CreateTodoAsAdmin,
			Roles:   []string{"admin"},
		},
		{
			Method:  http.MethodPost,
			Path:    "/todos",
			Handler: todosHandler.CreateTodoHandler,
			Roles:   []string{"user"},
		},
		{
			Method:  http.MethodGet,
			Path:    "/admin/todos",
			Handler: todosHandler.GetAllHandler,
			Roles:   []string{"admin"},
		},
		{
			Method:  http.MethodGet,
			Path:    "/admin/todos",
			Handler: todosHandler.GetTodosByUserIdAsAdmin,
			Roles:   []string{"admin"},
		},
		{
			Method:  http.MethodGet,
			Path:    "/todos",
			Handler: todosHandler.GetTodosHandler,
			Roles:   []string{"user"},
		},
		{
			Method:  http.MethodPut,
			Path:    "/admin/todos",
			Handler: todosHandler.UpdateTodoAsAdmin,
			Roles:   []string{"admin"},
		},
		{
			Method:  http.MethodPut,
			Path:    "/todos",
			Handler: todosHandler.UpdateTodoHandler,
			Roles:   []string{"user"},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/admin/todos",
			Handler: todosHandler.DeleteTodoAsAdmin,
			Roles:   []string{"admin"},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/todos",
			Handler: todosHandler.DeleteTodoHandler,
			Roles:   []string{"user"},
		},
	}
}
