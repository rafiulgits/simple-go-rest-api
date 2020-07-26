package app

import (
	"net/http"
	"restapi/config"
	"restapi/conn"
	"restapi/handlers"
	"restapi/repositories"
	"restapi/services"

	"github.com/go-chi/chi/middleware"

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

	//Handlers
	container.Provide(handlers.NewUserHandler)
	container.Provide(handlers.NewArticleHandler)

	container.Provide(NewServer)
	return container
}

//App :
type App struct {
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
	//handlers
	userHandler    handlers.IUserHandler
	articleHandler handlers.IArticleHandler

	//
	router    *chi.Mux
	dbContext *conn.DB
}

//NewServer : constructor of Server
func NewServer(
	userHandler handlers.IUserHandler,
	articleHandler handlers.IArticleHandler,

	dbContext *conn.DB) *Server {

	return &Server{
		userHandler:    userHandler,
		articleHandler: articleHandler,

		dbContext: dbContext,
		router:    chi.NewRouter(),
	}
}

func (s *Server) setMiddlewares() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.RealIP)

	s.dbContext.Migration()
}

func (s *Server) mapHandlers() {
	s.router.Route("/article", s.articleHandler.Handle)
	s.router.Route("/user", s.userHandler.Handle)
}

func (s *Server) dispose() {
	_ = s.dbContext.Close()
}

func (s *Server) run() {
	s.setMiddlewares()
	s.mapHandlers()
	defer s.dispose()
	http.ListenAndServe(":8080", s.router)

}
