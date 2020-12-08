package models

import (
	"database/sql"
	"github.com/to_do_list/auth"
	"net/http"
)

type User struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Token    string `json:"token"`
}

func (u *User) SaveUser(db *sql.DB) *Response {

	insertStatement := `INSERT INTO users(login,password,fullname) VALUES ($1,$2,$3) RETURNING id`
	row := db.QueryRow(insertStatement, u.Login, u.Password, u.FullName).Scan(&u.Id)
	if row != nil {
		return &Response{Status: http.StatusBadRequest, Message: "error when creating new query", Payload: nil}
	}

	return nil

}

func (u *User) SignIn(db *sql.DB) Response {

	response := Response{}

	findUserQuery := "SELECT *FROM users WHERE users.login=$1 and users.password=$2"
	db.QueryRow(findUserQuery, u.Login, u.Password).Scan(&u.Id, &u.FullName, &u.Password, &u.FullName)
	token, err := auth.CreateToken(u.Id)

	if err != nil {
		response.Message = "Error while inserting"
		response.Status = http.StatusBadRequest
		response.Payload = nil
	}

	u.Token = token
	response.Message = "Successfully created"
	response.Status = http.StatusOK
	response.Payload = u

	return response

}

func (u *User) GetAllUser(db *sql.DB) *Response {

	selectAllStateMen := "Select * from users"

	rows, err := db.Query(selectAllStateMen)
	if err != nil {
		return &Response{Status: http.StatusNotFound, Message: "Error while reading from db", Payload: nil}
	}
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Login, &user.Password, &user.FullName)
		users = append(users, user)
	}

	response := Response{Status: http.StatusOK, Message: "Users was not found", Payload: nil}
	if len(users) > 0 {
		response.Message = "Users was selected"
		response.Payload = users
	}

	return &response

}
