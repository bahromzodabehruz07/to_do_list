package controllers

import (
	"encoding/json"
	"github.com/to_do_list/auth"
	"github.com/to_do_list/models"
	"github.com/to_do_list/response"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	e := models.Response{}

	if err != nil {
		e = models.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad date",
			Payload: nil,
		}
		response.JSON(w, http.StatusUnprocessableEntity, e)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		e = models.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad date",
			Payload: nil,
		}
		response.JSON(w, http.StatusUnprocessableEntity, e)
		return
	}

	var res = user.SaveUser(s.DB)
	if res != nil {
		response.JSON(w, http.StatusBadRequest, res)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func (s *Server) getAllUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	res := u.GetAllUser(s.DB)

	response.JSON(w, http.StatusOK, res)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	e := models.Response{}

	if err != nil {
		e = models.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad date",
			Payload: nil,
		}
		response.JSON(w, http.StatusBadRequest, e)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		e = models.Response{
			Status:  http.StatusBadRequest,
			Message: "Bad date",
			Payload: nil,
		}
		response.JSON(w, http.StatusBadRequest, e)
	}
	res := user.SignIn(s.DB)
	response.JSON(w, http.StatusOK, res)

}

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
