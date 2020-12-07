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

	return response
}
