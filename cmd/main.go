package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	//Todo, dbconfig
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	//s.Router.Use(middleware.AllowContentType("application/json"))
	//routes.RegisterItemRoutes(s.Router)
	//Mound rest of the handlers here
}

func main() {

}
