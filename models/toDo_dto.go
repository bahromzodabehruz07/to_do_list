package models

import (
	"database/sql"
	"net/http"
)

type Todo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Date        string `json:"date"`
	UserId      int64  `json:"user_id"`
}

func (t *Todo) CreateToDo(db *sql.DB) Response {
	response := Response{}
	insertStateMent := "INSERT INTO todo(name,description,status,date,userId) VALUES($1,$2,$3,$4,$5) RETURNING id"
	err := db.QueryRow(insertStateMent, t.Name, t.Description, t.Status, t.Date, t.UserId).Scan(&t.Id)
	if err != nil {
		response.Status = http.StatusForbidden
		response.Message = err.Error()
		response.Payload = nil
		return response
	}
	response.Message = "Created todo"
	response.Status = http.StatusOK
	response.Payload = t

	return response
}

func (t *Todo) GetAllTodos(db *sql.DB) Response {

	response := Response{}

	getAllState := "SELECT * FROM todo"
	row, err := db.Query(getAllState)

	if err != nil {
		response.Status = http.StatusNotAcceptable
		response.Message = "No date"
		response.Payload = nil
	}

	var todos []Todo

	for row.Next() {
		var todo Todo
		row.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Date, &todo.Status, &todo.UserId)
		todos = append(todos, todo)
	}
	response.Status = http.StatusOK
	response.Message = "No date"
	response.Payload = nil

	if len(todos) > 0 {
		response.Message = "Succesfully getted"
		response.Payload = todos
	}
	return response
}

func (t *Todo) UpdateStatus(db *sql.DB) Response {
	response := Response{}

	updateStatus := "Update todo set status=$1 where id=$2 RETURNING id"
	err := db.QueryRow(updateStatus, t.Status, t.Id).Scan(&t.Id)

	if err != nil {
		response.Status = http.StatusForbidden
		response.Message = err.Error()
		response.Payload = nil
		return response
	}
	response.Status = http.StatusCreated
	response.Message = "Updated succesFully"
	response.Payload = t

	return response

}
