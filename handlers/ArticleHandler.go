package handlers

import (
	"net/http"
	"restapi/services"

	"github.com/go-chi/chi"
)

//ArticleHandler :
type ArticleHandler struct {
	articleService *services.IArticleService
}

//NewArticleHandler :
func NewArticleHandler(articleService *services.IArticleService) IHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

//Handle :
func (h *ArticleHandler) Handle(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	})
}
