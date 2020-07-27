package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/handlers/param"
	"restapi/models"
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
	router.Get("/", h.getAllArticles)
	router.Post("/", h.createArticle)

	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.getArticleByID)
	})
}

func (h *ArticleHandler) getAllArticles(w http.ResponseWriter, r *http.Request) {
	d, e := h.articleService.GetAll()
	if e != nil {
		NotFound(w, r)
		return
	}
	Ok(w, d)
}

func (h *ArticleHandler) getArticleByID(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	d, e := h.articleService.GetArticleByID(id)
	if e != nil {
		NotFound(w, r)
		return
	}
	Ok(w, d)
}

func (h *ArticleHandler) createArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		panic(err)
	}
	d, e := h.articleService.CreateArticle(&article)
	if e != nil {
		NotFound(w, r)
		return
	}
	Created(w, d, "")
}
