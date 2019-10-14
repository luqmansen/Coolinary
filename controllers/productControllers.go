package controllers

import (
	"encoding/json"
	"github.com/luqmansen/Coolinary/models"
	u "github.com/luqmansen/Coolinary/utils"
	"github.com/luqmansen/hanako/utils"
	"net/http"
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {

	seller := r.Context().Value("user").(uint) //Grab the Id of seller creator

	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		utils.Respond(w, u.Message(http.StatusInternalServerError, "Error while requesting body"))
		return
	}

	product.SellerID = seller
	resp, _ := product.CreateProduct()
	utils.Respond(w, resp)

}

var GetAllProduct = func(w http.ResponseWriter, r *http.Request) {

	show := r.URL.Query().Get("show")

	data,_ := models.GetAllProduct(show)
	if data == nil {
		u.Respond(w, u.Message(http.StatusNoContent, "Not Found"))
	} else {
		resp := u.Message(http.StatusOK, "Success")
		resp["data"] = data
		u.Respond(w, resp)
	}
}
