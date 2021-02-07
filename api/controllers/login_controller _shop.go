package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AlexSwiss/prentice/api/auth"
	"github.com/AlexSwiss/prentice/api/models"
	"github.com/AlexSwiss/prentice/api/responses"
	"github.com/AlexSwiss/prentice/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

// Login takes inputs and logs user in
func (server *Server) LoginShop(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	shop := models.Shop{}
	err = json.Unmarshal(body, &shop)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	shop.Prepare()
	err = shop.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(shop.AdminEmail, shop.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

//SignIn validates email and password
func (server *Server) SignInShop(adminemail, password string) (string, error) {

	var err error

	shop := models.Shop{}

	err = server.DB.Debug().Model(models.Shop{}).Where("adminemail = ?", adminemail).Take(&shop).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(shop.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(shop.ID)
}
