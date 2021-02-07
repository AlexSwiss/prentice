package middlewares

import (
	"net/http"

	"github.com/AlexSwiss/prentice/api/auth"
	"github.com/AlexSwiss/prentice/api/response"
)

// SetMiddlewareJSON returns all responses in json format
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication checks and authenticates the token
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			response.Write(w, response.Response{
				Code:    http.StatusInternalServerError,
				Action:  "user",
				Message: "An internal error has occurred.",
			})
			return
		}
		next(w, r)
	}
}
