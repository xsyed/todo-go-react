package routes

import (
	"net/http"
	"todolist/controllers"

	"github.com/go-chi/chi"
)

func UserHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/", controllers.CreateUser)
	})
	return rg
}
