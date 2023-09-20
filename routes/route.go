package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/harshau007/go-api/controller"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.MethodFunc("GET", "/users", controller.GetAllUser)
	router.MethodFunc("POST", "/user", controller.CreateUser)
	router.MethodFunc("PUT", "/user/{id}", controller.UpdateUser)
	router.MethodFunc("DELETE","/deleteuser/{id}", controller.DeleteUser)
	router.MethodFunc("DELETE","/deletealluser", controller.DeleteAllUser)

	return router
}