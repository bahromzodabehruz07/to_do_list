package controllers

import (
	"github.com/to_do_list/response"
	"net/http"
)

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, "Ping to server")
}
