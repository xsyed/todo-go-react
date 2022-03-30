package routes

import (
	"net/http"
	"todolist/controllers"

	"github.com/go-chi/chi"
)

func TodoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", controllers.FetchTodos)
		r.Post("/", controllers.CreateTodo)
		r.Put("/{id}", controllers.UpdateTodo)
		r.Delete("/{id}", controllers.DeleteTodo)
	})
	return rg
}
