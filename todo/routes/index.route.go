package routes

import (
	"todolist/controllers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func AllRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", controllers.FetchTodos)

	r.Mount("/todo", TodoHandlers())
	r.Mount("/user", UserHandlers())
	r.Post("/login", controllers.LoginUser)
	r.Get("/logout", controllers.Logout)

	return r
}
