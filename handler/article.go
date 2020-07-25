package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"restapi/models"
	"strconv"

	"github.com/go-chi/chi"
)

// ArticleHandler root
func ArticleHandler(router chi.Router) {
	router.Get("/", getAllArticles)
	router.Post("/", createArticle)

	//Sub-Routers for :id
	router.Route("/{id}", func(router chi.Router) {
		router.Use(articleContext)
		router.Get("/", getArticleByID)
		router.Put("/", updateArticleByID)
		router.Delete("/", deleteArticleByID)
	})
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message" : "Created"}`))
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message" : "OK"}`))
}

func getArticleByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value(key("article")).(models.Article)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not Found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func deleteArticleByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func updateArticleByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value(key("article")).(models.Article)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

type key string

//ArticleContext manage article
func articleContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_id := chi.URLParam(r, "id")
		id, err := strconv.Atoi(_id)
		if err != nil {

		}
		article := models.Article{Title: "Hello", Body: "Hello World", ID: id, AuthorID: 1}
		ctx := context.WithValue(r.Context(), key("article"), article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
