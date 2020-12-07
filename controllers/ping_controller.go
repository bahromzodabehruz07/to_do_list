package controllers

import (
	"fmt"
	"github.com/to_do_list/response"
	"net/http"
	"time"
)

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	str := currentTime.Format("01-02-2006")
	fmt.Println(str)
	response.JSON(w, http.StatusOK, "Hello from s")
}
