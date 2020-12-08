package controllers

import (
	"encoding/json"
	"github.com/to_do_list/models"
	"github.com/to_do_list/response"
	"io/ioutil"
	"net/http"
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
	e.Message = "Successfully created"
	e.Status = http.StatusOK
	e.Payload = user
	response.JSON(w, http.StatusOK, e)
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
		return
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
		return
	}
	res := user.SignIn(s.DB)
	response.JSON(w, http.StatusOK, res)
	return

}
