package controllers

import (
	"encoding/json"
	"github.com/luqmansen/Coolinary/models"
	u "github.com/luqmansen/Coolinary/utils"
	"net/http"
)

var CreateSellerAccount = func(w http.ResponseWriter, r *http.Request) {

	seller := &models.Seller{}
	err := json.NewDecoder(r.Body).Decode(seller)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, u.InvalidReq))
		return
	}

	resp := seller.CreateStore()
	u.Respond(w, resp)

}

var AuthenticateSeller = func(w http.ResponseWriter, r *http.Request) {

	seller := &models.Seller{}
	err := json.NewDecoder(r.Body).Decode(seller)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, u.InvalidReq))
	}

	resp := models.LoginSeller(seller.Email, seller.Password)
	u.Respond(w, resp)
}
