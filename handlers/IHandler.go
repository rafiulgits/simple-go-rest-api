package handlers

import (
	"github.com/go-chi/chi"
)

//IHandler :
type IHandler interface {
	Handle(router chi.Router)
}
