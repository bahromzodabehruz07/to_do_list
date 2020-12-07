package controllers

import (
	"github.com/to_do_list/middelwares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/api/ping", middelwares.SetMiddlewareJSON(s.Ping)).Methods("GET")

	//User Route
	s.Router.HandleFunc("/api/user/createUser", middelwares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/user/getAllUsers", middelwares.SetMiddlewareJSON(s.getAllUser)).Methods("GET")
	s.Router.HandleFunc("/api/signIn", middelwares.SetMiddlewareJSON(s.SignIn)).Methods("POST")
	//Todos Route
	s.Router.HandleFunc("/api/createTodo", middelwares.SetMiddlewareJSON(s.CreateToDO)).Methods("POST")

}
