package handlers

import (
	"net/http"
	"restapi/services"

	"github.com/go-chi/chi"
)

//IUserHandler :
type IUserHandler interface {
	IHandler
}

//UserHandler :
type UserHandler struct {
	userService services.IUserService
}

//NewUserHandler :
func NewUserHandler(userService services.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

//Handle :
func (h *UserHandler) Handle(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		d, e := h.userService.GetAll()
		if e != nil {
			NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		JSONResponse(w, d)
	})
}
