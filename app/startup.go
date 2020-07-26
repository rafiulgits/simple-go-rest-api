package app

import (
	"net/http"
	"restapi/config"
	"restapi/conn"
	"restapi/repositories"
	"restapi/services"

	"github.com/go-chi/chi"
	"go.uber.org/dig"
)

// Dependency Injections
func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewDBConfig)
	container.Provide(conn.ConnectDB)

	// User
	container.Provide(repositories.NewUserRepository)
	container.Provide(services.NewUserService)

	//Article
	container.Provide(repositories.NewArticleRepository)
	container.Provide(services.NewArticleService)

	container.Provide(NewServer)
	return container
}

//App :
type App struct {
	_ *Server
}

//NewApp :
func NewApp() {
	container := buildContainer()
	err := container.Invoke(func(server *Server) {
		server.run()
	})
	if err != nil {
		panic(err)
	}

}

//Server :
type Server struct {
	userService    *services.IUserService
	articleService *services.IArticleService
}

//NewServer : constructor of Server
func NewServer(userService services.IUserService, articleService services.IArticleService) *Server {
	return &Server{
		userService:    &userService,
		articleService: &articleService,
	}
}

func (s *Server) handler() http.Handler {
	router := chi.NewRouter()
	return router
}

func (s *Server) run() {
	http.ListenAndServe(":8080", s.handler())
}
