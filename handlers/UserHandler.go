package handlers

import (
	"net/http"
	"restapi/handlers/param"
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
	router.Get("/", h.getAllUser)
	router.Get("/{id}", h.getUserByID)
}

func (h *UserHandler) getAllUser(w http.ResponseWriter, r *http.Request) {
	d, e := h.userService.GetAll()
	if e != nil {
		NotFound(w, r)
		return
	}
	Ok(w, d)
}

func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	d, e := h.userService.GetUserByID(id)
	if e != nil {
		NotFound(w, r)
		return
	}
	Ok(w, d)
}
