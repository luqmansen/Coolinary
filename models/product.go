package models

import (
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
)

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	SellerID    uint   `json:"seller_id"`
	Price       uint32 `json:"price"`
	SellingArea string `json:"selling_area"`
}

func (product *Product) ValidateProduct() (map[string]interface{}, bool) {

	if product.ProductName == "" {
		return u.Message(http.StatusBadRequest, "Name can't be empty"), false
	}

	if product.Price <= 0 {
		return u.Message(http.StatusBadRequest, "Price Must Be Set"), false
	}

	if product.SellingArea == "" {
		return u.Message(http.StatusBadRequest, "Selling Area can't be empty"), false
	}

	return u.Message(http.StatusOK, "All Requirement Satisfied"), true
}

func (product *Product) CreateProduct() (map[string]interface{}, bool) {

	if resp, ok := product.ValidateProduct(); !ok {
		return resp, false
	}

	GetDB().Create(product)

	resp := u.Message(http.StatusOK, "New Product Added")
	resp["product"] = product
	return resp, true

}
