package controllers

import (
	"encoding/json"
	"github.com/luqmansen/Coolinary/models"
	"github.com/luqmansen/Coolinary/utils"
	"net/http"
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {

	seller := r.Context().Value("user").(uint) //Grab the Id of seller creator

	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		utils.Respond(w, utils.Message(http.StatusInternalServerError, "Error while requesting body"))
		return
	}

	product.SellerID = seller
	resp, _ := product.CreateProduct()
	utils.Respond(w, resp)

}
