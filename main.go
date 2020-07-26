package main

import (
	"encoding/json"
	"net/http"
	"restapi/app"
	"restapi/config"
	"restapi/conn"
	"restapi/di"

	"github.com/go-chi/chi"
	"go.uber.org/dig"
)

//BuildContainer inject required dependencies
func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewDBConfig)
	container.Provide(conn.ConnectDB)
	container.Provide(di.NewArticleRepository)
	container.Provide(di.NewArticleService)
	container.Provide(NewServer)

	return container
}

//Server :
type Server struct {
	dbConfig       *config.DBConfig
	articleService *di.ArticleService
}

//Handler :
func (s *Server) Handler() http.Handler {
	router := chi.NewRouter()

	router.HandleFunc("/article", s.articleHandler)

	return router
}

func (s *Server) articleHandler(w http.ResponseWriter, r *http.Request) {
	x := s.articleService.GetArticle(2)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)

}

//NewServer :
func NewServer(config *config.DBConfig, service *di.ArticleService) *Server {

	return &Server{
		dbConfig:       config,
		articleService: service,
	}
}

//Run :
func (s *Server) Run() {
	http.ListenAndServe(":8080", s.Handler())
}
func main() {

	// router := chi.NewRouter()

	// //middlewares
	// router.Use(middleware.Logger)
	// router.Use(middleware.RealIP)

	// //RESTy routes for "articles"
	// router.Route("/articles", handler.ArticleHandler)

	// containier := BuildContainer()
	// d := Dependency{}
	// d.Inject()
	// err := containier.Invoke(func(server *Server) {
	// 	server.Run()
	// })

	// if err != nil {
	// 	panic(err)
	// }
	// http.ListenAndServe(":8080", router)
	app.NewApp()
}
