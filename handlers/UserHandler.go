package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/handlers/param"
	"restapi/models"
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
	router.Post("/", h.createUser)

	//Sub-Routes : /article/{id}
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.getUserByID)
		router.Delete("/", h.deleteUser)
	})
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

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	d, e := h.userService.CreateUser(&user)
	if e != nil {
		NotFound(w, r)
		return
	}
	Created(w, d, "")
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	e := h.userService.DeleteUser(id)
	if e != nil {
		NotFound(w, r)
		return
	}
	NoContent(w)
}
