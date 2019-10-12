package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/luqmansen/Coolinary/models"
	"github.com/luqmansen/hanako/utils"
	"net/http"
	"strconv"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint)  //Grab the Id of order creator
	order := &models.Order{}

	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Error while request body"))
		return
	}

	order.BuyerID = userID
	resp, _ := order.CreateOrder()
	utils.Respond(w, resp)

}

var PayOrder = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	userID := r.Context().Value("user").(uint)  //Grab the Id of order creator
	order := &models.Order{}

	err = json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Error while request body"))
		return
	}

	order.BuyerID = userID
	resp, _ := order.PayOrder(uint(id))
	utils.Respond(w, resp)

}

