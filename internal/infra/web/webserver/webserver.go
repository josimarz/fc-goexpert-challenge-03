package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router chi.Router
	port   string
}

func NewWebServer(port string) *WebServer {
	server := &WebServer{
		Router: chi.NewRouter(),
		port:   port,
	}
	server.Router.Use(middleware.Logger)
	return server
}

func (s *WebServer) Start() {
	http.ListenAndServe(s.port, s.Router)
}
