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
	s.Router.HandleFunc("/api/createTodo", middelwares.SetMiddlewareAuthentication(s.CreateToDO)).Methods("POST")
	s.Router.HandleFunc("/api/updateTodo/{id}/{status}", middelwares.SetMiddlewareAuthentication(s.UpdateToDo)).Methods("PUT")
	s.Router.HandleFunc("/api/deleteTodo/{id}", middelwares.SetMiddlewareAuthentication(s.DeleteToDo)).Methods("DELETE")
	s.Router.HandleFunc("/api/getAllTodo", middelwares.SetMiddlewareAuthentication(s.GetAllToDo)).Methods("GET")

}
