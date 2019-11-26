package controllers

import (
	"fmt"
	"github.com/luqmansen/Coolinary/models"
	u "github.com/luqmansen/Coolinary/utils"
	"net/http"
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {

	//seller := r.Context().Value("user").(uint) //Grab the Id of seller creator

	err := r.ParseForm()
	if err != nil {
		u.Respond(w, u.Message(http.StatusInternalServerError, "Error while requesting body"))
		fmt.Println(err)
		return
	}

	data := models.Product{}

	data.ProductName = r.Form.Get("nama")
	data.Price = (r.Form.Get("harga"))
	data.SellingArea = r.Form.Get("stok")


	//err := json.NewDecoder(r.Body).Decode(product)
	//if err != nil {
	//	u.Respond(w, u.Message(http.StatusInternalServerError, "Error while requesting body"))
	//	fmt.Println(err)
	//	return
	//}

	//data.SellerID = seller
	resp, _ := data.CreateProduct()
	u.Respond(w, resp)

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
