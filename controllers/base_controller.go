package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	dbName     = "toDo"
	dbUserName = "postgres"
	dbPassword = "2404"
)

func (s *Server) Initialize() {

	var err error

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbUserName, dbName, dbPassword)
	s.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", connStr)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", connStr)
	}

	s.Router = mux.NewRouter()

	s.initializeRoutes()
}
func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
