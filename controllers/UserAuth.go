package controllers

import (
	"encoding/json"
	"github.com/luqmansen/Coolinary/models"
	u "github.com/luqmansen/Coolinary/utils"
	"net/http"
)

var CreateUserAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.User{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, u.InvalidReq))
		return
	}

	resp := account.CreateUser()
	u.Respond(w, resp)

}

var AuthenticateUser = func(w http.ResponseWriter, r *http.Request) {

	account := &models.User{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, u.InvalidReq))
	}

	resp := models.LoginUser(account.Email, account.Password)
	u.Respond(w, resp)
}
