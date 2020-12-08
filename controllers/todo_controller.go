package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/to_do_list/auth"
	"github.com/to_do_list/models"
	"github.com/to_do_list/response"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body types.Todo true "New Todo"
// @Success 201 {object} types.Todo
// @Failure 400 {object} HTTPError
// @Router /todos [post]

func (s *Server) CreateToDO(w http.ResponseWriter, r *http.Request) {
	userId, _ := auth.ExtractTokenID(r)
	responsePayload := models.Response{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responsePayload.Status = http.StatusNotAcceptable
		responsePayload.Message = err.Error()
		responsePayload.Payload = nil
		response.JSON(w, http.StatusNotAcceptable, responsePayload)
		return
	}

	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)

	if err != nil {
		responsePayload.Status = http.StatusBadRequest
		responsePayload.Message = "bad date"
		responsePayload.Payload = nil
		response.JSON(w, http.StatusNotAcceptable, responsePayload)
		return
	}
	currentTime := time.Now()
	currentDate := currentTime.Format("01-02-2006")
	todo.UserId = userId
	todo.Date = currentDate
	res := todo.CreateToDo(s.DB)
	response.JSON(w, http.StatusNotAcceptable, res)

}

func (s *Server) UpdateToDo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	t, ter := strconv.ParseBool(vars["status"])
	todo := models.Todo{}
	todo.Id = int64(pid)
	todo.Status = t
	res := todo.UpdateStatus(s.DB)
	fmt.Println(res)
	fmt.Println(ter)
	fmt.Println(t)
	fmt.Println(pid)
	fmt.Println(err)
}

func (s *Server) DeleteToDo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	responsePayload := models.Response{}
	// Check if the post id is valid
	tId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responsePayload.Status = http.StatusBadRequest
		responsePayload.Message = "Bad_date"
		responsePayload.Payload = nil
		response.JSON(w, http.StatusBadRequest, responsePayload)
		return
	}
	deleteStateMent := "DELETE from todo WHERE id=$1"
	_, err = s.DB.Exec(deleteStateMent, tId)
	if err != nil {
		responsePayload.Status = http.StatusNotAcceptable
		responsePayload.Message = err.Error()
		responsePayload.Payload = nil
		response.JSON(w, http.StatusNotAcceptable, responsePayload)
		return
	}
	responsePayload.Status = http.StatusOK
	responsePayload.Message = "Deleted successFully"
	responsePayload.Payload = nil
	response.JSON(w, http.StatusOK, responsePayload)
}

func (s *Server) GetAllToDo(w http.ResponseWriter, r *http.Request) {
	userId, _ := auth.ExtractTokenID(r)
	todo := models.Todo{}
	todo.UserId = userId
	res := todo.GetAllTodos(s.DB)
	response.JSON(w, http.StatusOK, res)
}
