package models

import (
	"database/sql"
	"fmt"
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

	//password := []byte(u.Password)
	//// Hashing the password with the default cost of 10
	//
	//hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	//if err != nil {
	//	//panic(err)
	//}

	insertStatement := "INSERT INTO users(login,password,fullname) VALUES ($1,$2,$3)"
	_, err := db.Exec(insertStatement, u.Login, u.Password, u.FullName)
	if err != nil {
		return &Response{Status: http.StatusBadRequest, Message: "error when creating new query", Payload: nil}
	}

	return nil

}
func (u *User) SignIn(db *sql.DB) *Response {
	//password := []byte(u.Password)
	//// Hashing the password with the default cost of 10
	//
	////hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	////if err != nil {
	////	//panic(err)
	////	fmt.Println(hashedPassword)
	////}
	////user := User{}
	//fmt.Println(u)
	//fmt.Println(hashedPassword)
	findUserQuery := "SELECT *FROM users WHERE users.login=$1 and users.password=$2"
	db.QueryRow(findUserQuery, u.Login, u.Password).Scan(&u.Id, &u.FullName, &u.Password, &u.FullName)
	token, err := auth.CreateToken(u.Id)
	if err != nil {

	}
	u.Token = token
	fmt.Println(token)
	fmt.Println(u)
	return nil

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
