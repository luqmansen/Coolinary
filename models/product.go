package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/Coolinary/utils"
	"net/http"
	"strconv"
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

func GetAllProduct(number string) ([]*Product, bool){

	if number == "" {
		number = "20"
	} else if _, err := strconv.Atoi(number); err != nil {
		number = "20"
	}

	product := make([]*Product, 0)
	err := GetDB().Limit(number).Find(&product).Error
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return product, true

}