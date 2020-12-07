package middelwares

import (
	"github.com/to_do_list/auth"
	"github.com/to_do_list/models"
	"github.com/to_do_list/response"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, models.Response{Status: http.StatusUnauthorized, Message: "invalid token", Payload: nil})
			return
		}
		next(w, r)
	}
}
