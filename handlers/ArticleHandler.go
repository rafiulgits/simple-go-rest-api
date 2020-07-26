package handlers

import (
	"net/http"
	"restapi/services"

	"github.com/go-chi/chi"
)

//IArticleHandler :
type IArticleHandler interface {
	IHandler
}

//ArticleHandler :
type ArticleHandler struct {
	articleService services.IArticleService
}

//NewArticleHandler :
func NewArticleHandler(articleService services.IArticleService) IArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

//Handle :
func (h *ArticleHandler) Handle(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		d, e := h.articleService.GetAll()
		if e != nil {
			NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		JSONResponse(w, d)
	})
}
