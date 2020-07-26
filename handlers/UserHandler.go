package handlers

import (
	"net/http"
	"restapi/services"

	"github.com/go-chi/chi"
)

//UserHandler :
type UserHandler struct {
	userService *services.IUserService
}

//NewUserHandler :
func NewUserHandler(userService *services.IUserService) IHandler {
	return &UserHandler{
		userService: userService,
	}
}

//Handle :
func (h *UserHandler) Handle(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	})
}
