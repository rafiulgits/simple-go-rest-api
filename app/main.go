package main

import (
	"net/http"
	"restapi/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	router := chi.NewRouter()

	//middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	//RESTy routes for "articles"
	router.Route("/articles", handler.ArticleHandler)

	http.ListenAndServe(":8080", router)
}
