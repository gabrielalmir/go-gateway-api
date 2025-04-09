package server

import (
	"net/http"

	"github.com/gabrielalmir/go-gateway-api/internal/service"
	"github.com/gabrielalmir/go-gateway-api/internal/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/api/v1/accounts", accountHandler.CreateAccount)
	s.router.Get("/api/v1/accounts", accountHandler.GetAccount)
}

func (s *Server) Start() error {
	s.router.Use(middleware.Logger)

	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	s.ConfigureRoutes()

	return s.server.ListenAndServe()
}
