package controllers

import (
	"net/http"

	"github.com/AlexSwiss/prentice/api/responses"
)

// Home returns welcome message
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To The Prentice API")

}
